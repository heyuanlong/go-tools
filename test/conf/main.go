package main

import (
	"log"

	"github.com/heyuanlong/go-tools/conf"
)

func main() {
	kconf, err := conf.NewKconf("./conf.json")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(kconf.GetInt64("server.port"))
	log.Println(kconf.GetInt("server.timestamp"))
	log.Println(kconf.GetBool("server.istest"))
	log.Println(kconf.GetFloat64("server.num"))
	log.Println(kconf.GetString("redis.auth"))
	log.Println()

	log.Println(kconf.GetInt64("redis.xxx"))
	log.Println(kconf.GetInt("redis.xxx"))
	log.Println(kconf.GetBool("redis.xxx"))
	log.Println(kconf.GetFloat64("redis.xxx"))
	log.Println(kconf.GetString("redis.xxx"))
	log.Println()

	arrs, _ := kconf.Container.Path("strings").Children()
	for _, v := range arrs {
		v, ok := v.Data().(string)
		if ok {
			log.Println(v)
		}
	}
}
