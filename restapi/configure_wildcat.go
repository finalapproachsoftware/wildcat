package restapi

import (
	"crypto/tls"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"os"

	"github.com/finalapproachsoftware/wildcat/biz"
	"github.com/finalapproachsoftware/wildcat/models"
	"github.com/finalapproachsoftware/wildcat/restapi/operations"
	"github.com/finalapproachsoftware/wildcat/restapi/operations/tenants"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name wildcat --spec ..\swagger.yml

func configureFlags(api *operations.WildcatAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.WildcatAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	mongoHost := os.Getenv("MONGO_HOST")
	session, err := mgo.Dial(mongoHost)

	if err != nil {

	}

	regTenantHandlers(api, session)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func regTenantHandlers(api *operations.WildcatAPI, session *mgo.Session) {

	api.TenantsCreateHandler = tenants.CreateHandlerFunc(func(params tenants.CreateParams) middleware.Responder {
		tenant := params.Body
		tmgr := biz.NewTenantManager(session)
		err := tmgr.Create(tenant)

		if err != nil {
			msg := err.Error()
			e := &models.Error{
				Message: &msg,
			}
			return tenants.NewCreateDefault(500).WithPayload(e)
		}

		return tenants.NewCreateCreated().WithPayload(tenant)
	})

	api.TenantsSearchHandler = tenants.SearchHandlerFunc(func(params tenants.SearchParams) middleware.Responder {
		return middleware.NotImplemented("operation tenants.Search has not yet been implemented")
	})
}
