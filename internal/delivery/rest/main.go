package rest

import (
	"github.com/febriliankr/go-clean-architecture/internal/app"
	"github.com/febriliankr/go-clean-architecture/internal/delivery/rest/server"
	"github.com/febriliankr/go-clean-architecture/internal/delivery/rest/service"
)

// Start server
func Start(app *app.ShipmentApp) {
	srv := server.New()
	svc := service.GetServices(app)

	// Init Trucks Handlers
	trucks := srv.Group("/trucks")
	trucks.GET("", svc.GetTruckListHandler)
	trucks.GET("/:id", svc.GetTruckByIDHandler)
	trucks.POST("", svc.CreateNewTruckHandler)
	trucks.PUT("/:id", svc.UpdateTruckHandler)
	trucks.DELETE("/:id", svc.DeleteTruckHandler)

	// Init Drivers Handlers
	drivers := srv.Group("/drivers")
	drivers.GET("", svc.GetDriverListHandler)
	drivers.GET("/:id", svc.GetDriverByIDHandler)
	drivers.POST("", svc.CreateNewDriverHandler)
	drivers.PUT("/:id", svc.UpdateDriverHandler)
	drivers.DELETE("/:id", svc.DeleteDriverHandler)

	// Init Shipments Handlers
	shipments := srv.Group("/shipments")
	shipments.GET("", svc.GetShipmentListHandler)
	shipments.GET("/:id", svc.GetShipmentByIDHandler)
	shipments.POST("", svc.CreateNewShipmentHandler)
	shipments.POST("/allocate", svc.AllocateShipmentHandler)
	shipments.PUT("/:id", svc.UpdateShipmentHandler)
	shipments.DELETE("/:id", svc.DeleteShipmentHandler)

	server.Start(srv, &app.Cfg.Server)
}
