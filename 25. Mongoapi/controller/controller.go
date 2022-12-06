package controller

import "go.mongodb.org/mongo-driver/mongo"
import "encoding/json"
import "github.com/gorilla/mux"
import "log"
import "fmt"

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


//DELETE 1 record
func deleteOneMovie(movieID string){
	id,_:=primitive.ObjectIDFromHex(movieID)
	filter:=bson.M{"_id": id}
	deleteCount,err:=collection.DeleteOne(contex.Backgroud(),filter)
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("Movie got deleted with delete count: ",deleteCount)
}

//Delete All records from mongo
func deleteAllMovie() int64{
	
	deleteResult,err:=collection.DeleteMany(contex.Backgroud(),bson.D{{}},nil) //{{}} This will select all 
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("Number of movie deleted: ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//	Get All movies from mongo
func getAllMovies() []primitive.M{
	cur,err:=collection.Find(contex.Backgroud(),bson.D{{}})
	if err != nil {
		log.Fatal(err) 
	}
	var movies []primitive.M
	for cur.Next(contex.Backgroud()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err) 
		}
		movies = append(movies,movie)
	}
	defer cur.Close(contex.Backgroud())

}

//Actul controller
func GetMyAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies:=getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request){
c
	var movie model.Netflix
	_:= json.NewDecoder(r.Body).Decode(&movie)
	insetOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params:=mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params:=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])	
}

func DeleteAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count:=deleteAllMovie()
	json.NewEncoder(w).Encode(count)	
}