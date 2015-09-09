package glog

// Mirrors the internal configuration struct used by glog, to expose settings
// to applications.
type GlogConfig struct {
	// log to standard error instead of files
	ToStderr bool

	// log to standard error as well as files
	AlsoToStderr bool

	// logs at or above this threshold go to stderr
	StderrThreshold string

	// when logging hits line file:N, emit a stack trace
	TraceLocation string

	// comma-separated list of pattern=N settings for file-filtered logging
	Vmodule string

	// log level for V logs: 0, 1, 2, 3, etc.
	Verbosity string
}

func NewConfig() *GlogConfig {
	return &GlogConfig{
		ToStderr:        false,
		AlsoToStderr:    false,
		Verbosity:       "",
		StderrThreshold: "info",
		Vmodule:         "",
		TraceLocation:   "",
	}
}

// Apply the logging configuration to the glog system.
func (c *GlogConfig) Init() {
	logging.toStderr = c.ToStderr
	logging.alsoToStderr = c.AlsoToStderr
	logging.verbosity.Set(c.Verbosity)
	logging.vmodule.Set(c.Vmodule)
	logging.traceLocation.Set(c.TraceLocation)

	setThreshold(c.StderrThreshold)

	logging.setVState(0, nil, false)
	go logging.flushDaemon()
}

// Set the stderr threshold using a string: "info", "warn", "error", or "fatal".
// If not recognized or not set, defaults to "info".
func setThreshold(level string) {
	switch level {
	case "warn":
		logging.stderrThreshold = warningLog
	case "error":
		logging.stderrThreshold = errorLog
	case "fatal":
		logging.stderrThreshold = fatalLog
	default:
		logging.stderrThreshold = infoLog
	}
}
