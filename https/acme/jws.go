// fork from https://github.com/rsc/letsencrypt/tree/master/vendor/github.com/xenolf/lego/acme
// fork from https://github.com/xenolf/lego/tree/master/acme
package acme

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/winlinvip/mod_migrate/https/jose"
)

type jws struct {
	directoryURL string
	privKey      crypto.PrivateKey
	nonces       []string
}

func keyAsJWK(key interface{}) *jose.JsonWebKey {
	switch k := key.(type) {
	case *ecdsa.PublicKey:
		return &jose.JsonWebKey{Key: k, Algorithm: "EC"}
	case *rsa.PublicKey:
		return &jose.JsonWebKey{Key: k, Algorithm: "RSA"}

	default:
		return nil
	}
}

// Posts a JWS signed message to the specified URL
func (j *jws) post(url string, content []byte) (*http.Response, error) {
	signedContent, err := j.signContent(content)
	if err != nil {
		return nil, err
	}

	resp, err := httpPost(url, "application/jose+json", bytes.NewBuffer([]byte(signedContent.FullSerialize())))
	if err != nil {
		return nil, err
	}

	j.getNonceFromResponse(resp)

	return resp, err
}

func (j *jws) signContent(content []byte) (*jose.JsonWebSignature, error) {

	var alg jose.SignatureAlgorithm
	switch k := j.privKey.(type) {
	case *rsa.PrivateKey:
		alg = jose.RS256
	case *ecdsa.PrivateKey:
		if k.Curve == elliptic.P256() {
			alg = jose.ES256
		} else if k.Curve == elliptic.P384() {
			alg = jose.ES384
		}
	}

	signer, err := jose.NewSigner(alg, j.privKey)
	if err != nil {
		return nil, err
	}
	signer.SetNonceSource(j)

	signed, err := signer.Sign(content)
	if err != nil {
		return nil, err
	}
	return signed, nil
}

func (j *jws) getNonceFromResponse(resp *http.Response) error {
	nonce := resp.Header.Get("Replay-Nonce")
	if nonce == "" {
		return fmt.Errorf("Server did not respond with a proper nonce header.")
	}

	j.nonces = append(j.nonces, nonce)
	return nil
}

func (j *jws) getNonce() error {
	resp, err := httpHead(j.directoryURL)
	if err != nil {
		return err
	}

	return j.getNonceFromResponse(resp)
}

func (j *jws) Nonce() (string, error) {
	nonce := ""
	if len(j.nonces) == 0 {
		err := j.getNonce()
		if err != nil {
			return nonce, err
		}
	}

	nonce, j.nonces = j.nonces[len(j.nonces)-1], j.nonces[:len(j.nonces)-1]
	return nonce, nil
}
