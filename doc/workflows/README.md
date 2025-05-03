# Flujos de Trabajo - txeo-tools-library

## Patrones de Uso Comunes

### Gestión de Archivos

#### Procesamiento de CSV
```go
package main

import (
    "fmt"
    "github.com/txeo/tools-library/files"
    "github.com/txeo/tools-library/logger"
)

func main() {
    log := logger.New("csv-processor")
    
    // Leer archivo CSV
    content, err := files.ReadCSV("datos.csv")
    if err != nil {
        log.Error("Error al leer CSV: %v", err)
        return
    }
    
    // Procesar filas
    for i, row := range content {
        if i == 0 {
            log.Info("Encabezados: %v", row)
            continue
        }
        
        // Procesar cada fila
        processRow(row)
    }
    
    // Escribir resultados
    resultRows := [][]string{
        {"ID", "Nombre", "Resultado"},
        {"1", "Producto A", "Procesado"},
        {"2", "Producto B", "Procesado"},
    }
    
    err = files.WriteCSV("resultados.csv", resultRows)
    if err != nil {
        log.Error("Error al escribir CSV: %v", err)
        return
    }
    
    log.Info("Procesamiento completado")
}

func processRow(row []string) {
    // Lógica de procesamiento
}
```

#### Operaciones por Lotes con Archivos
```go
// Buscar y procesar archivos en lote
func processDirectory(dirPath string) {
    // Encontrar todos los archivos .txt
    files, err := files.Find(dirPath, "*.txt")
    if err != nil {
        log.Error("Error al buscar archivos: %v", err)
        return
    }
    
    // Procesar cada archivo
    for _, filePath := range files {
        content, err := files.ReadFile(filePath)
        if err != nil {
            log.Warn("Error al leer %s: %v", filePath, err)
            continue
        }
        
        // Procesar contenido
        processedContent := processContent(content)
        
        // Guardar resultado
        outputPath := filePath + ".processed"
        err = files.WriteFile(outputPath, processedContent, 0644)
        if err != nil {
            log.Error("Error al escribir %s: %v", outputPath, err)
            continue
        }
        
        log.Info("Procesado: %s", filePath)
    }
}
```

### Logging y Monitoreo

#### Configuración Avanzada de Logger
```go
// Configuración para diferentes entornos
func setupLogging(environment string) {
    switch environment {
    case "development":
        logger.SetLevel(logger.LevelDebug)
        logger.SetFormat("[{time}] {level} {component} - {message}")
        logger.SetOutput(os.Stdout)
    
    case "production":
        // Log a archivo
        logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            fmt.Printf("Error al abrir archivo de log: %v\n", err)
            os.Exit(1)
        }
        
        logger.SetLevel(logger.LevelInfo)
        logger.SetFormat("[{time}] {level} {component} - {message}")
        logger.SetOutput(logFile)
        
        // Opcional: logger de errores separado
        errorLogFile, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        logger.SetErrorOutput(errorLogFile)
    
    case "testing":
        logger.SetLevel(logger.LevelWarn)
        logger.SetOutput(io.Discard) // Descartar logs en pruebas
    }
}
```

#### Logging con Contexto
```go
func processRequest(requestID string, userID int, data map[string]interface{}) {
    // Crear logger con contexto
    log := logger.New("request-handler").WithContext(map[string]interface{}{
        "requestID": requestID,
        "userID": userID,
    })
    
    log.Info("Inicio de procesamiento de solicitud")
    
    // Procesar datos
    for key, value := range data {
        log.Debug("Procesando campo: %s = %v", key, value)
        
        // Lógica de procesamiento
        result, err := processField(key, value)
        
        if err != nil {
            log.Error("Error al procesar campo %s: %v", key, err)
            continue
        }
        
        log.Debug("Campo procesado: %s, resultado: %v", key, result)
    }
    
    log.Info("Solicitud procesada completamente")
}
```

### Concurrencia y Operaciones Asíncronas

