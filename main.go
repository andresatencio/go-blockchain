package main

func main() {
	// bc := NewBlockchain()

	// bc.AddBlock("Send 1 BTC to Andy")
	// bc.AddBlock("Send 2 BTC to Andy")

	// for _, block := range bc.blocks {
	// 	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Println()
	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	// 	fmt.Printf("Nonce: %v\n", block.Nonce)

	// 	fmt.Println()
	// }

	// bc := NewBlockchain()
	// defer bc.db.Close()

	cli := CLI{}
	cli.Run()
}
