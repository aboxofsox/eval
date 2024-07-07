# eval
A simple implementation of Shunting Yard.

## Usage
Given an infix expression, `eval.Eval(expression)` will evaluate the expression using Shunting Yard. The infix expression is converted to Reverse Polish Notation, and then evaluated.

***Note**: because of how naive tokenization is, each token is expected to be separated by 1 space, including parentheses.*
```sh
"( 8 + 10 ) * 2" // okay
"(8 + 10) * 2" // not okay
```

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