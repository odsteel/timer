package poll

func Start() <-chan Input {
	input := make(chan Input)
	go poll(input)
	return input
}
