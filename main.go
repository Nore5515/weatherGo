package main

import (
	//"fmt"

	"context"
	"fmt"
	"io/ioutil"
	"time"

	//"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//mongoStuff()
	go forever()
	select {} // block forever
}

func forever() {
	for {
		//30.520035124090917, -97.77904654503592
		//api.openweathermap.org/data/2.5/weather?lat=30.520035124090917&lon=-97.77904654503592&appid=f4383eb27993d08083af66327f662a77
		//fmt.Printf("%v+\n", time.Now())
		URL := "https://api.openweathermap.org/data/2.5/weather?lat=30.520035124090917&lon=-97.77904654503592&appid=f4383eb27993d08083af66327f662a77"
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)
		log.Printf(sb)

		// HELP
		//foo := bson.D{{"poop", "peepee"}, {"ding dong", "dorkie"}}
		mongoStuff(body)

		time.Sleep(time.Second * 100000)
	}
}

func mongoStuff(foo []byte) {

	//uri := os.Getenv("MONGODB_URI")
	uri := "mongodb+srv://Nore5515:thealphabetbackwards@cluster0.oijvz.mongodb.net/eek?retryWrites=true&w=majority"
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/#environment-variable")
	}

	fmt.Println("1")
	credential := options.Credential{
		Username: "Nore5515",
		Password: "thealphabetbackwards",
	}

	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	_ = client

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	fmt.Println("2")

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	fmt.Println("3")

	coll := client.Database("Weather").Collection("weathers")
	// DA STUFF

	var doc interface{}
	err = bson.UnmarshalExtJSON(foo, true, &doc) //bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}
	if err != nil {
		fmt.Println("Err 1")
		log.Fatal(err)
	}
	//doc := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println("Err 2")
		panic(err)
	}

	fmt.Println(result)
	/*
		coll := client.Database("Cluster0").Collection("eek")
		var result bson.M
		err = coll.FindOne(context.TODO(), bson.D{{"title", "The Room"}}).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				return
			}
			panic(err)
		}
	*/

}
