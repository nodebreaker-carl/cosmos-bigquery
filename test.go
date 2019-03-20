package main

import (

           "fmt"

           "log"

 

           "cloud.google.com/go/bigquery"

           "golang.org/x/net/context"

)

 

type Item struct {

           Name  string

           Count int

}

 

// Save implements the ValueSaver interface.

func (i *Item) Save() (map[string]bigquery.Value, string, error) {

           return map[string]bigquery.Value{

                      "Name":  i.Name,

                      "Count": i.Count,

           }, "", nil

}

 

func main() {

           ctx := context.Background()

 

           // Sets your Google Cloud Platform project ID.

           projectID := "cystest-1" //프로젝트ID

           datasetID := "gotest" //데이터셋ID

           tableID := "gotable" //테이블ID

 

           // Creates a client.

           client, err := bigquery.NewClient(ctx, projectID)

           if err != nil {

                      log.Fatalf("Failed to create client: %v", err)

           }

           insertRows(client, datasetID, tableID)    

}

 

func insertRows(client *bigquery.Client, datasetID, tableID string) error {

           ctx := context.Background()

                     

           // [START bigquery_insert_stream]

           u := client.Dataset(datasetID).Table(tableID).Uploader()

           items := []*Item{

                      // Item implements the ValueSaver interface.

                      // 아래에 테이블에 삽입할 레코드를 입력한다.

                      {Name: "n1", Count: 7},

                      {Name: "n2", Count: 2},

                      {Name: "n3", Count: 1},

           }

           if err := u.Put(ctx, items); err != nil {

                      return err

           }

           // [END bigquery_insert_stream]

           fmt.Printf("insert complete\n")

           return nil

}
