# GORoutines (ejecución Asíncrona - Promesas en GO)

En Go, cada unidad de ejecución concurrente se llama goroutina. La comunicación entre goroutines es a través de canales
```
go gorutina()
```
El siguiente ejemplo muestra que 2 Gorutines (main() y hola()):
```
package main

import (  
    "fmt"
    "time"
)

func hola() {  
    fmt.Println("Hola mundo goroutine")
}
func main() {  
    go hola()
    time.Sleep(1 * time.Second)
    fmt.Println("función main")
}
```
La función ```hello()``` se ejecuta concurrentemente con la función ```main()```. Si no pusieramos ```time.Sleep``` se mostraría solo  ```función main```. Esto ocurre porque:
- Cuando una gorutine empieza devuelve return inmediatamente, no espera a que la Gorutine se ejecute. El control se devuelve inmediatamente a la siguiente línea.
- La Gorutine ```main()``` debe estar ejecutándose para que se ejecuten otras Gorutines. Si la Gorutine main finaliza, el programa terminará y no se ejecutará ninguna otra Gorutine.

Ponemos ```time.Sleep(1 * time.Second)``` para que el programa espere un tiempo necesario para que termine la Gorutine ```hola()```
