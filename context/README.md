# Context 내부 속으로 (1)
- 최근 이미지 업로드 / 다운로드 서버를 개발하면서 고루틴을 많이 다루게 되었다.
- 이 기회에 고루틴 및 context에 개념을 정리하고 가고자 한다.
- 잘 모르면 주석부터 해석해보는 것이 진리 아니겠는가
- 1부는 패키지에 소개된 주석들을 살펴보는 것이다.

<br>
<br>

## context란?
- 패키지의 주석들을 해석해서 살펴보자.
```
Package context defines the Context type, which carries deadlines,
cancellation signals, and other request-scoped values across API boundaries and between processes.
```
```
- context 패키지는 Context라는 타입으로 정의할 수 있는데
  - 마감시간
  - 취소 신호
  - API 경계를 넘어서 프로세스 간에 요청 범위의 값
들을 전달할 수 있다.
```
- 즉 `값을 전달`할 수 있는 일종의 통로인가보다.
- 다른 블로그들의 글을 인용하면 `맥락을 유지하는 통로` 라고 표현을 볼 수 있다.
  - 현재 맥락(흐름) 안에서 유지해야할 값을 context에 담아 전달하고, 필요한 곳에서 context에서 값을 꺼내 사용할 수 있다고 이해하자.

<br>
<br>
<br>


```
Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. 
The chain of function calls between them must propagate the Context, optionally replacing 
it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. 
When a Context is canceled, all Contexts derived from it are also canceled.
```
```
1. 서버에 들어오는 요청은 컨텍스트를 생성해야 하고, 서버에서 나가는 호출은 컨텍스트를 수락해야 합니다.
2. 서버로 들어와서 나가기까지 함수 호출 체이닝간 컨텍스트를 전파해야하며, 선택적으로 `WithCancel`, `WithDeadline`, `WithTimeout`, `WithValue`를 사용하여 생성퇸 컨텍스트에서 파생된 컨텍스트를 사용한다.
3. 컨텍스트가 취소되면 해당 컨텍스트에서 파생된 모든 컨텍스트도 취소가 된다.
```

요약을 해보면 
1. 서버에서 나가는 호출은 컨텍스트를 수락해야한다 => in-out으로 이해가 되어 호출의 종료를 말하는 것 같다.
  - 즉, 컨텍스트를 종료시켜야한다고 이해가 된다.
2. 서버로 들어온 요청이 나가기 전까지 연쇄로 호출되는 함수간에 컨텍스트를 파라미터 등의 방법을 통해 전달해야한다.
   - 또한 위 4개의 함수를 이용해서 선택적으로 컨텍스트를 파생시킬 수 있는 것 같다.
   - 컨텍스트를 파생한다는 것은 새로운 컨텍스트를 생성할 수 있는데 부모 - 자식 관계라고 이해가 된다.
   - 게다가 위 4개 함수는 모두 자식 컨텍스트를 생성하는 함수들인 것 같다.
3. 부모 컨텍스트가 취소되면 해당 부모의 자식 컨텍스트들도 모두 취소가 되는 것 같다.

<br>
<br>
<br>

```
1. The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent)
and return a derived Context (the child) and a CancelFunc.
2. Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers.
3. Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires. 
4. The go vet tool checks that CancelFuncs are used on all control-flow paths.
```
```
1. WithCancel, WithDeadline 및 WithTimeout 기능은 컨텍스트(상위)를 사용합니다. 파생 컨텍스트(하위)와 CancelFunc를 반환합니다.
2. CancelFunc를 호출하면 자식 컨테스트와 해당 자식 컨텍스트의 자식 컨텍스트(?)가 취소되고 자식 컨텍스트에 대한 부모 컨텍스트와의 참조가 제거되며 연결된 타이머가 중지됩니다.
3. CancelFunc 호출을 실패하면 부모 컨텍스트가 취소되거나 타이머가 실행될 때까지 자식 컨테스트와 해당 자식 컨텍스트의 자식 컨텍스트(?)가 leak 된다.
4. $ go vet 는 CancelFuncs가 모든 제어 흐름 경로에 사용되는지 확인한다.
```
요약하면,
1. 위에서 이해한대로 WithCancel, WithDeadline 및 WithTimeout은 부모 컨텍스트 기준으로 자식 컨텍스트를 생성한다.
   - 대신 반환 값으로 cancelFunc라는 취소와 관련된 함수를 반환하는 것 같다.
