package drivers

var DriverRegistry *_driverRegistry

type _driverRegistry struct {
	driverHandlers map[string]DriverHandler
}

func RegisterHandler(id string, handler DriverHandler) {
	if DriverRegistry.driverHandlers == nil {
		DriverRegistry.driverHandlers = make(map[string]DriverHandler)
	}
	DriverRegistry.driverHandlers[id] = handler
}
