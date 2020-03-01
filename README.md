# Ejercicio Mercadolibre
## Contenido
<!--ts-->
* [Ejecución local](#ejecución-local)
  * [Requisitos](#requisitos)
  * [Configuración](#configuración)
  * [Instrucciones](#instrucciones)
  * [Pruebas](#pruebas)
* [Uso](#uso)
  * [POST /mutant](#post-/mutant)
  * [GET /stats](#get-/stats)
<!--te-->

## Ejecución local
### Requisitos
* MongoDB
* Redis
### Configuración
La configuración se gestiona a través de variables de entorno para cada servicio requerido.

---
Variables para MongoDB y valores por defecto:
* MONGODB_CREDS_USER: Usuario de base de datos ("mutant")
* MONGODB_CREDS_PWD: Contraseña del usuario de base de datos ("mutant")
* MONGODB_SERVER_ADDR: Dirección del servidor de base de datos ("localhost")
* MONGODB_SERVER_PORT: Puerto del servidor de base de datos ("27017")
* MONGODB_DATABASE: Nombre de la base de datos ("mutantdb")


---
Variables para Redis:
* REDIS_SERVER_ADDR: 
* REDIS_CREDS_PWD: 


### Instrucciones
- Asignar variables de entorno requeridas
- Compilar y ejecutar main.go

### Pruebas
Para ejecutar las pruebas se requiere una instancia local de MongoDB y una de Redis, ambas sin autenticación y corriendo en su puerto por defecto (27017 para MongoDB y 6379 para Redis)

El método más sencillo consiste en ejecutar estos servicios con docker para mantenerlos efímeros y separados de cualquier otro ambiente.

#### MongoDB
```
  docker run --rm -p 27017:27017 -d mongo:4.2.3-bionic
```

#### Redis
```
  docker run --rm -p 6379:6379 -d redis:5.0.7-buster
```

## Uso
El programa responde a través de una interfaz HTTP REST a las siguientes solicitudes:

### POST /mutant
Este servicio permite verificar si una cadena de ADN específica pertenece a un mutante o a un humano.

#### Headers
Debe especificar el formato JSON del contenido
```
Content-Type: application/json
```

#### Body
Un objeto JSON con clave "dna" y como valor la cadena de ADN a verificar.

La cadena debe ser una lista de seis (6) cadenas de texto, cada una con seis (6) caracteres.

Los únicos caracteres válidos son **A T G C**

##### Ejemplo
```
{
    "dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```

#### Respuesta
La aplicación responde únicamente con un header con código de estado:
* HTTP 200 (OK): si se envió un ADN mutante
* HTTP 403 (Forbidden): si se envió un ADN humano


### GET /stats
Este servicio muestra las estadísticas de humanos y mutantes descubiertos a través de este programa. No recibe parámetros.

#### Respuesta
La aplicación responde con un código de estado HTTP 200 (OK) y un objeto JSON con las estadisticas organizadas en los siguientes atributos:

* count_mutant_dna: Cantidad de secuencias de ADN mutante examinadas
* count_human_dna: Cantidad de secuencias de ADN humano examinadas
* ratio: Relación entre la cantidad de secuencias de ADN mutantes y humanas examinadas

##### Ejemplo
```
{
    "count_mutant_dna": 40,
    "count_human_dna": 100,
    "ratio": 0.4
}
```