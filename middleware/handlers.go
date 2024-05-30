package middleware

import(
	"fmt"
	"log"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	"encoding/json"
	"go-postgres-crud/models"
)

type response struct{
	ID int64 `json:"id",omitempty`
	Message string `json:"message",omitempty`
}

func CreateConnection() *sql.DB{
	err:= godotenv.Load(".env")
	if err!= nil{ log.Fatal("Error loading .env file") }

	db, err:= sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err!= nil{ panic(err) }

	err := db.Ping()
	if err != nil{ panic(err) }

	fmt.Println("Successfully connected to postgres")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request){

	var stock models.Stock

	err:=json.NewDecoder(r.Body).Decode(&stock)
	if err!=nil{
		log.Fatal("Unable to decode the request body. %v",err)
	}
	insertID = 	insertStock(stock)

	res:= response{
		ID: insertID,
		Message: "stock created successfully"
	}
	json.NewEncoder(w).Encode(res)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request){
	stocks,err:= getAllStocks()
	if err!=nil{
		log.Fatal("Unable to get all stocks. %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func GetStockById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if err!=nil{
		log.Fatal("Unable to convert the string into int. %v", err)
	}
	stock,err:=	getStockById(int64(id))
	if err!=nil{
		log.Fatal("Unable to get stock. %v",err)
	}

	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err!=nil{
		log.Fatal("Unable to convert the string  to int. %v",err)
	}
	
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err!=nil{
		log.Fatal("Unable to decode the request body. %v", err)
	}
	updatedRows := updateStock(int64(id),stock)

	msg:= fmt.Sprintf("Stock updated successfully. Total rows affected %v", updatedRows)
	res:= response{
		ID:=int64(id)
		Message: msg
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if err!=nil{
		log.Fatal("Unablr to convert the string to int. %v", err)
	}

	deletedRows:=deleteStock(int64(id))
	ms:= fmt.Sprintf("Stock deleted successfully. total rows deleted %v", deletedRows)
	res:= response{
		ID: int64(id)
		Message: msg
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64{

}

func getStockById(id int64) (models.Stock, error){

}

func getAllStocks()([]models.Stock, error){

}

func updateStock(id int64, stock models.Stock) int64{

}

func deleteStock(id int64) int64{

}