# Nivel 7: Creemos un mini juego - desplazamiento del personaje

¡Bienvenid@! Este repositorio acompaña al tutorial "Creemos un mini juego - desplazamiento del personaje" de nuestro canal de YouTube, donde aprendemos a crear un juego básico utilizando Go y la biblioteca Ebitengine.

## 📋 Descripción

Este proyecto implementa un juego simple donde un personaje puede moverse, cambiar de tamaño y ajustar su velocidad. A través de este proyecto, aprenderemos conceptos fundamentales de Go como estructuras, métodos, manejo de errores, y cómo integrar una biblioteca externa para crear aplicaciones gráficas.

## 🚀 Comenzando

Para utilizar este juego, asegúrate de tener Go instalado en tu sistema.

### Prerrequisitos

* Go (versión 1.16 o superior recomendada)
* Ebitengine

### Instalación

1. Haz un fork de este repositorio haciendo clic en el botón "Fork" en la parte superior derecha de esta página.

Clona tu fork a tu máquina local:

git clone https://github.com/TU_USUARIO/gognition-nivel7-desplazamiento-personaje

2. Instala Ebitengine:

```
go get github.com/hajimehoshi/ebiten/v2
go get github.com/hajimehoshi/ebiten/v2/ebitenutil
```

Para más detalles sobre la instalación de Ebitengine, visita la [documentación oficial de instalación de Ebitengine](https://ebitengine.org/en/documents/install.html).

## 💻️ Uso

Para ejecutar el juego:

```
go run main.go
```

Controles del juego:
- Flechas: Mover el personaje
- Q/E: Disminuir/Aumentar el tamaño del personaje
- A/D: Disminuir/Aumentar la velocidad del personaje

## 🏗️ Compilación (Build)

Para compilar el proyecto y crear un ejecutable:

```
go build -o juego-simple
```

Para Windows (desde Linux o macOS):

```
GOOS=windows GOARCH=amd64 go build -o juego-simple.exe
```

Para macOS (desde Windows o Linux):

```
GOOS=darwin GOARCH=amd64 go build -o juego-simple-mac
```

Para Linux (desde Windows o macOS):

```
GOOS=linux GOARCH=amd64 go build -o juego-simple-linux
```

Después de la compilación, puedes ejecutar el programa usando:

```
./gopher-walk  # En Linux o macOS
gopher-walk.exe  # En Windows
```

## 🧠 Conceptos de Go aprendidos

1. Declaración de paquetes e importaciones
2. Uso de constantes y variables
3. Definición y uso de estructuras (structs)
4. Métodos en Go
5. Manejo de errores
6. Funciones y control de flujo
7. Operaciones matemáticas y lógica condicional
8. Ciclos y animación básica
9. Manejo de entrada de usuario
10. Renderizado gráfico básico

## 🔗 Enlaces útiles

- [Documentación oficial de Go](https://golang.org/doc/)
- [Documentación oficial de Ebitengine](https://ebitengine.org/en/)
- [Canal de YouTube](https://www.youtube.com/@harleyzapata)

## 🤝 Contribuir

Las contribuciones son bienvenidas. Si tienes alguna idea para mejorar el juego o el tutorial, no dudes en crear un pull request o abrir un issue.

## 📝 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo `LICENSE` para más detalles.

---

Visita [gognition.pro](https://www.gognition.pro/) para más tutoriales y recursos sobre Go.
