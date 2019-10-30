// Please use library.
package main

import (
	"fmt"
	_ "github.com/winlinvip/mod_migrate/aac"
	_ "github.com/winlinvip/mod_migrate/asprocess"
	_ "github.com/winlinvip/mod_migrate/errors"
	_ "github.com/winlinvip/mod_migrate/flv"
	_ "github.com/winlinvip/mod_migrate/gmoryx"
	_ "github.com/winlinvip/mod_migrate/http"
	_ "github.com/winlinvip/mod_migrate/https"
	_ "github.com/winlinvip/mod_migrate/json"
	_ "github.com/winlinvip/mod_migrate/kxps"
	_ "github.com/winlinvip/mod_migrate/logger"
	_ "github.com/winlinvip/mod_migrate/options"
	_ "github.com/winlinvip/mod_migrate/websocket"
)

const (
	Major, Minor, Revision = 0, 0, 1
)

func Version() string {
	return fmt.Sprintf("%v.%v.%v", Major, Minor, Revision)
}

func main() {
	fmt.Println(fmt.Sprintf("GO-ORYX-LIB/%v, please use as library in your project.", Version()))
	return
}
