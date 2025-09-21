# Hello Go API

API REST desarrollada en Go usando arquitectura hexagonal con Gin como framework web y MySQL como base de datos.

## 📋 Requisitos Previos

- **Go 1.25.1** o superior
- **MySQL 5.7** o superior
- **Git**

## 🚀 Setup del Proyecto

### 1. Clonar el Repositorio
```bash
git clone <url-del-repositorio>
cd hello-go
```

### 2. Instalar Dependencias
```bash
go mod download
```

### 3. Configurar Base de Datos

#### Crear la Base de Datos
```sql
CREATE DATABASE `hello-go`;
```

#### Configuración de Conexión
La aplicación se conecta a MySQL con los siguientes parámetros por defecto:
- **Host**: localhost
- **Puerto**: 3306
- **Usuario**: root
- **Contraseña**: root
- **Base de datos**: hello-go

> **Nota**: Para cambiar estos valores, modifica las constantes en `cmd/api/bootstrap/bootstrap.go`

### 4. Crear Tablas (si es necesario)
```sql
USE `hello-go`;

-- Ejemplo de tabla para cursos
CREATE TABLE courses (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    duration VARCHAR(100) NOT NULL
);
```

## 🏃‍♂️ Ejecutar la Aplicación

### Desarrollo
```bash
go run cmd/api/main.go
```

### Compilar y Ejecutar
```bash
go build -o bin/hello-go cmd/api/main.go
./bin/hello-go
```

La aplicación se ejecutará en `http://localhost:8080`

## 🔧 Endpoints Disponibles

### Health Check
```http
GET /health
```

### Crear Curso
```http
POST /courses
Content-Type: application/json

{
    "id": "course-uuid",
    "name": "Curso de Go",
    "duration": "40h"
}
```

## 🧪 Testing

### Ejecutar Todos los Tests
```bash
go test ./...
```

### Ejecutar Tests con Verbose
```bash
go test -v ./...
```

### Ejecutar Tests de un Paquete Específico
```bash
go test ./internal/creating/
go test ./internal/platform/server/handler/courses/
```

### Ejecutar Tests con Coverage
```bash
go test -cover ./...
```

### Generar Reporte de Coverage HTML
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 🎭 Generación de Mocks

Este proyecto utiliza [Mockery](https://github.com/vektra/mockery) para generar mocks automáticamente.

### Instalar Mockery
```bash
go install github.com/vektra/mockery/v2@latest
```

### Generar Mocks

#### Opción 1: Usar go generate
```bash
go generate ./internal/courses.go
```

#### Opción 2: Usar mockery directamente
```bash
# Asegurar que mockery esté en el PATH
export PATH=$PATH:~/go/bin

# Generar todos los mocks según configuración
mockery --config .mockery.yaml

# Generar mock específico
mockery --case=snake --outpkg=storagemocks --output=internal/platform/storage/storagemocks --name=CourseRepository --dir=internal
```

### Configuración de Mocks

Los mocks se configuran en `.mockery.yaml` y se generan automáticamente en:
- `internal/platform/storage/storagemocks/` - Mocks de repositorios

### Ejemplo de Uso en Tests
```go
func TestSomething(t *testing.T) {
    // Crear mock
    mockRepo := storagemocks.NewCourseRepository(t)
    
    // Configurar comportamiento esperado
    course := core.NewCourse("id", "name", "duration")
    mockRepo.On("Save", mock.Anything, course).Return(nil)
    
    // Usar el mock en tu test
    service := creating.NewCourseService(mockRepo)
    err := service.CreateCourse(context.Background(), "id", "name", "duration")
    
    // Las expectativas se verifican automáticamente
    assert.NoError(t, err)
}
```

## 📁 Estructura del Proyecto

```
hello-go/
├── cmd/
│   └── api/
│       ├── bootstrap/          # Configuración e inicialización
│       └── main.go            # Punto de entrada
├── internal/
│   ├── courses.go             # Entidades y contratos del dominio
│   ├── creating/              # Casos de uso de creación
│   └── platform/
│       ├── server/            # Servidor HTTP y handlers
│       │   └── handler/
│       │       ├── courses/   # Handlers de cursos
│       │       └── health/    # Handler de salud
│       └── storage/           # Implementaciones de persistencia
│           ├── mysql/         # Repositorio MySQL
│           └── storagemocks/  # Mocks generados
├── .mockery.yaml              # Configuración de mockery
├── go.mod                     # Dependencias
└── README.md                  # Este archivo
```

## 🛠️ Herramientas de Desarrollo

### Linting
```bash
# Instalar golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Ejecutar linter
golangci-lint run
```

### Formateo de Código
```bash
# Formatear todo el código
go fmt ./...

# Imports automáticos
goimports -w .
```

## 🔍 Comandos Útiles

### Ver Dependencias
```bash
go mod tidy
go mod graph
```

### Limpiar Módulos
```bash
go clean -modcache
```

### Ver Tests que Fallan
```bash
go test ./... | grep FAIL
```

### Benchmark
```bash
go test -bench=. ./...
```

## 🐛 Troubleshooting

### Error de Conexión a MySQL
1. Verificar que MySQL esté corriendo
2. Confirmar credenciales en `bootstrap.go`
3. Verificar que la base de datos existe

### Mocks No Se Generan
1. Verificar que mockery esté instalado: `which mockery`
2. Añadir `~/go/bin` al PATH: `export PATH=$PATH:~/go/bin`
3. Verificar configuración en `.mockery.yaml`

### Tests Fallan
1. Verificar que los mocks estén generados
2. Revisar imports en archivos de test
3. Ejecutar `go mod tidy` para resolver dependencias

## 📚 Recursos Adicionales

- [Documentación de Go](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [Testify](https://github.com/stretchr/testify)
- [Mockery](https://github.com/vektra/mockery)
- [Arquitectura Hexagonal](https://alistair.cockburn.us/hexagonal-architecture/)

## 🤝 Contribución

1. Fork el proyecto
2. Crear una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abrir un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.