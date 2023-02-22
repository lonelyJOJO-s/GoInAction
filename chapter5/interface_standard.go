package chapter5

import (
	// "fmt"
	// "io"
	// "net/http"
	// "os"

	"fmt"
)

// func init() {
// 	if len(os.Args) != 2 { // notcie that system have a param already
// 		fmt.Println("...")
// 		os.Exit(-1)
// 	}
// }

// func main() {
// 	r, err := http.Get(os.Args[1]) // return a http.response pointer
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	io.Copy(os.Stdout, r.Body) // r.Body 是一个io.ReadCloser接口类型的值
// 	if err := r.Body.Close(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func main(){
// 	var b bytes.Buffer

// 	b.Write([]byte("Hello"))

// 	fmt.Fprintf(&b, "world!")

// 	io.Copy(os.Stdout, &b)

// }

// type user struct {
// 	name  string
// 	email string
// }

// type notifier interface {
// 	notify()
// }

// func (u *user) notify() {
// 	fmt.Println("name: "+u.name, "email: "+u.email)
// }

// func main() {
// 	bill := user{"bill", "sadsad"}
// 	sendNotification(&bill) // 注意此处不能把user对象bill作为接收者
// }

// func sendNotification(n notifier) {
// 	n.notify()
// }

// type duration int

// func (d *duration) pretty() {
// 	fmt.Print(d)
// }
// func mian(){
// 	a := 22
// 	fmt.Print(&a)

// }

type user struct {
	name     string
	password string
}

func (u *user) notify() {
	fmt.Print(u.name, u.password)
}

type admin struct {
	user  // 注意此处的写法， user user的写法的话不是内部置入，而是拥有一个其对象
	level string
}

func main() {
	ad := admin{
		user: user{
			name:     "lisa",
			password: "123",
		},
		level: "haha",
	}
	ad.user.notify() // ad为user类的超类， user实现的接口ad也自动继承实现
	ad.notify()      // 在外部类型仍然可以访问
	ad.name = "lll"
}
