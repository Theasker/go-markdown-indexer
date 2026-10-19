# Markdown indexer in Go

Herramienta de línea de comandos (CLI) escrita en Go que genera de forma automática un archivo `index.md` con la tabla de contenidos jerárquica de cualquier estructura de carpetas con documentación Markdown.

## Características

- **Recorrido recursivo**: Inspecciona directorios y subdirectorios a cualquier nivel de profundidad.
- **Extracción automática de títulos**:
  - Si el archivo contiene un encabezado `# Título (H1)`, se utiliza como texto del enlace.
  - Si no contiene ningún H1, utiliza el nombre del archivo sin extensión.
- **Estructura en árbol fiel**: Representa la jerarquía real de carpetas utilizando listas anidadas en Markdown.
- **Rutas normalizadas**: Enlaces relativos limpios utilizando siempre separadores `/`, compatibles con cualquier sistema operativo, GitHub, GitLab y visores web.
- **Prevención de bucles**: Descarta automáticamente el propio `index.md` raíz si ya existe para evitar autorreferencias.
- **Sin dependencias externas en producción**: Desarrollado utilizando exclusivamente la librería estándar de Go.

---

## Requisitos

- [Go](https://go.dev/dl/) 1.20 o superior instalado en el sistema.

---

## Instalación y compilación

Puedes compilar la herramienta para generar un binario ejecutable independiente:

```bash
# Clonar o entrar al repositorio
cd md-indexer-go

# Compilar el binario
go build -o md-indexer .
# Linux (64 bits)
GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64-md-indexer .

# Windows (64 bits)
GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64-md-indexer.exe .

# macOS (Apple Silicon / ARM64 - usar amd64 para Intel)
GOOS=darwin GOARCH=arm64 go build -o bin/app-darwin-arm64-md-indexer .

# Android (64 bits ARM)
GOOS=android GOARCH=arm64 go build -o bin/app-android-arm64-md-indexer .
```

Esto generará el ejecutable `md-indexer` (o `md-indexer.exe` en Windows) en la raíz del proyecto.

---

## Modo de uso

Ejecuta el binario pasando como argumento la ruta del directorio que deseas indexar:

```bash
# Si usas el binario compilado:
./md-indexer ./docs

# O ejecutándolo directamente con Go:
go run . ./docs
```

El programa escaneará el directorio indicado y creará o actualizará el archivo `index.md` directamente en la raíz de esa carpeta.

---

## Ejemplo de funcionamiento

Dada la siguiente estructura de carpetas:

```text
docs/
├── README.md                  # Contiene: "# Mi Documentación"
├── guias/
│   ├── instalacion.md         # Contiene: "# Guía de Instalación"
│   └── configuracion.md       # Sin H1
└── notas/
    └── reunion.md             # Contiene: "# Notas de Reunión"
```

Al ejecutar `md-indexer ./docs`, se generará el archivo `docs/index.md` con este contenido:

```markdown
- [Mi Documentación](README.md)
- **guias**
  - [Guía de Instalación](guias/instalacion.md)
  - [configuracion](guias/configuracion.md)
- **notas**
  - [Notas de Reunión](notas/reunion.md)
```

---

## Arquitectura del código

El proyecto está diseñado de forma modular dentro del paquete `main`, separando responsabilidades en archivos independientes:

- `main.go`: Punto de entrada de la CLI, validación de argumentos y orquestación del flujo principal.
- `1.markdown.go`: Recorrido recursivo (`filepath.WalkDir`), filtrado de extensiones y extracción eficiente de títulos H1 con `bufio.Scanner`.
- `2.tree.go`: Estructura de datos jerárquica en memoria (`Node`) y construcción del árbol mediante navegación por cursor y punteros.
- `3.generador.go`: Renderizado recursivo del árbol a formato Markdown con listas anidadas utilizando `strings.Builder`.
- `4.output.go`: Capa de persistencia encargada de escribir de forma segura el archivo final `index.md` en disco mediante `os.WriteFile`.

---

## Licencia

MIT
