# Go Introduction 2

## 출력 함수

* * *

Go 는 텍스트를 출력하기 위한 세가지의 함수가 있음

* `Print()`
* `Println()`
* `Printf()`

* * *

### `Print()` 함수

`Print()` 함수는 기본 포멧으로 전달된 인자를 출력함

```go
package main
import ("fmt")

func main() {
    var i, j string = "Hello", "World"

    fmt.Print(i)
    fmt.Print(j)
}
```

결과:

```bash
HelloWorld
```

* * *

만약 인자들을 새로운 줄에 출력하고 싶다면 `\n` 을 사용함

```go
package main
import ("fmt")

func main() {
    var i, j string = "Hello", "World"

    fmt.Print(i, "\n")
    fmt.Print(j, "\n")
}
```

결과:

```bash
Hello
World
```

* * *

한개의 `Print()` 함수로 여러 변수를 출력하는 것 또한 가능함

```go
package main
import ("fmt")

func main() {
    var i, j string = "Hello", "World"

    fmt.Print(i, "\n", j)
}
```

결과:

```bash
Hello
World
```

* * *

만약 두 문자열 매개변수에 공백을 넣고 싶다면 `" "` 를 사용함

```go
package main
import ("fmt")

func main() {
    var i, j string = "Hello", "World"

    fmt.Print(i, " ", j)
}
```

결과:

```bash
Hello World
```

* * *

`Print()` 는 매개변수가 전부 문자열이 아닌 경우 사이에 공백을 추가함

```go
package main
import ("fmt")

func main() {
    var i, j = 10, 20

    fmt.Print(i, j)
}
```

결과:

```bash
10 20
```

* * *

### `Println()` 함수

`Println()` 함수는 `Print()` 와 비슷하지만 매개변수 사이에 공백이 추가되고, 끝에 개행문자가 추가된다는 차이점이 존재함

```go
package main
import ("fmt")

func main() {
    var i, j string = "Hello", "World"

    fmt.Println(i, j)
}
```

결과:

```bash
Hello World
```

* * *

### `Printf()` 함수

`Printf()` 함수는 전달된 format 구절로 매개변수를 변환한 뒤 출력함

여기선 두 개의 format 구절을 사용함

* `%v` 는 매개변수의 **value** 를 출력하기 위해 사용됨
* `%T` 는 매개변수의 **type** 을 출력하기 위해 사용됨

```go
package main
import ("fmt")

func main() {
    var i string = "Hello"
    var j int = 15

    fmt.Printf("i 는 %v 를 가지고 있고 %T 타입임\n", i, i)
    fmt.Printf("j 는 %v 를 가지고 있고 %T 타입임\n", j, j)
}
```

결과:

```bash
i 는 Hello 를 가지고 있고 string 타입임
j 는 15 를 가지고 있고 int 타입임
```
