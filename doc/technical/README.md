# Documentación Técnica - txeo-tools-library

## Arquitectura

La biblioteca txeo-tools-library está organizada como una colección de paquetes independientes pero complementarios, cada uno centrado en un dominio específico de funcionalidad. Cada paquete sigue estos principios de diseño:

1. **Independencia**: Mínimas dependencias entre paquetes
2. **Simplicidad**: API sencilla y directa
3. **Reutilización**: Componentes altamente reutilizables
4. **Rendimiento**: Optimización para casos de uso frecuentes
5. **Fiabilidad**: Manejo exhaustivo de errores

## Paquetes Principales

### files

Proporciona funcionalidades para trabajar con archivos y directorios:

```go
// Leer contenido de archivo
content, err := files.ReadFile("config.json")

// Escribir contenido a archivo
err := files.WriteFile("output.txt", []byte("Contenido"), 0644)

// Comprobar si un archivo existe
exists := files.Exists("data.csv")

// Leer directorio
entries, err := files.ReadDir("/path/to/dir")

// Crear directorio recursivamente
err := files.MkdirAll("/path/to/new/dir", 0755)

// Copiar archivo
err := files.Copy("source.txt", "destination.txt")

// Mover archivo
err := files.Move("old.txt", "new.txt")

// Buscar archivos por patrón
matches, err := files.Find("/path/to/dir", "*.go")
```

### convert

Facilita la conversión entre diferentes tipos de datos:

```go
// Conversiones de strings
intVal := convert.StringToInt("123", 0)            // Default 0 si hay error
floatVal := convert.StringToFloat("123.45", 0.0)   // Default 0.0 si hay error
boolVal := convert.StringToBool("true", false)     // Default false si hay error

// Conversiones a string
str := convert.IntToString(123)
str := convert.FloatToString(123.45, 2)  // 2 decimales
str := convert.BoolToString(true)

// Conversiones de estructuras
json := convert.StructToJSON(myStruct)
xml := convert.StructToXML(myStruct)
map := convert.StructToMap(myStruct)

// Conversiones a estructuras
err := convert.JSONToStruct(json, &myStruct)
err := convert.XMLToStruct(xml, &myStruct)
err := convert.MapToStruct(map, &myStruct)
```

### format

Proporciona funciones para formatear datos:

```go
// Formateo de números
str := format.Number(1234567.89, 2)      // "1,234,567.89"
str := format.Currency(1234.56, "$")     // "$1,234.56"
str := format.Percentage(0.1234, 2)      // "12.34%"

// Formateo de textos
str := format.PadLeft("text", 10, ' ')   // "      text"
str := format.PadRight("text", 10, ' ')  // "text      "
str := format.Truncate("texto largo", 5) // "texto..."
str := format.CamelCase("hello world")   // "helloWorld"
str := format.SnakeCase("HelloWorld")    // "hello_world"

// Formateo de tablas
table := format.NewTable([]string{"ID", "Nombre", "Edad"})
table.AddRow([]string{"1", "Juan", "30"})
table.AddRow([]string{"2", "María", "25"})
str := table.String()
```

### env

Gestiona variables de entorno y configuración:

```go
// Acceso a variables de entorno
val := env.Get("HOME", "/default/home")

// Carga de variables desde archivo .env
err := env.Load(".env")

// Verificar existencia de variable
exists := env.Exists("API_KEY")

// Configuración tipada
port := env.GetInt("PORT", 8080)
debug := env.GetBool("DEBUG", false)
timeout := env.GetDuration("TIMEOUT", 30*time.Second)

// Carga de configuración desde JSON/YAML
config := env.LoadConfig("config.json")
val := config.GetString("database.host")
```

### logger

Sistema de logging flexible:

