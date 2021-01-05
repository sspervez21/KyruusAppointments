// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"KyruusAppointments/app"
	"KyruusAppointments/restapi/operations"
)

//go:generate swagger generate server --target ..\..\KyruusAppointments --name Kyruus --spec ..\swagger.yml

func configureFlags(api *operations.KyruusAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.KyruusAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CreateDoctorHandler = operations.CreateDoctorHandlerFunc(func(params operations.CreateDoctorParams) middleware.Responder {
		return app.CreateDoctor(params)
	})
	api.GetDoctorHandler = operations.GetDoctorHandlerFunc(func(params operations.GetDoctorParams) middleware.Responder {
		return app.GetDoctor(params)
	})
	api.DeleteDoctorHandler = operations.DeleteDoctorHandlerFunc(func(params operations.DeleteDoctorParams) middleware.Responder {
		return app.DeleteDoctor(params)
	})
	api.UpdateDoctorHandler = operations.UpdateDoctorHandlerFunc(func(params operations.UpdateDoctorParams) middleware.Responder {
		return app.UpdateDoctor(params)
	})

	api.CreateAppointmentHandler = operations.CreateAppointmentHandlerFunc(func(params operations.CreateAppointmentParams) middleware.Responder {
		return app.CreateAppointment(params)
	})
	api.GetAppointmentsHandler = operations.GetAppointmentsHandlerFunc(func(params operations.GetAppointmentsParams) middleware.Responder {
		return app.GetAppointments(params)
	})
	api.DeleteAppointmentHandler = operations.DeleteAppointmentHandlerFunc(func(params operations.DeleteAppointmentParams) middleware.Responder {
		return app.DeleteAppointment(params)
	})

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
func configureServer(s *http.Server, scheme, addr string) {
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
