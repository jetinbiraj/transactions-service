package main

// @title			Transactions Service
// @version		v1
// @description	Exposed APIs for accounts and transactions related operations
// @contact.name	Jetin Biraj
// @contact.email	birajdarjk1106@gmail.com
func main() {
	if err := StartTransactionsService(); err != nil {
		panic(err)
	}

	// TODO: Extend for graceful shutdown
}
