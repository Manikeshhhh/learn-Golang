package controller

import "go.mongodb.org/mongo-driver/mongo"
import "log"

//before you try to use the password,firewall rule is set to my ip and DB is already revoked
const connectionString="mongodb+srv://manikesh:manikesh@cluster0.fjsdtsc.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "watchlist"

//Important
var collection *mongo.Collection

//Making Connection
func init(){
	//client option
	clientOption:=options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client,err:=mongo.Connect(contex.TODO(),clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	collection = client.Database(dbName).collection(colName)

	//collection instance 
	fmt.Println("collection instance is ready")
}
