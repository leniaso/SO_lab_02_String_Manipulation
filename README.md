# SO_lab_02_String_Manipulation

Este documento explica las funciones que ya tenemos hechas, para que sea fácil
retomar el trabajo desde acá. La tarea pendiente es el **conteo de vocales y
consonantes**.
 
## Idea general del patrón usado
 
En Go los strings son inmutables, así que no podemos modificarlos directamente
carácter por carácter. Por eso, el flujo del programa es:
 
1. Convertir el string de entrada a `[]rune` (slice de runas).
2. Trabajar sobre ese slice usando **punteros** (`&runas[i]` para obtener la
   dirección, `*puntero` para leer/escribir el valor ahí guardado).
3. Al final, convertir el slice de vuelta a `string` para imprimir.
### ¿Por qué `rune` y no `byte`?
 
`byte` ocupa 1 byte y representa bien el ASCII, pero las vocales con tilde
(á, é, í, ó, ú) ocupan **2 bytes** en UTF-8, así que un solo `byte` no las
representa correctamente. `rune` (alias de `int32`) sí representa un carácter
completo sin importar cuántos bytes use internamente. Por eso todo el
programa usa `[]rune`, no `[]byte`.
 
### El patrón "puntero de lectura / puntero de escritura"
 
Varias funciones recorren el slice con dos índices:
 
- `idxLectura`: avanza **siempre**, en cada vuelta del `for`. Es el que
  recorre todo el slice de principio a fin.
- `idxEscritura`: solo avanza **cuando se necesita escribir algo**. Se usa
  para "comprimir" el slice cuando estamos descartando caracteres.
Regla clave: `idxEscritura` nunca puede ser mayor que `idxLectura`. Por eso
es seguro escribir en `runas[idxEscritura]` sin perder datos que aún no se
han leído: siempre se está pisando una posición ya procesada.
 
En funciones donde **no se comprime nada** (como reemplazar espacios por
guion bajo, que no cambia el tamaño del string), no hace falta un segundo
índice — basta con el puntero de lectura, ya que se escribe en la misma
posición que se lee.
 
---
 
## Funciones ya implementadas
 
### `getString() string`
 
Lee el argumento de línea de comandos (`os.Args[1]`), valida que no supere
los 100 caracteres. Si es inválido, retorna `""`.
 
### `stringToRune(s string) []rune`
 
Convierte el `string` de entrada a `[]rune`, para poder manipular cada
carácter (incluidas tildes) como una sola posición del slice.
 
### `filtrarCaracteres(runas []rune) []rune`
 
**Qué hace:** elimina del slice todo lo que no sea letra (a-z, A-Z),
vocal con tilde (áéíóú, ÁÉÍÓÚ) o espacio. Números, signos de puntuación,
etc. quedan fuera.
 
**Cómo:** usa el patrón lectura/escritura. Recorre con `idxLectura`; cada
vez que encuentra un carácter válido, lo copia a `runas[idxEscritura]` (vía
puntero) y solo entonces avanza `idxEscritura`. Al final devuelve
`runas[:idxEscritura]`, recortando la "basura" que quedó al final.
 
**Importante:** esta función debe llamarse **primero**, antes que cualquier
otra, porque las demás asumen que el string ya está "limpio" (solo letras
y espacios), tal como lo exige el enunciado.
 
### `revertirCaracteres(runas []rune) []rune`
 
**Qué hace:** invierte el orden del slice, in place (sin usar memoria
adicional para otro string).
 
**Cómo:** dos punteros que caminan uno hacia el otro — `idxInicio` desde
el principio, `idxFin` desde el final. En cada vuelta se intercambian los
valores (`*ptrInicio, *ptrFin = *ptrFin, *ptrInicio`) y los índices se
acercan (`idxInicio++`, `idxFin--`) hasta cruzarse.
 
**⚠️ Cuidado — efecto secundario importante:** como los slices en Go
apuntan a la misma memoria subyacente, esta función modifica el slice
**original**, no una copia. Si le pasas `entrada_rune` y luego sigues
usando `entrada_rune` más adelante en el programa, ya vendrá reversado.
Por eso, si se necesita el string reversado *y también* el original o
modificado de otra forma, hay que sacar una copia antes de llamar a esta
función (con `make` + `copy`), y pasarle la copia a `revertirCaracteres`.
 
### `cambiarEspacios(runas []rune) []rune`
 
**Qué hace:** reemplaza cada espacio `' '` por guion bajo `'_'`.
 
**Cómo:** como reemplazar no cambia el tamaño del string, no hace falta
"comprimir" nada — se recorre con un solo índice (`idxLectura`), y cuando
se encuentra un espacio, se escribe `'_'` directamente en esa misma
posición vía puntero.
 
---
 
## Pendiente: conteo de vocales y consonantes
 
Falta una función que reciba el `[]rune` ya filtrado (antes de reversar o
cambiar espacios) y devuelva:
 
- Cantidad de vocales (a, e, i, o, u — incluyendo con tilde, mayúsculas y
  minúsculas cuentan igual)
- Cantidad de cada vocal por separado, en orden: a, e, i, o, u
- Cantidad de consonantes
**Sugerencia de enfoque:** seguir el mismo patrón de recorrido con un solo
puntero (similar a `cambiarEspacios`, ya que no se modifica el slice, solo
se cuenta), con contadores separados para cada vocal y uno para
consonantes. Ojo con las vocales acentuadas: hay que normalizarlas (por
ejemplo, contar `'á'` como una "a") o llevar contadores que sumen tanto la
versión con tilde como sin tilde en el mismo total.
 
---
 
## Orden correcto de las operaciones en `main`
 
Recordar el flujo pedido por el enunciado, todo a partir del string ya
filtrado:
 
1. `filtrarCaracteres` (una sola vez, al inicio)
2. A partir de ahí, generar **copias independientes** para:
   - Reversar (`revertirCaracteres`)
   - Contar vocales/consonantes (función pendiente, no modifica nada)
   - Cambiar espacios por `_` (`cambiarEspacios`)
3. Armar el string de salida final, separado por espacios, como pide el
   enunciado:
```
   <reversado> <num_vocales> <a> <e> <i> <o> <u> <consonantes> <con_guiones>
```
