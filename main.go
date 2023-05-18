package main

import (
	"database/sql"

	"flag"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/mstoews/xlmparser/utils"
	"log"
	"time"
)

var (
	use_database      = flag.Bool("use_database", false, "enable debugging")
	debug             = flag.Bool("debug", false, "enable debugging")
	password          = flag.String("password", "admin123", "the database password")
	port         *int = flag.Int("port", 31432, "the database port")
	server            = flag.String("server", "localhost", "the database server")
	user              = flag.String("user", "admin", "the database user")
	database          = flag.String("database", "prd-backup-api", "database name")
	source            = flag.String("source", "activity", "Source type (activity)")
	filename          = flag.String("filename", "./xml/activity.xml", "XML file name")
)

func main() {
	start := time.Now()
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" database:%s\n", *database)
		fmt.Printf(" source:%s\n", *source)
		fmt.Printf(" filename:%s\n", *filename)
	}

	//fmt.Printf("Source:  %s\n", *source)
	//fmt.Printf("File name:  %s\n", *filename)

	var psqlconn string

	if *use_database == true {
		psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *server, *port, *user, *password, *database)
		fmt.Printf("Connecting to database %s\n", psqlconn)
		db, err := sql.Open("postgres", psqlconn)
		CheckError(err)
		defer db.Close()
		switch *source {
		case "activity":
			utils.ProcessActivity(*filename, db)
		default:
			panic("Invalid source")
		}
		log.Printf("main, execution time %s\n", time.Since(start))
	} else {
		switch *source {
		case "activity":
			utils.ProcessActivityNoDB(*filename)
		default:
			panic("Invalid source")
		}
		log.Printf("main, execution time %s\n", time.Since(start))
	}

	// Open our xmlFile

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
