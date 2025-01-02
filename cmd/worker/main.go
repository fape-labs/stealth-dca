package main

import (
	"context"
	"github.com/fape-labs/stealth-dca/internal/jup"
	"github.com/fape-labs/stealth-dca/internal/log"
	"github.com/fape-labs/stealth-dca/internal/solanaclient"
	"github.com/gagliardetto/solana-go"
	"go.uber.org/zap"
	"time"
)

var logger = log.CreateLogger()

func main() {
	logger.Debug("Starting Stealth DCA Worker")
	jupClient := jup.JupClient{BaseUrl: "https://quote-api.jup.ag/v6"}
	rpcClient, err := solanaclient.NewClient(context.Background(), "https://api.mainnet-beta.solana.com", "wss://api.mainnet-beta.solana.com")
	if err != nil {
		panic(err)
	}
	fapeMint := solana.MustPublicKeyFromBase58("GA1UvKdQLi3BQ7jbresAT4JQA7R9K9AZW9X9MvFSuvZk")
	amountEach := solana.LAMPORTS_PER_SOL / 10

	signer, err := solana.NewRandomPrivateKey()
	if err != nil {
		panic(err)
	}

	for {

		tx, err := jupClient.Swap(rpcClient, signer, solana.SolMint, fapeMint, amountEach, 100)
		if err != nil {
			logger.Error("failed to swap", zap.Error(err))
		}

		logger.Info("swap finished", zap.String("tx", tx.String()))
		<-time.After(1 * time.Minute)
	}

	<-make(chan struct{})
}
