# gRPC con Golang


* gRPC es un framework gratis y open-source desarrollado por Google. Forma parte de la Cloud Native Computing Foundation. A Alto nivel, permite definir peticiones y respuestas 
para RPC (Remote Procedure Calls) y maneja el resto por ti. 

* Es moderno, rápido y eficiente. Está construido sobre HTTP/2, tiene baja latencia, soporta streamings, es independiente del lenguaje de programación y es muy fácil incluir autenticación, balanceo de carga, logging y monitorización. 

* Los mensajes y servicios en gRPC se definen mediante Protocol Buffers, hay soporte para unos 12 lenguages de programación difenrtes. En estos tipos de mensaje los datos se almacenan en binario y serializados, por lo que el payload se reduce 

* Usa HTTP/2, que es más eficiente que su versión anterior. 

## Tipos de API gRPC

Existen cuatro tipos de API gRPC: unary, server streaming, client streaming y streaming bidireccional. 


La API de tipo **unary** es el tipo de API clásica basada en request/response. Una de tipo **Server streaming** hace que para una sola petición del cliente, el servidor pueda devolver varias respuestas. La opuesta a esta última es la de tipo **Client Streaming**, el cliente abre una conexión y envía todos los mensajes que quiera y el servidor responde una vez en algún momento. En la última, **Streaming bidreccional** tanto cliente como servidor envían varios mensajes, conviertiendo el proceso en algo asíncrono. 


```javascript

service GreetService {

    //Unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

    //Streaming Server
    rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesResponse){};

    //Streaming Client
    rpc LongGreet(stream longGreetRequest) returns (LongGreetResponse) {};

    // Bi directional streaming
    rpc GreetEveryone(stream GreetEveryoneRequest) returns (stream GreetEveryoneResponse) {};
}
```

## Escalabilidad

gRPC es extremandemente fácil de escalar. Por defecto, los servidores de gRPC son asíncronos, por lo que no bloquean los hilos con cada petición puediendo, así atender millones de petciciones en paralelo. 

El cliente de gRPC tiene la opción de ser síncrono o asíncrono, cada uno ofrece un rendimiento distintos. ***gRPC Clients can perform client side load balancing***.

## Seguridad

Es un framework seguro por defecto, haciendo que se tenga SSL por defecto. Además se pueden usar Interceptor para proporcionar a la API autenticación. 

## REST vs gRPC

| gRPC | REST |
|:----:|:----:|
|Protocol buffers - smaller,faster | JSON - text based, slower, bigger|
|HTTP/2 (lower latency) - from 2015 | HTTP1.1 (higher latency) - from 1997 |
| Bidirectional & Async | Client => Server requests only |
| Stream support | Request/Response support only |
|API Oriented - "What" (no constraints - free design) | CRUD Oriened |
| Code Generation through Protocol Buffers in any language | Code generation trough OpenAPI/Swagger (add-on)|
|RPC Based - gRPC does the plumbing for us | HTTP verbs based - we have to write the plumbing or use a use a third party library|

https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html


## Unary API

Este tipo de API supone una basada en petición/respuesta, de forma que el cliente envía una peticicón al servidor y este le responde con una respuesta. 

Son muy útiles cuando los datos a trasnmitir son pequeños. Se suele empezar con este tipo de APIs y se evolucionan a streaming cuando el rendimiento supone un problema. 


 ### Scripts

Variables de entorno necesarias:
```bash
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```


 Generar código en Go a partir de un protocol buffer (.proto): 
 ```bash
$ protoc -I . greet/greetpb/greet.proto --go_out=plugins=grpc:.
 ```

