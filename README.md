Este proyecto implementa un intérprete de expresiones matemáticas utilizando ANTLR4 para generar un analizador léxico y sintáctico. A continuación, se explica cómo está estructurado y cómo funciona:

### 1. **Definición de la gramática (`Calc.g4`)**
La gramática define las reglas para interpretar expresiones matemáticas simples. Contiene:
- **Tokens**: Define operadores (`*`, `/`, `+`, `-`), números (`NUMBER`) y espacios en blanco (`WHITESPACE`).
- **Reglas**:
    - `start`: Punto de entrada que espera una expresión seguida del final del archivo (`EOF`).
    - `expression`: Define operaciones de multiplicación/división (`MulDiv`), suma/resta (`AddSub`) y números (`Number`).

### 2. **Generación del código con ANTLR**
El archivo `Calc.g4` se utiliza para generar el código fuente necesario para el analizador léxico y sintáctico. Esto incluye:
- **Lexer**: Divide la entrada en tokens.
- **Parser**: Construye un árbol de análisis basado en las reglas de la gramática.
- **Listener**: Permite realizar acciones al entrar o salir de nodos del árbol.

### 3. **Implementación del Listener (`listener/calc_listener.go`)**
El listener personalizado (`calcListener`) extiende el listener base generado por ANTLR. Este se encarga de:
- **Pila (`stack`)**: Utiliza una pila para evaluar las expresiones.
- **Operaciones**:
    - `ExitMulDiv`: Realiza multiplicaciones o divisiones.
    - `ExitAddSub`: Realiza sumas o restas.
    - `ExitNumber`: Convierte un número del árbol en un entero y lo empuja a la pila.
- **Evaluación**: Al recorrer el árbol, las operaciones se realizan en el orden correcto gracias a la estructura del árbol.

### 4. **Función principal (`Calc`)**
La función `Calc` toma una expresión matemática como entrada, la analiza y devuelve el resultado:
1. Convierte la entrada en un flujo de tokens.
2. Crea un lexer y un parser basados en la gramática.
3. Construye el árbol de análisis y lo recorre con el listener personalizado.
4. Devuelve el resultado final desde la pila.

### 5. **Punto de entrada (`main.go`)**
El archivo `main.go` utiliza la función `Calc` para evaluar una expresión matemática y muestra el resultado en la consola.

### Ejemplo de ejecución
Si se evalúa la expresión `"3 * 5 + 2 * 10"`, el programa:
1. Construye el árbol de análisis.
2. Realiza las operaciones en el orden correcto (multiplicación antes de suma).
3. Devuelve el resultado `35`.

Este proyecto es un ejemplo práctico de cómo usar ANTLR4 para construir un intérprete o compilador básico, aprovechando la potencia de los analizadores léxicos y sintácticos generados automáticamente.