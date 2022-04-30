package repo

// Driver
const (
	queryCreateNewDriver = `INSERT INTO drivers(name, phone, id_card, driver_license, status, shipment_status) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	queryGetAllDrivers   = `SELECT * FROM drivers  WHERE ($1 = '' OR name ILIKE $1) OFFSET $2 LIMIT $3`
	queryGetDriverById   = `SELECT * FROM drivers WHERE id=$1`
	queryUpdateDriver    = `UPDATE drivers SET name=$1, phone=$2, id_card=$3, driver_license=$4, status=$5, shipment_status=$6, updated_at=$7 WHERE id=$8`
	queryDeleteDriver    = `UPDATE trucks status=true WHERE id=$1`
)

// Truck
const (
	queryCreateNewTruck = `INSERT INTO trucks(license_number, truck_type, plate_type, production_year, status, shipment_status, stnk, kir) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	queryGetAllTrucks   = `SELECT * FROM trucks WHERE ($1 = '' OR license_number ILIKE $1) AND  ($1 = '' OR truck_type ILIKE $2) OFFSET $3 LIMIT $4`
	queryGetTruckById   = `SELECT * FROM trucks WHERE id=$1`
	queryUpdateTruck    = `UPDATE trucks SET license_number=$1, truck_type=$2, plate_type=$3, production_year=$4, status=$5, stnk=$6, kir=$7, updated_at=now() WHERE id=$8`
	queryDeleteTruck    = `UPDATE trucks status=true WHERE id=$1` // SOFT DELETE
)

// Shipment
const (
	queryCreateNewShipment = `INSERT INTO shipments(no_shipment, origin, destination, status, loading_date) VALUES($1, $2, $3, 1, $4) RETURNING id`
	queryAllocateShipment  = `UPDATE shipments(id_driver, id_truck, status) VALUES ($1, $2, 2) RETURNING id`
	queryGetAllShipments   = `SELECT * FROM trucks`
	queryGetShipmentsById  = `SELECT * FROM shipments WHERE id=$1`
	queryUpdateShipment    = `UPDATE shipments SET no_shipment=$1, id_driver=$2, id_truck=$3, origin=$4, destination=$5, status=$6, loading_date=$7, updated_at=$9 WHERE id=$10`
	queryDeleteShipment    = `UPDATE shipment deleted_at=$1 WHERE id=$2`
)
