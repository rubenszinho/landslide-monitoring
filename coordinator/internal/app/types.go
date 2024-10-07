package app

type SensorType int

const (
	SoilSensor SensorType = iota
	PluviometerSensor
)

func (s SensorType) String() string {
	return [...]string{"soil", "pluviometer"}[s]
}

type MeasurementType int

const (
	Temperature MeasurementType = iota
	Humidity
	Salinity
)

func (m MeasurementType) String() string {
	return [...]string{"temperature", "humidity", "salinity"}[m]
}

type RestartType string

const (
	SoftRestart RestartType = "soft"
	FullRestart RestartType = "full"
)

type ShutdownType string

const (
	GracefulShutdown  ShutdownType = "graceful"
	ImmediateShutdown ShutdownType = "immediate"
)
