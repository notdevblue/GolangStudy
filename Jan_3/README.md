# Go Introduction

## Go 란?

* cross-platform, open source 프로그래밍 언어
* 고성능 애플리케이션 제작이 가능
* 빠르고, 정적 타입, 컴파일 언어이지만 동적 타입, 인터프리터 언어 처럼 사용이 가능
* 2007년에 구글에 속한 Robert Griesemer, Rob Pike, 그리고 Ken Thompson 에 의해 개발됨
* C++ 과 문법이 비슷함

* * *

### Go 는 어디에 사용되는가?

* 웹 개발 (server-side)
* 네트워크 기반 프로그램 개발
* 크로스 플렛폼 기업형 애플리케이션 개발
* 클라우드 네이티브 개발

* * *

### 왜 Go?

* 재밌고 배우기 쉬움
* 빠른 컴파일, 런타임을 가짐
* 동시성을 지원함
* 메모리 메니지먼트가 존재함
* 여러 플렛폼에서 돌아감

* * *

## Go 문법

* * *

* Go 에서는 모든 프로그램이 패키지의 일부가 됨.

  `package` 키워드를 통해 이를 정의함

```go
package main
// code...
```

* `{` 는 라인의 맨 처음에 올 수 없음

```go
package main

func main()
{
    // 에러
}

func main() {
    // 노 에러
}
```

* 주석은 C 개열 언어와 동일함

```go
package main
import ("fmt") // 출력을 위한 페키지

func main() {
    // fmt.Println("이거 실행 안됨")

    /*
    fmt.Println("이것들도")
    fmt.Println("실행")
    fmt.Println("안됨")
    */
}
```

* * *

## Go 변수

* * *

### 변수 타입

* `int`
  * 정수를 저장함 (154, -154)
* `float32`
  * 부동소수점을 저장함 (154.1, -154.1)
* `string`
  * 문자열을 저장함. `string` 값은 큰따옴표로 둘러싸임
* `bool`
  * `true`, `false` 두 상태를 가질 수 있는 값을 저장함

* * *

### 변수 선언

#### `var` 키워드

* `var` 키워드를 변수명, 타입과 같이 사용

```go
var variablename type = value
```

* `type` 또는 `value` 중 하나를 반드시 명시해야 함

* * *

#### `:=` 연산자

* `:=` 연산자를 값과 같이 사용

```go
variablename := value
```

* 이 경우, 변수의 타입은 추론됨
* `value` 를 전달하지 않고 사용할 수 없음

* * *

#### 변수 선언 시 초기값 전달

* 변수의 값이 처음부터 알고 있는 상태라면, 한줄로 정의할 수 있음

```go
package main
import ("fmt")

func main() {
    var student1 string = "Han" // 타입은 string
    var student2 = "Jan" // 타입은 추론됨
    x := 2 // 타입은 추론됨

    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(x)
}
```

* * *

#### 초기값 전달 없이 변수 선언

* Go 언어는 모든 변수를 초기화함

  그래서 초기값 없이 변수를 선언한 경우, 선언된 변수의 타입이 가진 기본값이 전달됨

```go
package main
import ("fmt")

func main() {
    var a string
    var b int
    var c bool

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
```

위의 예시에는 변수가 3개 있음

* `a`
* `b`
* `c`

이 변수들은 선언되었지만 초기값이 전달되지 않았지만,

위의 코드를 실행시키면 각각의 타입에 따른 기본값이 저장되있음을 알 수 있음

* `a` 는 `""` 임
* `b` 는 `0` 임
* `c` 는 `false` 임

* * *

#### 선언 후 값 할당

* 변수가 선언된 뒤 값을 할당하는 것이 가능함.

  이는 초기에 대입될 값을 알 수 없을 때 유용함

```go
package main
import ("fmt")

func main() {
    var student1 string
    student1 = "Han"
    fmt.Println(student1)
}
```

* * *

### `var` 와 `:=` 의 차이점

`var` 와 `:=` 는 서로 작은 차이점이 있음

| `var` | `:=` |
| ----- | ---- |
| 함수 **안**과 **밖**에서 사용이 가능함 | 함수 **안**에서만 사용이 가능함 |
| 변수 선언과 값 할당이 **단독으로 이루어질 수 있음** | 변수 선언과 값 할당은 **단독으로 이루어질 수 없음** (같은 줄에서 이뤄저야 함) |

#### *예시*

```go
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
    a = 1

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
```

위의 예시는 `var` 키워드를 이용해 함수 안과 밖에서 변수를 정의함

