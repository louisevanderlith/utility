package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)

	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewService))).Methods(http.MethodGet)

	r.Handle("/info", mw.Handler(http.HandlerFunc(GetServices))).Methods(http.MethodGet)
	r.Handle("/info/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchServices))).Methods(http.MethodGet)

	r.Handle("/info", mw.Handler(http.HandlerFunc(CreateService))).Methods(http.MethodPost)
	r.Handle("/info/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateService))).Methods(http.MethodPut)

	/*

		//services
		r.Handle("/services", mw.Handler(http.HandlerFunc(GetServices))).Methods(http.MethodGet)
		r.Handle("/services/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewService))).Methods(http.MethodGet)
		r.Handle("/services/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchServices))).Methods(http.MethodGet)
		r.Handle("/services/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchServices))).Methods(http.MethodGet)
		r.Handle("/services", mw.Handler(http.HandlerFunc(CreateService))).Methods(http.MethodPost)
		r.Handle("/services/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateService))).Methods(http.MethodPut)

	*/

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
