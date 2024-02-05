package main

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"os"

	mycrypto "datacollection/protos"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

func main() {
    apiKey := os.Getenv("API_KEY")
    secret := os.Getenv("SECRET")

    u := url.URL{Scheme: "wss", Host: "stream.data.alpaca.markets", Path: "/v1beta3/crypto/us"}
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    // Send authentication message
    authMessage := map[string]string{
        "action": "auth",
        "key":    apiKey,
        "secret": secret,
    }
    authMessageJSON, _ := json.Marshal(authMessage)
    err = c.WriteMessage(websocket.TextMessage, authMessageJSON)
    if err != nil {
        log.Println("write:", err)
        return
    }

    // Send subscription message
    subscriptionMessage := map[string]interface{}{
        "action": "subscribe",
        "trades": []string{},
        "quotes": []string{
            "BTC/USD", "ETH/USD", "BNB/USD", "USDT/USD", "USDC/USD",
            "XRP/USD", "ADA/USD", "DOT/USD", "LTC/USD", "BCH/USD",
            "LINK/USD", "XLM/USD", "UNI/USD", "DOGE/USD", "WBTC/USD",
            "AAVE/USD", "ATOM/USD", "XMR/USD", "XTZ/USD", "EOS/USD",
            "BSV/USD", "TRX/USD", "VET/USD", "SOL/USD", "MIOTA/USD",
            "THETA/USD", "SNX/USD", "NEO/USD", "MKR/USD", "COMP/USD",
        },
        "bars": []string{},
    }
    subscriptionMessageJSON, _ := json.Marshal(subscriptionMessage)
    err = c.WriteMessage(websocket.TextMessage, subscriptionMessageJSON)
    if err != nil {
        log.Println("write:", err)
        return
    }

	conn, err := grpc.Dial("detection-algos:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		// Don't return, just log the error
	}
	defer conn.Close()

	var stream mycrypto.CryptoStream_StreamCryptoClient
	if err == nil {
		client := mycrypto.NewCryptoStreamClient(conn)

		// Start the stream
		stream, err = client.StreamCrypto(context.Background())
		if err != nil {
			log.Printf("Error on stream: %v", err)
			// Don't return, just log the error
		}
	}

		// Handle incoming messages
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		type JSONCryptoData struct {
			T string  `json:"T"`
			S string  `json:"S"`
			Bp float64 `json:"bp"`
			Bs float64 `json:"bs"`
			Ap float64 `json:"ap"`
			As float64 `json:"as"`
			Time string `json:"t"`
		}
		
		
		// Unmarshal the message into a slice of JSONCryptoData objects
		var jsonData []JSONCryptoData
		err = json.Unmarshal(message, &jsonData)
		if err != nil {
			log.Println("json unmarshal:", err)
			continue
		}
		
		// Convert JSONCryptoData objects to CryptoData objects
		var cryptoData []mycrypto.CryptoData
		for _, data := range jsonData {

			// Ignore control messages
			if data.T == "success" || data.T == "subscription" {
				continue
			}

			cryptoData = append(cryptoData, mycrypto.CryptoData{
				T: data.T,
				S: data.S,
				Bp: data.Bp,
				Bs: data.Bs,
				Ap: data.Ap,
				As: data.As,
				Time: data.Time,
			})
		}
	
		// Send each CryptoData object to the server
		for _, data := range cryptoData {
			if stream != nil {
				if err := stream.Send(&data); err != nil {
					log.Println("stream send:", err)
					return
				}
			}
		}
	}
}