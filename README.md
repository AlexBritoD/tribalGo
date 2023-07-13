# Ejercicio induccion GO

Este proyecto consiste en una API en Go que centraliza varios servicios de búsqueda de canciones, consolidando los resultados, almacenándolos en una base de datos PostgreSQL y exponiendo un único endpoint para realizar búsquedas. La API requiere autenticación para ser consumida y se puede ejecutar utilizando Docker Compose con NGINX como servidor web.

## Instalación

Asegúrate de tener instalado Docker y Docker compose en tu sistema. Puedes obtenerlo en https://www.docker.com/.

Clona este repositorio en tu máquina local: https://github.com/AlexBritoD/tribalGo.git
Accede al directorio raíz del proyecto.
Ejecuta el siguiente comando para ejecutar el sistema desde Docker:

```bash
#Este bajara las imagenes necesarias para poder ejecutar el sistema en el puerto 8080
docker-compose up -d --build
```

## Usage

Una vez que la API está en funcionamiento, puedes realizar solicitudes a través del único endpoint disponible:

### Busqueda de canciones

```bash
GET /api/search?term=Nirvana
```

### Ejemplo de respuesta 
```json
{
    "success": true,
    "message": "Songs found",
    "data": [
        {
            "id": 518854672,
            "name": "Nirvana",
            "artist": "Adam Lambert",
            "duration": "4:22",
            "album": "Trespassing (Deluxe Version)",
            "artwork": "https://is2-ssl.mzstatic.com/image/thumb/Music124/v4/a3/aa/b1/a3aab1b1-b8fd-3139-8317-067b563c5552/886443458536.jpg/60x60bb.jpg",
            "price": "11.99 USD",
            "origin": "apple"
        },
        {
            "id": 1413921155,
            "name": "Smells Like Teen Spirit",
            "artist": "Nirvana",
            "duration": "5:01",
            "album": "Nirvana",
            "artwork": "https://is5-ssl.mzstatic.com/image/thumb/Music115/v4/7b/58/c2/7b58c21a-2b51-2bb2-e59a-9bb9b96ad8c3/00602567924166.rgb.jpg/60x60bb.jpg",
            "price": "9.99 USD",
            "origin": "apple"
        },
        {
            "id": 1413921156,
            "name": "Come As You Are",
            "artist": "Nirvana",
            "duration": "3:39",
            "album": "Nirvana",
            "artwork": "https://is5-ssl.mzstatic.com/image/thumb/Music115/v4/7b/58/c2/7b58c21a-2b51-2bb2-e59a-9bb9b96ad8c3/00602567924166.rgb.jpg/60x60bb.jpg",
            "price": "9.99 USD",
            "origin": "apple"
        }
    ]
}

```
