package utils

import (
	"context"
	"fmt"
	//"github.com/elastic/go-elasticsearch"
	"gopkg.in/olivere/elastic.v6"
	"log"
	"os"
)

/*
https://blog.csdn.net/tflasd1157/article/details/81981915
*/

var client *elastic.Client
var host = "http://192.168.1.211:8030/"

func DemoEs() {
	esInit()
}

//func esInit2(){
//	cfg := elasticsearch.Config{
//		Addresses: []string{
//			"http://localhost:9200",
//			"http://localhost:9201",
//		},
//	}
//	es, err := elasticsearch.NewClient(cfg)
//	println(es)
//	println(err)
//}

//初始化
func esInit() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

func esSearch() {

	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
	//printEmployee(res, err)

	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "数据",
	//		},
	//	},
	//}

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	println(res)
	//printEmployee(res, err)

	//if res.Hits.TotalHits > 0 {
	//	fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)
	//
	//	for _, hit := range res.Hits.Hits {
	//
	//		var t Employee
	//		err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
	//		if err != nil {
	//			fmt.Println("Deserialization failed")
	//		}
	//
	//		fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
	//	}
	//} else {
	//	fmt.Printf("Found no Employee \n")
	//}

	////条件查询
	////年龄大于30岁的
	//boolQ := elastic.NewBoolQuery()
	//boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	//boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	//res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	//printEmployee(res, err)
	//
	////短语搜索 搜索about字段中有 rock climbing
	//matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	//res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	//printEmployee(res, err)
	//
	////分析 interests
	//aggs := elastic.NewTermsAggregation().Field("interests")
	//res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	//printEmployee(res, err)

}

func helperDemo() {
	//esutil.JSONReader()
}
