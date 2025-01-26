package core

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
	"txeo-tui-library/ui"

	"github.com/charmbracelet/lipgloss"
	"github.com/gin-gonic/gin"
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
		FullTimestamp:    true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// filename := filepath.Base(f.File)
			// return fmt.Sprintf("%s:%d", filename, f.Line), filepath.Base(f.Function)
			return "", ""
		},
		TimestampFormat: "2006-01-02 15:04:05",
		DisableQuote:    true,
		DisableSorting:  true,
		// FullTimestamp:   true,
		// DisableColors:   false,

		// DisableColors: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// ExampleLogrus()
	return log.New()
}
func InitLogRusRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			}
		}()
		c.Next()
	}
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

/* LoggerForGin returns a Gin middleware that logs through a given Logrus logger.*/
func LoggerForGin(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Inicia el contador de tiempo
		start := time.Now()

		// Procesa la petición
		c.Next()

		// Calcula la latencia
		stop := time.Now()
		latency := stop.Sub(start)
		colorLatencyAsString := GetBackgroundColorForLatency(float64(latency.Seconds()))

		// Formatea la latencia
		latencyAsString := lipgloss.NewStyle().Background(lipgloss.Color(colorLatencyAsString)).Render(fmt.Sprintf(" %.3fs ", latency.Seconds()))
		// Formatea la fecha-hora; Gin usa algo como "2006/01/02 - 15:04:05"
		stopTime := stop.Format("2006/01/02 - 15:04:05")

		// Datos principales
		method := c.Request.Method
		status := c.Writer.Status()
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery
		ipColored := ui.ColoredString(c.ClientIP())
		path = ui.ColoredString(path)
		method = ui.ColoredString(method)
		if rawQuery != "" {
			path = path + "?" + rawQuery
		}
		// Aquí coloreamos el IP

		// Si hubo un error (por ejemplo con c.AbortWithError)
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Selección de colores de estado
		var statusColor aurora.Value
		statusSpaced := " " + strconv.Itoa(status) + " "
		switch {
		case status >= http.StatusInternalServerError:
			statusColor = aurora.BgRed(aurora.White(statusSpaced))
		case status >= http.StatusBadRequest:
			statusColor = aurora.BgYellow(aurora.Black(statusSpaced))
		case status >= http.StatusMultipleChoices:
			statusColor = aurora.BgCyan(aurora.Black(statusSpaced))
		default:
			statusColor = aurora.BgGreen(aurora.Black(statusSpaced))
		}

		// Colores de método
		var methodColor aurora.Value
		methodSpaced := " " + method + " "
		switch method {
		case "GET":
			methodColor = aurora.Bold(aurora.Green(methodSpaced))
		case "POST":
			methodColor = aurora.Bold(aurora.BgBrightCyan(methodSpaced))
		case "PUT":
			methodColor = aurora.Bold(aurora.Blue(methodSpaced))
		case "DELETE":
			methodColor = aurora.Bold(aurora.Red(methodSpaced))
		default:
			methodColor = aurora.Bold(aurora.White(methodSpaced))
		}

		// Muestra el error si existe (en rojo)
		var redError string
		if errorMessage != "" {
			redError = " " + aurora.Red(errorMessage).String()
		}

		// Crea la entrada de log con un par de campos, si lo deseas
		entry := logger.WithTime(time.Now())
		if errorMessage != "" {
			entry = entry.WithField("error", errorMessage)
		}

		// Formato con la fecha/hora
		// Ejemplo: [GIN] 2025/01/26 - 09:10:46 | 401 | 1.448125ms | 99.80.255.172 | POST    "/..." <error>
		lineFormat := "[GIN] %s | %v | %10v | %15s | %-7s %s%s"
		//           fecha/hora  status latencia   IP            metodo  path     error

		switch {
		case status >= http.StatusInternalServerError:
			entry.Errorf(lineFormat,
				stopTime, statusColor, latencyAsString, ipColored, methodColor, path, redError)
		case status >= http.StatusBadRequest:
			entry.Warnf(lineFormat,
				stopTime, statusColor, latencyAsString, ipColored, methodColor, path, redError)
		default:
			entry.Infof(lineFormat,
				stopTime, statusColor, latencyAsString, ipColored, methodColor, path, redError)
		}
	}
}
func GetBackgroundColorForLatency(latency float64) string {
	// Si la latencia es negativa (invalida), forzamos un color "neutro"
	if latency < 0 {
		return "#A2079A" // Morado (indicador de error)
	}

	// Si la latencia supera 5 segundos, usamos rojo máximo
	if latency >= 5.0 {
		return "#FF0000"
	}

	// Busca el rango más cercano en LatencyColorMap
	var lowerBound, upperBound float64
	for bound := range ui.LatencyColorMap {
		if bound <= latency && bound > lowerBound {
			lowerBound = bound
		}
		if bound > latency && (upperBound == 0 || bound < upperBound) {
			upperBound = bound
		}
	}

	// Interpolar entre lowerBound y upperBound
	lowerColor := ui.LatencyColorMap[lowerBound]
	upperColor := ui.LatencyColorMap[upperBound]

	// Si no hay interpolación posible, devuelve el color más cercano
	if upperBound == 0 || lowerBound == upperBound {
		return lowerColor
	}

	// Porcentaje de interpolación
	t := (latency - lowerBound) / (upperBound - lowerBound)

	// Interpolar colores (de #RRGGBB)
	return ui.InterpolateHexColor(lowerColor, upperColor, t)
}
