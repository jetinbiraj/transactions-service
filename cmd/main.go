package main

func main() {
	if err := StartTransactionsService(); err != nil {
		panic(err)
	}

	// TODO: Extend for graceful shutdown
}
