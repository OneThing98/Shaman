package sensors

type TriggerEvent struct {
	TriggerType string
	Payload     map[string]interface{} //key-value data describing the event
	Meta        map[string]string      //optional metadata like timestamps
}

//Sensor is the interface every sensor plugin must implement

type Sensor interface {
	//returns the unique name for a sensor, used in logs and monitoring
	Name() string

	//optional version info
	Version() string

	//called before sensor starts. It receives sensor specific config.
	Init(config map[string]interface{}) error

	//Starts the sensor. The sensor should emit trigger event objects on eventCh
	Start(eventCh chan<- TriggerEvent) error

	//Stops the error gracefully returning an error if something goes wrong
}
