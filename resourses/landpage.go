package resourses

import (
	"github.com/go-chi/chi/v5"
	"github.com/marioteik/helpers"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type LandingPage struct {
	DB *gorm.DB
	TH *helpers.TemplateHelper
	//Session *scs.SessionManager
}

func (lp LandingPage) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", lp.Page)

	return r
}

func (lp LandingPage) Page(w http.ResponseWriter, r *http.Request) {
	//remoteIP := r.RemoteAddr
	//admin.Session.Put(r.Context(), "remote_ip", remoteIP)

	err := lp.TH.RenderTemplate(w, "landing.page.gohtml", nil)
	if err != nil {
		log.Println(err)
		//helpers.ServerError(w, err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}
