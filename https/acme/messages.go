// fork from https://github.com/rsc/letsencrypt/tree/master/vendor/github.com/xenolf/lego/acme
// fork from https://github.com/xenolf/lego/tree/master/acme
package acme

import (
	"time"

	"github.com/winlinvip/mod_migrate/https/jose"
)

type directory struct {
	NewAuthzURL   string `json:"new-authz"`
	NewCertURL    string `json:"new-cert"`
	NewRegURL     string `json:"new-reg"`
	RevokeCertURL string `json:"revoke-cert"`
}

type recoveryKeyMessage struct {
	Length int             `json:"length,omitempty"`
	Client jose.JsonWebKey `json:"client,omitempty"`
	Server jose.JsonWebKey `json:"client,omitempty"`
}

type registrationMessage struct {
	Resource string   `json:"resource"`
	Contact  []string `json:"contact"`
	//	RecoveryKey recoveryKeyMessage `json:"recoveryKey,omitempty"`
}

// Registration is returned by the ACME server after the registration
// The client implementation should save this registration somewhere.
type Registration struct {
	Resource       string          `json:"resource,omitempty"`
	ID             int             `json:"id"`
	Key            jose.JsonWebKey `json:"key"`
	Contact        []string        `json:"contact"`
	Agreement      string          `json:"agreement,omitempty"`
	Authorizations string          `json:"authorizations,omitempty"`
	Certificates   string          `json:"certificates,omitempty"`
	//	RecoveryKey    recoveryKeyMessage `json:"recoveryKey,omitempty"`
}

// RegistrationResource represents all important informations about a registration
// of which the client needs to keep track itself.
type RegistrationResource struct {
	Body        Registration `json:"body,omitempty"`
	URI         string       `json:"uri,omitempty"`
	NewAuthzURL string       `json:"new_authzr_uri,omitempty"`
	TosURL      string       `json:"terms_of_service,omitempty"`
}

type authorizationResource struct {
	Body       authorization
	Domain     string
	NewCertURL string
	AuthURL    string
}

type authorization struct {
	Resource     string      `json:"resource,omitempty"`
	Identifier   identifier  `json:"identifier"`
	Status       string      `json:"status,omitempty"`
	Expires      time.Time   `json:"expires,omitempty"`
	Challenges   []challenge `json:"challenges,omitempty"`
	Combinations [][]int     `json:"combinations,omitempty"`
}

type identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type validationRecord struct {
	URI               string   `json:"url,omitempty"`
	Hostname          string   `json:"hostname,omitempty"`
	Port              string   `json:"port,omitempty"`
	ResolvedAddresses []string `json:"addressesResolved,omitempty"`
	UsedAddress       string   `json:"addressUsed,omitempty"`
}

type challenge struct {
	Resource          string             `json:"resource,omitempty"`
	Type              Challenge          `json:"type,omitempty"`
	Status            string             `json:"status,omitempty"`
	URI               string             `json:"uri,omitempty"`
	Token             string             `json:"token,omitempty"`
	KeyAuthorization  string             `json:"keyAuthorization,omitempty"`
	TLS               bool               `json:"tls,omitempty"`
	Iterations        int                `json:"n,omitempty"`
	Error             RemoteError        `json:"error,omitempty"`
	ValidationRecords []validationRecord `json:"validationRecord,omitempty"`
}

type csrMessage struct {
	Resource       string   `json:"resource,omitempty"`
	Csr            string   `json:"csr"`
	Authorizations []string `json:"authorizations"`
}

type revokeCertMessage struct {
	Resource    string `json:"resource"`
	Certificate string `json:"certificate"`
}

// CertificateResource represents a CA issued certificate.
// PrivateKey and Certificate are both already PEM encoded
// and can be directly written to disk. Certificate may
// be a certificate bundle, depending on the options supplied
// to create it.
type CertificateResource struct {
	Domain        string `json:"domain"`
	CertURL       string `json:"certUrl"`
	CertStableURL string `json:"certStableUrl"`
	AccountRef    string `json:"accountRef,omitempty"`
	PrivateKey    []byte `json:"-"`
	Certificate   []byte `json:"-"`
}
