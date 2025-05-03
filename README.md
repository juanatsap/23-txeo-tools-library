# 23-txeo-tools-library

## Descripción General
Una biblioteca de utilidades y herramientas para desarrollo en Go, que proporciona funcionalidades comunes y reutilizables para una variedad de aplicaciones. Esta biblioteca centraliza código útil para operaciones frecuentes, permitiendo un desarrollo más rápido y consistente.

## Información del Repositorio
- **Proyecto ID**: 23
- **Repositorio**: `/Users/txeo/Git/go/23-txeo-tools-library`
- **Lenguaje Principal**: Go

## Características Principales
- Colección de utilidades para operaciones comunes
- Funciones de ayuda para manipulación de datos
- Herramientas para gestión de archivos y directorios
- Componentes para logging y depuración
- Utilidades para conversión y formateo de datos
- Herramientas para configuración y variables de entorno
- Funciones para operaciones asíncronas y concurrencia

## Paquetes Principales
- **files**: Utilidades para gestión de archivos y directorios
- **convert**: Herramientas para conversión entre tipos de datos
- **format**: Funciones para formateo de datos
- **env**: Gestión de variables de entorno y configuración
- **logger**: Sistema de logging extensible
- **time**: Utilidades para manejo de fechas y horas
- **crypto**: Funciones para operaciones criptográficas básicas
- **net**: Herramientas para operaciones de red
- **async**: Utilidades para concurrencia y operaciones asíncronas

## Estructura del Proyecto
```
23-txeo-tools-library/
├── doc/                  # Documentación
│   ├── technical/        # Documentación técnica
│   └── workflows/        # Ejemplos de uso y flujos de trabajo
├── files/                # Utilidades para gestión de archivos
├── convert/              # Herramientas de conversión
├── format/               # Funciones de formateo
├── env/                  # Gestión de entorno
├── logger/               # Sistema de logging
├── time/                 # Utilidades de fecha/hora
├── crypto/               # Funciones criptográficas
├── net/                  # Utilidades de red
└── async/                # Herramientas para concurrencia
```

## Instalación
```bash
go get github.com/txeo/tools-library
```

## Uso Básico
```go
package main

import (
    "fmt"
    "github.com/txeo/tools-library/files"
    "github.com/txeo/tools-library/logger"
    "github.com/txeo/tools-library/time"
)

func main() {
    // Inicializar logger
    log := logger.New("main")
    log.Info("Aplicación iniciada")
    
    // Utilizar funciones de tiempo
    now := time.Now()
    formattedDate := time.Format(now, "YYYY-MM-DD")
    log.Info("Fecha actual: %s", formattedDate)
    
    // Operaciones con archivos
    fileContent, err := files.ReadFile("config.json")
    if err != nil {
        log.Error("Error leyendo archivo: %v", err)
        return
    }
    
    // Procesar contenido
    fmt.Println("Contenido:", string(fileContent))
}
```

## Documentación Adicional
Para más información sobre los componentes y su uso, consulta:
- [Documentación Técnica](doc/technical/README.md)
- [Ejemplos y Flujos de Trabajo](doc/workflows/README.md)