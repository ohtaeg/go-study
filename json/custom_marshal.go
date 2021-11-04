package json

import (
	"encoding/json"
	"time"
)

var layout = "2006-01-02 15:04:05"

type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
}

// MarshalJSON stack overflow
// 내부 구조체는 "승격 과정"을 통해 외부 타입으로 승격한다.
// 승격된 내부 타입은 마치 외부에 직접 선언된 것 처럼 동작하기 때문에 MarshalJSON을 상속받아 무한 호출 된다.
/*
func (u *User) MarshalJSON() ([]byte, error) {
	converter := func(now time.Time) time.Time {
		parse, _ := time.Parse(layout, now.Format(layout))
		return parse
	}

	return json.Marshal(&struct {
		CreateAt time.Time `json:"createAt"`
		*User
	}{
		CreateAt: converter(u.CreateAt),
		User:   u,
	})
}
*/

// MarshalJSON Named Type 을 이용해서 무한 호출을 해결
// 어떤 블로그에서는 Type Alias 라고 하는데 개인적으로 Named Type 이라고 생각한다. (type MyUser = User 주석을 풀고 실행해볼 것)
// Named Type 은 기존 타입의 정의된 내용을 그대로 이용하면서 다른 새로운 타입으로 취급할 수 있다.
// 즉, Named type으로 선언한 새로운 타입과 원본이 된 타입은 호환성이 없기 때문에
// 서로 다른 타입이므로 명시적인 캐스팅을 통해 호환되도록 하고, 다른 타입의 MarshalJSON()을 호출함으로써 해결할 수 있다.
func (u *User) MarshalJSON() ([]byte, error) {
	// type MyUser = User
	type MyUser User
	converter := func(now time.Time) time.Time {
		parse, _ := time.Parse(layout, now.Format(layout))
		return parse
	}

	return json.Marshal(&struct {
		CreateAt time.Time `json:"createAt"`
		*MyUser
	}{
		CreateAt: converter(u.CreateAt),
		MyUser:   (*MyUser)(u),
		// *MyUser(u)  -> same as *(MyUser(u))
		//(*MyUser)(u) -> u is converted to *MyUser pointer
	})
}

//func main() {
//	_ = json.NewEncoder(os.Stdout).Encode(
//		&User{ID: 1, Name: "ohtaeg", CreateAt: time.Now()},
//	)
//}
