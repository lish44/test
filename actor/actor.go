package actor

type Actor struct {
	mailbox chan interface{}
}

func (a *Actor) Send(msg interface{}) {
	a.mailbox <- msg
}

func (a *Actor) Receive() interface{} {
	return <-a.mailbox
}

func NewActor() *Actor {
	return &Actor{mailbox: make(chan interface{})}
}

func Hello() {
	println("Hello, World!")
}

var Person string = "John Doe"
