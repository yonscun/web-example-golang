/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/habibiefaried/web-example-golang/controller"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var portBind string

var rootCmd = &cobra.Command{
	Use: "web",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		controller.Basic()
		http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))) // For assets
		if _, err := os.Stat("./assets/uploads"); os.IsNotExist(err) {
			log.Println("Creating upload directory")
			errDir := os.MkdirAll("./assets/uploads", 0755)
			if errDir != nil {
				log.Println(err)
			}
		}

		log.Println("server started at localhost:" + portBind)
		http.ListenAndServe(":"+portBind, nil)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&portBind, "port", "p", "8080", "choose port to bind to")
}
