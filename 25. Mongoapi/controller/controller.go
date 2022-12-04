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

//Mongodb helper

//insert 1 record

func insetOneMovie(movie model.Netflix){
	inserted,err:=collection.InsertOne(contex.Backgroud(),movie)
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("Inserted 1 movie in DB with Id: ",inserted.InsertedID)
}

//update 1 record

func updateOneMovie(movieID string){
	id,_:=primitive.ObjectIDFromHex(movieID)
	filter:=bson.M{"_id": id}
	update:=bson.M{"$set":bson.M{"watched": true}}

	result,err:=collection.UpdateOne(contex.Backgroud(),filter,update)
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("modified count: ",result.ModifiedCount)
}