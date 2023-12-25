package database

//
//import (
//	"fmt"
//	"github.com/spf13/viper"
//	"log"
//)
//
//func Set() {
//
//	viper.SetConfigName("db")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath("database")
//
//	if err := viper.ReadInConfig(); err != nil {
//		log.Fatal("Error reading the configs")
//	}
//
//	err := viper.Unmarshal(&dbInstance)
//	fmt.Println("host in setter is", dbInstance.Host)
//	if err != nil {
//		log.Fatal("unable to decode into struct")
//	}
//
//}
