# eval
A simple implementation of Shunting Yard.

## Usage
Given an infix expression, `eval.Eval(expression)` will evaluate the expression using Shunting Yard. The infix expression is converted to Reverse Polish Notation, and then evaluated.

```go
package main

import "github.com/aboxofsox/eval"

func main() {
        exp := "4 * 1024 + 1024" // 5120
        n, err := eval.Eval(exp)
        if err != nil {
                panic(err)
        }
        fmt.Println(n) // 5120
}
```