#### Procesamiento Paralelo con Worker Pool
```go
func processItems(items []Item) []Result {
    results := make([]Result, len(items))
    
    // Crear pool de workers
    pool := async.NewWorkerPool(runtime.NumCPU())
    
    // Semáforo para sincronización
    var wg sync.WaitGroup
    
    // Enviar tareas al pool
    for i, item := range items {
        wg.Add(1)
        
        // Capturar variables del loop
        index := i
        currentItem := item
        
        pool.Submit(func() {
            defer wg.Done()
            
            // Procesar item (operación que puede llevar tiempo)
            result, err := processItem(currentItem)
            
            if err != nil {
                log.Error("Error procesando item %d: %v", index, err)
                results[index] = Result{Error: err}
                return
            }
            
            results[index] = result
        })
    }
    
    // Esperar a que todas las tareas terminen
    wg.Wait()
    
    return results
}
```

#### Operaciones con Timeout
```go
func fetchDataWithTimeout(url string, timeout time.Duration) ([]byte, error) {
    result, err := async.WithTimeout(func() (interface{}, error) {
        // Operación que puede tardar
        return net.Get(url)
    }, timeout)
    
    if err != nil {
        if async.IsTimeoutError(err) {
            return nil, fmt.Errorf("timeout al obtener datos de %s", url)
        }
        return nil, err
    }
    
    return result.([]byte), nil
}
```

### Configuración y Variables de Entorno

#### Carga de Configuración desde Múltiples Fuentes
```go
func loadConfiguration() *Config {
    config := &Config{}
    
    // 1. Valores predeterminados
    config.Port = 8080
    config.Timeout = 30 * time.Second
    config.DatabaseURL = "localhost:5432"
    
    // 2. Archivo de configuración
    if files.Exists("config.json") {
        configData, err := env.LoadConfig("config.json")
        if err == nil {
            config.Port = configData.GetInt("server.port", config.Port)
            config.Timeout = configData.GetDuration("server.timeout", config.Timeout)
            config.DatabaseURL = configData.GetString("database.url", config.DatabaseURL)
        }
    }
    
    // 3. Variables de entorno (tienen prioridad)
    if env.Exists("APP_PORT") {
        config.Port = env.GetInt("APP_PORT", config.Port)
    }
    if env.Exists("APP_TIMEOUT") {
        config.Timeout = env.GetDuration("APP_TIMEOUT", config.Timeout)
    }
    if env.Exists("APP_DB_URL") {
        config.DatabaseURL = env.Get("APP_DB_URL", config.DatabaseURL)
    }
    
    return config
}
```

## Escenarios de Uso Común

### Importación y Transformación de Datos

```go
func importDataFromCSV(filePath string) error {
    // Leer CSV
    data, err := files.ReadCSV(filePath)
    if err != nil {
        return fmt.Errorf("error al leer CSV: %w", err)
    }
    
    // Verificar estructura
    if len(data) < 2 {
        return fmt.Errorf("CSV vacío o sin datos suficientes")
    }
    
    headers := data[0]
    headerMap := make(map[string]int)
    for i, header := range headers {
        headerMap[header] = i
    }
    
    // Verificar columnas requeridas
    requiredColumns := []string{"ID", "Nombre", "Valor"}
    for _, col := range requiredColumns {
        if _, exists := headerMap[col]; !exists {
            return fmt.Errorf("columna requerida no encontrada: %s", col)
        }
    }
    
    // Procesar filas
    for i, row := range data {
        if i == 0 {
            continue // Saltar encabezados
        }
        
        // Extraer datos
        id := row[headerMap["ID"]]
        nombre := row[headerMap["Nombre"]]
        valorStr := row[headerMap["Valor"]]
        
        // Convertir tipos
        valor := convert.StringToFloat(valorStr, 0.0)
        
        // Crear objeto de modelo
        item := Item{
            ID: id,
            Nombre: nombre,
            Valor: valor,
        }
        
        // Almacenar (ejemplo)
        err := saveItem(item)
        if err != nil {
            log.Error("Error al guardar item %s: %v", id, err)
            continue
        }
    }
    
    return nil
}
```