2. 반환된 취소 함수를 실행하면 부모의 자식들이 모두 취소가 되는 것 같다. 대신 부모는 취소가되는지는 별도로 2부에서 확인이 필요해보인다.
   - 모든 자식 컨텍스트들은 부모의 참조가 끊킨다.
   - 타이머가 중단된다는데, 이 부분도 2부에서 확인해보자.
3. 반환된 취소 함수를 실행했는데 실패할 수도 있는 것 같다. 이때, 자식 컨텍스트들이 leak이 된다는데 메모리 누수인지, 그냥 노출이 된다는 건지는 모르겠다.
4. go vet라는 표준 패키지로써, 코드 정적 분석 툴인데 해당 툴이 함수 호출하면서 작업 플로우간 CancelFuncs가 잘 사용되었는지 확인하는 것 같다.


<br>
<br>
<br>


```
Programs that use Contexts should follow these rules 
to keep interfaces consistent across packages and 
enable static analysis tools to check context
propagation:
```
```
1. Context를 사용하는 프로그램은 다음 룰들을 따라야 한다.
패키지 간에 인터페이스를 일관되게 유지하고 컨텍스트 전파를 확인하기 위해 정적 분석 툴을 활성화해야한다.
```
- 패키지 간 인터페이스를 일관되게 유지한다는 것은 일관된 행동을 할 수 있도록 유지하라는 것으로 이해가된다.
- 다음 주석도 살펴보자.

<br>
<br>
<br>

```
Do not store Contexts inside a struct type; instead, pass a Context
explicitly to each function that needs it. The Context should be the first
parameter, typically named ctx:
```
```
```

<br>
<br>
<br>

```go
func DoSomething(ctx context.Context, arg Arg) error {
    // ... use ctx ...
}
```
```
Do not pass a nil Context, even if a function permits it. 
Pass context.TODO if you are unsure about which Context to use.
```
```
함수가 허락할지라도, Context에 nil을 전달하지 말 것.
만약 사용할 Context가  확실하지 않는다면, Context.TODO를 전달해라
```
- 말 그대로 Context를 전달하면서 nil을 전달하지말고 어떤 Context를 전달해야할지 모르겠다면
- `Context.TODO`를 전달하라고 한다. 꿀팁 ㄱㅅ

<br>
<br>
<br>

```
Use context Values only for request-scoped data that transits processes and
APIs, not for passing optional parameters to functions.
```
```
오직 프로세스를 전송하는 요청 범위 내의 데이터만 Conext 값들을 사용하고
Conext는 함수 파라미터에 값을 전달하기 위한 API가 아닙니다.
```
- Conext를 함수에 값을 전달하는 파라미터 용도로만 활용하지 말라는 뜻으로 이해가 되며,
- A라는 Proccess가 있을 때 B라는 Process로 흐름을 전파시킬 때, 사용하라는 뜻인가? 싶다.
  - 아니면 Proccess 내에서 제어할 수 있는 혹은 유효한 데이터만 Context를 이용해서 사용하라는 건가?

<br>
<br>
<br>

## Context의 func
- 한번 생성된 Context는 변경할 수 없다.
- 컨텍스트에 값을 추가하고 싶을 때는 context.WithValue() 함수로 `새로운 컨텍스트`를 만들어줘야한다.
- 컨텍스트의 값을 가져올 때는 Value 메서드를 사용한다.
- `context.WithCancel` 함수로 생성한 컨텍스트는 취소 신호를 보낼 수 있다.
- 일정시간이 지나면 자동으로 컨텍스트 취소 신호가 전달되도록 하려면 `context.WithDeadline` or `context.WithTimeout` 함수를 사용하여 컨텍스트를 생성하면 된다.