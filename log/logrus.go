package core

import (
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"golang.org/x/mod/modfile"
)

/* ╭──────────────────────────────────────────╮ */
/* │             DUAL OUPUT TYPE              │ */
/* ╰──────────────────────────────────────────╯ */
type dualOutputHook struct {
	consoleFormatter *logrus.TextFormatter
	fileFormatter    *logrus.TextFormatter
	fileWriter       *os.File
}

func (hook *dualOutputHook) Levels() []logrus.Level {
	return logrus.AllLevels // Apply hook to all log levels
}
func (hook *dualOutputHook) Fire(entry *logrus.Entry) error {
	// Serialize log message for console
	consoleMessage, err := hook.consoleFormatter.Format(entry)
	if err != nil {
		return err
	}
	// Write to console
	_, err = os.Stdout.Write(consoleMessage)
	if err != nil {
		return err
	}

	// Serialize log message for file
	fileMessage, err := hook.fileFormatter.Format(entry)
	if err != nil {
		return err
	}
	// Write to file
	_, err = hook.fileWriter.Write(fileMessage)
	return err
}

/* ╭──────────────────────────────────────────╮ */
/* │             LOGRUS FUNCTIONS             │ */
/* ╰──────────────────────────────────────────╯ */
func InitLogRus() *log.Logger {
	// Log as JSON instead of the default ASCII formatter.

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// URI de tu base de datos MongoDB remota
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,

		// DisableColors: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// ExampleLogrus()
	return log.New()
}
func ExampleLogrus() {

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
	// Get TUI instance and load all
	// framework.InitTUI()
}
func InitLogRusWithFile(level logrus.Level) (*log.Logger, func()) {
	// Obtener el nombre del módulo del proyecto
	moduleName := getModuleName()
	logFilePath := moduleName + ".log"

	log := InitLogRus()

	// Set log level
	log.SetLevel(level)

	// Disable default Logrus output to avoid duplicates
	log.SetOutput(io.Discard)

	// Open log file
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// Create custom hook
	hook := &dualOutputHook{
		consoleFormatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: false,
		},
		fileFormatter: &logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		},
		fileWriter: logFile,
	}

	// Add hook to logger
	log.AddHook(hook)

	// Clean up the log file when the program exits
	cleanup := func() error {
		return logFile.Close()
	}
	cleanUpfunc := func() {
		log.Infof("😎 Thanks for using %s", aurora.Bold(aurora.BrightBlue(moduleName)))

		if err := cleanup(); err != nil {
			log.Fatalf("failed to close log file: %v", err)
			os.Exit(1)
		}
	}

	// Example log messages
	log.Infof("🤘 Starting %s", aurora.Bold(aurora.BrightBlue(moduleName)))
	// log.Warnf("⚠️ This is a warning message")
	// log.Errorf("🧨 This is an error message")
	// log.Debugf("👀 This is a debug message")

	return log, cleanUpfunc
}
func InitLogrusOnlyFile(level logrus.Level) (*log.Logger, func()) {
	// Obtener el nombre del módulo del proyecto
	moduleName := getModuleName()
	logFilePath := moduleName + ".log"

	log := InitLogRus()

	// Set log level
	log.SetLevel(level)

	// Disable default Logrus output to avoid duplicates
	log.SetOutput(io.Discard)

	// Formateador para el archivo
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05", // Formato claro para fechas
	})

	// Abrir archivo de log
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// Configurar salida al archivo
	log.SetOutput(logFile)

	// Cleanup function para cerrar el archivo al finalizar
	cleanup := func() error {
		return logFile.Close()
	}

	// Función para manejar errores al cerrar
	cleanUpfunc := func() {
		if err := cleanup(); err != nil {
			log.Errorf("failed to close log file: %v", err)
			os.Exit(1)
		}
	}

	return log, cleanUpfunc
}

/* ╭──────────────────────────────────────────╮ */
/* │              AUX FUNCTIONS               │ */
/* ╰──────────────────────────────────────────╯ */
// getModuleName parses the go.mod file and returns the module name.
func getModuleName() string {
	// Read the go.mod file
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatalf("failed to read go.mod: %v", err)
	}

	// Parse the go.mod file
	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatalf("failed to parse go.mod: %v", err)
	}

	// Return the module name
	return modFile.Module.Mod.Path
}