### API REST Básica

```go
func startAPIServer(config *Config) error {
    log := logger.New("api-server")
    
    // Iniciar servidor HTTP
    http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        // Logger contextual
        reqLog := log.WithContext(map[string]interface{}{
            "method": r.Method,
            "path": r.URL.Path,
            "remote": r.RemoteAddr,
        })
        
        reqLog.Info("Recibida solicitud")
        
        // Verificar método
        if r.Method != "GET" {
            w.WriteHeader(http.StatusMethodNotAllowed)
            fmt.Fprintf(w, "Método no permitido")
            reqLog.Warn("Método no permitido: %s", r.Method)
            return
        }
        
        // Obtener datos
        data, err := fetchData()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "Error interno del servidor")
            reqLog.Error("Error al obtener datos: %v", err)
            return
        }
        
        // Convertir a JSON
        jsonData := convert.StructToJSON(data)
        
        // Enviar respuesta
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(jsonData))
        
        reqLog.Info("Solicitud procesada correctamente")
    })
    
    // Iniciar servidor
    serverAddr := fmt.Sprintf(":%d", config.Port)
    log.Info("Iniciando servidor en %s", serverAddr)
    return http.ListenAndServe(serverAddr, nil)
}
```

## Patrones de Implementación

### Patrón Repository

```go
// Definir interfaz
type ItemRepository interface {
    Find(id string) (*Item, error)
    FindAll() ([]*Item, error)
    Save(item *Item) error
    Delete(id string) error
}

// Implementación con archivos
type FileItemRepository struct {
    basePath string
}

func NewFileItemRepository(basePath string) ItemRepository {
    return &FileItemRepository{basePath: basePath}
}

func (r *FileItemRepository) Find(id string) (*Item, error) {
    filePath := fmt.Sprintf("%s/%s.json", r.basePath, id)
    
    if !files.Exists(filePath) {
        return nil, fmt.Errorf("item no encontrado: %s", id)
    }
    
    data, err := files.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    
    var item Item
    err = convert.JSONToStruct(string(data), &item)
    if err != nil {
        return nil, err
    }
    
    return &item, nil
}

// Implementar los demás métodos...
```

### Patrón Service

```go
// Definir servicio
type ItemService struct {
    repo ItemRepository
    log  *logger.Logger
}

func NewItemService(repo ItemRepository) *ItemService {
    return &ItemService{
        repo: repo,
        log:  logger.New("item-service"),
    }
}

func (s *ItemService) ProcessItem(id string) (*ProcessedItem, error) {
    s.log.Info("Procesando item: %s", id)
    
    // Obtener item
    item, err := s.repo.Find(id)
    if err != nil {
        s.log.Error("Error al obtener item: %v", err)
        return nil, err
    }
    
    // Procesar
    result := &ProcessedItem{
        ID:        item.ID,
        Name:      item.Name,
        Processed: true,
        Value:     item.Value * 1.1, // Ejemplo de procesamiento
        Timestamp: time.Now(),
    }
    
    s.log.Info("Item procesado correctamente: %s", id)
    return result, nil
}
```

## Mejores Prácticas

1. **Manejo Uniforme de Errores**
   - Usar `fmt.Errorf` con `%w` para envolver errores
   - Proporcionar contexto útil en mensajes de error
   - Verificar tipos específicos de error para manejarlos adecuadamente

2. **Logging Efectivo**
   - Incluir contexto relevante en logs
   - Usar niveles apropiados (Debug, Info, Warn, Error)
   - No exponer información sensible en logs

3. **Configuración**
   - Usar valores predeterminados razonables
   - Permitir configuración desde múltiples fuentes
   - Validar configuraciones al inicio

4. **Rendimiento**
   - Usar buffers y pools para operaciones frecuentes
   - Implementar concurrencia para operaciones independientes
   - Monitorear y optimizar puntos críticos

5. **Seguridad**
   - Sanitizar entradas de usuario
   - Usar funciones criptográficas actualizadas
   - Manejar datos sensibles con precaución