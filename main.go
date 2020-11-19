package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Wallet struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

var wallets [] Wallet

func getWallets(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(wallets)
}
func getWalletById(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	//loop through wallets
	for _, item := range wallets{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Wallet{})
}
func createWallet(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var wallet Wallet
	_ = json.NewDecoder(r.Body).Decode(&wallet)
	//wallet.ID = strconv.Itoa(rand.Intn(10000))
	wallets = append(wallets,wallet)
	json.NewEncoder(w).Encode(wallet)

}
func updateWallet(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item :=range wallets{
		if item.ID==params["id"]{
			wallets =append(wallets[:index],wallets[index+1:]...)
			var wallet Wallet
			_ = json.NewDecoder(r.Body).Decode(&wallet)
			wallet.ID = params["id"]
			wallets = append(wallets,wallet)
			json.NewEncoder(w).Encode(wallet)
			return
		}
	}
	json.NewEncoder(w).Encode(wallets)
}
func deleteWallets(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item :=range wallets{
		if item.ID==params["id"]{
			wallets =append(wallets[:index],wallets[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(wallets)
}

func main()  {
	//init mux
	fmt.Println("server started")
	r:=mux.NewRouter()
	//wallets =append(wallets,Wallet{ID: "1",Name: "Phonepe",Amount: 1000,CreatedAt: time.Now()})
	//wallets =append(wallets,Wallet{ID: "2",Name: "Paytm",Amount: 1500,CreatedAt: time.Now()})
	//handing routes
	r.HandleFunc("/api/wallets",getWallets).Methods("GET")
	r.HandleFunc("/api/wallets/{id}",getWalletById).Methods("GET")
	r.HandleFunc("/api/wallets",createWallet).Methods("POST")
	r.HandleFunc("/api/wallets/{id}",updateWallet).Methods("PUT")
	r.HandleFunc("/api/wallets/{id}",deleteWallets).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))
}
