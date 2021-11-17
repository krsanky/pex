package session

import (
	"database/sql"
	"net/http"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"go.d34d.net/pex/lg"
)

var Session *scs.Session

func Init(db *sql.DB) {
	lg.Log.Printf("session.Init()")
	Session = scs.New()
	lg.Log.Printf("session.Init() 32")
	Session.Store = postgresstore.New(db)
	lg.Log.Printf("session.Init() 34")

	Session.Cookie.Name = "_pex_sess_"
	Session.Cookie.Persist = true
	lg.Log.Printf("session.Init()...")
}

func GetWithDefault(r *http.Request, key string, def string) (value string, err error) {
	value = Session.GetString(r.Context(), key)
	if value == "" {
		value = def
	}
	return
}