```go
package main
import ("fmt")

a := 1

func main() {
    fmt.Println(a)
}
```

위의 예시는 `:=` 키워드를 이용해 함수 밖에서 변수를 정의함

함수 밖에서 변수를 정의했기 때문에, 아래와 같은 에러가 발생함

`syntax error: non-declaration statement outside function body`

* * *

### 다중 변수 선언

* Go 언어는 같은 줄에서 복수의 변수를 선언할 수 있음

```go
package main
import ("fmt")

func main() {
    var a, b, c, d int = 1, 3, 5, 7

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
```

`type` 키워드를 사용한 경우, 같은 줄에서 한개의 타입만 선언할 수 있음

`type` 키워드가 없는 경우, 같은 줄에서 여러 타입을 선언할 수 있음

```go
package main
import ("fmt")

func main() {
    var a, b = 6, "Hello"
    c, d := 7, "World!"

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
```

* * *

#### 블록 안에 변수 선언

* 변수의 다중 선언을 블록으로 그룹지어 가독성을 높일 수 있음

```go
package main
import ("fmt")

func main() {
    var (
        a int
        b int = 1
        c string = "Hello"
    )

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
```

* * *

### 변수 명명 규칙

변수는 짧은 이름 (x, y) 또는 의미가 담긴 이름 (age, price, carname 등등) 을 가질 수 있음

* 변수 이름은 단어나 언더스코어 문자 (`_`) 로 시작되어야 함
* 변수 이름은 숫자로 시작할 수 없음
* 변수 이름은 영어 단어와 숫자, 그리고 언더스코어 문자만 들어갈 수 있음 (`a-z`, `A-Z`, `0-9`, 그리고 `_`)
* 변수 이름은 대소문자를 구분함 (age, Age, AGE 는 각각 다른 변수)
* 변수 이름 길이 제한은 없음
* 변수 이름은 공백을 포함할 수 없음
* 변수 이름은 Go 키워드일 수 없음

* * *

### 복수 단어 변수명

여러 단어를 포함한 변수 이름은 읽기 어려울 수 있음

이를 읽기 쉽게 하기 위한 여러 기술들이 존재함

* * *

#### camelCase

첫 단어를 제외한 각각 단어의 시작은 대문자임

```go
myVariableName = "Han"
```

* * *

#### PascalCase

각각 단어의 시작은 대문자임

```go
MyVariableName = "Han"
```

* * *

#### snake_case

각각의 단어는 언더스코어 문자로 나누어짐

```go
my_variable_name = "Han"
```

* * *

### 상수

만약 변수가 변경할 수 없는 상수값을 가져야 한다면 `const` 키워드를 사용함

`const` 키워드는 변수를 상수로 선언함. 이는 이 변수는 불변하고 읽기 전용이라는 것을 나타냄

```go
const CONSTNAME type = value
```

상수 변수는 선언 시 `value` 를 전달해 줘야 함

* * *

#### 상수 정의 예시

```go
package main
import ("fmt")

const PI = 3.14

func main() {
    fmt.Println(PI)
}
```

* * *

#### 상수 규칙

* 상수 명명 규칙은 변수 명명 규칙과 동일함
* 상수 이름은 모두 대문자로 작성됨 (일반 변수와 차이를 쉽게 알기 위해)
* 상수는 함수 안과 밖에서 선언될 수 있음

* * *

#### 상수 타입

상수에는 두가지 타입이 있음

* Typed constants
* Untyped constants

* * *

#### Typed Constants

Typed Constants 는 정의된 타입을 이용해 정의함

```go
package main
import ("fmt")

const A int = 1

func main() {
    fmt.Println(A)
}
```

* * *

#### Untyped Contants

Untyped constants 는 타입 없이 정의함

```go
package main
import ("fmt")

const A = 1

func main() {
    fmt.Println(A)
}
```

이 상황에서는 상수의 타입은 전달된 값으로 추론됨

* * *

#### 상수: 불변성과 읽기 전용

상수가 정의된 경우, 정의된 후 값을 변경할 수 없음

```go
package main
import ("fmt")

func main() {
    const A = 1
    A = 2
    fmt.Println(A)
}
```

위는 아레와 같은 오류를 출력함

`cannot assign to A`

* * *

#### 다중 상수 선언

* 상수의 다중 선언을 블록으로 그룹지어 가독성을 높일 수 있음

```go
package main
import ("fmt")

const (
    A int = 1
    B = 3.14
    C = "Han"
)

func main() {
    fmt.Println(A)
    fmt.Println(B)
    fmt.Println(C)
}
```
