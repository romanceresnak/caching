package main
import
(
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/patrickmn/go-cache"
)
const
(
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)
var newCache *cache.Cache

//init run always as a first
func init() {
	//set time in cahce to 5 minute
	newCache = cache.New(5*time.Minute, 10*time.Minute)
	//key is foo value is bar
	newCache.Set("foo", "bar", cache.DefaultExpiration)
}
func getFromCache(w http.ResponseWriter, r *http.Request) {
	//find object by key
	foo, found := newCache.Get("foo")
	//if exist key in memory like this write the result
	if found {
	log.Print("Key Found in Cache with value as :: ",
	foo.(string))
	fmt.Fprintf(w, "Hello "+foo.(string))
	} else
	{
	log.Print("Key Not Found in Cache :: ", "foo")
	fmt.Fprintf(w, "Key Not Found in Cache")
	}
}
func main() {
	http.HandleFunc("/", getFromCache)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
	log.Fatal("error starting http server : ", err)
	return
	}
}