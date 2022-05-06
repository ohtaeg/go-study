# Context 내부 속으로 (2)
- 이번엔 내부적으로 어떻게 구성되어 있는지 그리고 주석을 생략하고 문서 기반으로 해석을 해보면서 살펴보자.

```go
package context

import (
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct{}

	Err() error

	Value(key any) any
}
```

- Context는 인터페이스로써 총 4개의 메서드가 정의되어 있으며 컨텍스트의 메서드는 여러 고루틴에 의해 동시에 호출될 수 있다.

<br>

## DeadLine()
- `Deadline() (deadline time.Time, ok bool)`
- 시간과 bool값을 반환하는데 
- 반환된 시간은 **컨텍스트로 취소 신호가 전달될 때까지 남은 시간**, 즉 Context가 끝나야할 시간이다.
- bool 값은 deadline이 있는지 여부를 표시한다.
  - false 일 경우 deadline이 없다는 것이며
  - true 일 경우 deadline이 존재한다는 것이다.
  - 작업을 시작하기 전 남은 시간을 먼저 확인을 해서, 충분한 시간이 있을 때만 작업을 수행하도록 할 수 있다.
- 이해가 쉽게 가지 않는다면 간단한 예제로 살펴보자
```go
package main

import (
	"context"
	"fmt"
	"time"
)

var format = "2006-01-02 15:04:05"

func main() {
  ctx := context.Background()
  fmt.Println("시작 전 : " + time.Now().Format(format))
  timeout, cancel := context.WithTimeout(ctx, 5*time.Second)
  deadline, ok := timeout.Deadline()
  fmt.Println("데드라인 : " + deadline.Format(format))
  fmt.Println(ok)

  cancel()

  time.Sleep(6 * time.Second)
  _, after := timeout.Deadline()
  fmt.Println(after)
}
```
- 해당 예제를 통해 5초후에 context에 timeout을 넘겨서 일정 시간이 되면 자동으로 컨텍스트에 취소 신호가 전달되도록하였다.
- 즉, 해당 context의 마감 시간 (5초)을 확인하는 것과 context 취소 후 ok 값이 변경되었는지 확인하는 예제이다.
- 실제로 실행해보면 다음과 같이 출력이 된다.
```
시작 전 : 2022-05-05 23:40:54
데드라인 : 2022-05-05 23:40:59
true
true
```
- 시작 전과 데드라인이 5초 차이가 있음을 확인하였다. 
- 다만 첫번째 ok 값은 데드라인이 있어서 true로 출력됨을 확인하였는데 6초후 after 값은 데드라인이 지난 시간이라 false를 예상해볼 순 있지만 true이다.
- 왜냐하면 한번 set 된 데드라인은 계속 호출하여도 동일한 결과를 반환한다고 문서에 설명되어있으니 확인해보면 좋다.
  ```text
  Deadline returns ok==false when no deadline is set.
  Successive calls to Deadline return the same results.
  ```

<br>
<br>

## Done()
- `Done() <-chan struct{}`
- Done() 메서드는 채널을 반환하는데 어떤 채널인지 문서를 읽어보자.
```
Done returns a channel that's closed when work done on behalf of this context should be canceled
```
- **Context가 취소 되거나 시간을 초과할 때 종료 신호를 전달받을 수 있는 close channel을 반환한다.**
  - 즉, cancel 함수를 실행하여 컨텍스트에 종료 신호를 보내면 종료 상황을 Done() 메서드를 통해 알 수 있다.

<br>

```
Done may return nil if this context can never be canceled.
```
- 취소되지 않은 컨텍스트의 경우 nil이 반환된다.

<br>

```
Successive calls to Done return the same value.
```
- DeadLine()과 마찬가지로 완료에 대한 연속 호출은 동일한 값을 반환한다.

<br>

```
The close of the Done channel may happen asynchronously, after the cancel function returns.
```
- Done()이 반환하는 채널의 종료는 cancel 함수가 반환된 후 비동기적으로 발생할 수 있다고 한다.

<br>

```
WithCancel arranges for Done to be closed when cancel is called;
WithDeadline arranges for Done to be closed when the deadline expires;
WithTimeout arranges for Done to be closed when the timeout elapses.
```
- WithCancel(), WithDeadline(), WithTimeout() 메서드들은 해당 함수들이 반환하는 cancel() 함수가 호출되었을 때 Done()의 반환 값인 종료 신호를 전달받을 수 있는 close channel이 닫히도록 처리한다.

<br>

```
Done() is provided for use in select statements
```
- Done()은 **select 문**에서 사용하기 위해 제공된다.
