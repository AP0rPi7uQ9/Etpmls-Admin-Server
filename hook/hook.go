package hook

import (
	"Etpmls-Admin-Server/library"
)

type Hook struct {

}

func (this *Hook) ExitApplication() {
	_ = library.NewConsul().CancelService()
}