```go
// Crear logger
log := logger.New("component")

// Niveles de log
log.Debug("Mensaje de debug")
log.Info("Mensaje informativo")
log.Warn("Advertencia")
log.Error("Error: %v", err)
log.Fatal("Error fatal")

// Configuración
logger.SetLevel(logger.LevelInfo)      // Establecer nivel mínimo
logger.SetOutput(os.Stdout)            // Establecer salida
logger.SetFormat("[{time}] {level} {component} - {message}")  // Formato

// Logger con contexto
contextLog := log.WithContext(map[string]interface{}{
    "requestID": "abc123",
    "userID": 456,
})
contextLog.Info("Procesando solicitud")
```

### time

Utilidades para manejo de tiempo:

```go
// Formateadores
str := time.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
date := time.Parse("2006-01-02", "2025-05-03")

// Manipulación
tomorrow := time.AddDays(time.Now(), 1)
nextMonth := time.AddMonths(time.Now(), 1)
start, end := time.MonthRange(time.Now())

// Diferencias
days := time.DaysBetween(start, end)
duration := time.Duration(start, end)
humanReadable := time.HumanDuration(duration) // "2 días 3 horas"

// Intervalos
isWithin := time.IsWithinRange(time.Now(), start, end)
```

### crypto

Funciones criptográficas básicas:

```go
// Hashing
hash := crypto.MD5("texto")
hash := crypto.SHA256("texto")

// Encriptación/Desencriptación
encrypted, err := crypto.Encrypt("texto", "clave")
decrypted, err := crypto.Decrypt(encrypted, "clave")

// Generación de tokens
token := crypto.GenerateToken(32)  // 32 bytes aleatorios

// UUID
uuid := crypto.UUID()

// Password
hashedPwd := crypto.HashPassword("password")
isMatch := crypto.VerifyPassword("password", hashedPwd)
```

### net

Utilidades para operaciones de red:

```go
// HTTP
response, err := net.Get("https://api.example.com/data")
response, err := net.Post("https://api.example.com/data", data, "application/json")

// Validación
isValid := net.IsValidURL("https://example.com")
isValid := net.IsValidEmail("user@example.com")
isValid := net.IsValidIP("192.168.1.1")

// Utilidades
ip := net.GetPublicIP()
isReachable := net.Ping("example.com", 1*time.Second)
```

### async

Herramientas para operaciones asíncronas:

```go
// Pool de Workers
pool := async.NewWorkerPool(10)
for i := 0; i < 100; i++ {
    task := func(id int) {
        // Trabajo a realizar
    }
    pool.Submit(func() { task(i) })
}
pool.Wait()

// Tareas con timeout
result, err := async.WithTimeout(func() (interface{}, error) {
    // Tarea larga
    return "resultado", nil
}, 5*time.Second)

// Ejecución periódica
stopFunc := async.Schedule(func() {
    // Tarea programada
}, 1*time.Minute)
// Más tarde...
stopFunc()

// Semáforo
sem := async.NewSemaphore(5)
sem.Acquire()
// Sección crítica
sem.Release()
```

## Manejo de Errores

La biblioteca implementa un enfoque consistente para el manejo de errores:

1. **Errores Tipados**: Los paquetes definen tipos de error específicos
2. **Información Contextual**: Los errores incluyen información detallada
3. **Capacidad de Anidamiento**: Soporte para envolver errores
4. **Recuperación**: Utilidades para recuperación de pánico

```go
// Creación de errores
err := files.NewError("file_not_found", "El archivo no existe: %s", path)

// Comprobación de tipos de error
if files.IsNotExistError(err) {
    // Manejo específico
}

// Envoltura de errores
err = files.WrapError(originalErr, "error al procesar archivo")

// Recuperación de pánico
defer func() {
    if r := recover(); r != nil {
        err = fmt.Errorf("recuperado de pánico: %v", r)
    }
}()
```

## Extensibilidad

Los paquetes están diseñados para ser extendidos:

1. **Interfaces**: Uso de interfaces para componentes clave
2. **Fábricas**: Patrones de fábrica para creación personalizada
3. **Middlewares**: Soporte para intercepción y modificación
4. **Plugins**: Sistema modular para extensiones