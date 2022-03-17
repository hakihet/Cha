package Cha

type Cha struct {
	quit chan struct{}
	done chan struct{}
}

func Stop(stop ...Cha) {
	for _, cha := range stop {
		cha.Stop()
	}
}

func (cha Cha) Stop() {
	close(cha.quit)
}

func (cha Cha) Quit() bool {
	select {
	case <-cha.quit:
		return true
	default:
	}
	return false
}

func (cha Cha) Done() {
	close(cha.done)
}

func Wait(wait ...Cha) {
	for _, cha := range wait {
		<-cha.done
	}
}

func (cha Cha) Wait() {
	<-cha.done
}

func NewCha() Cha {
	return Cha{make(chan struct{}), make(chan struct{})}
}

