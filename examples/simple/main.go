package main

import (
	"fmt"
	"nginx-marshal/nginx"
	"time"
)

func main()  {
	r := nginx.Root{
		Lines: []string{
			"daemon off",
			"worker_processes 2",
			"user www-data",
		},
		Submodules: []nginx.Submodule{{
			Name: "http",
			Lines: []string{
				"server_tokens off",
				"include       mime.types",
				"charset utf-8",
				"",
				"access_log logs/access.log combined;",
			},
			Submodules: []nginx.Submodule{{
				Name: "server",
				Lines: []string{
					"server_name localhost",
					"listen 127.0.0.1:80",
					"error_page 500 502 503 504  /50x.html",
				},
				Submodules: []nginx.Submodule{{
					Name: "location /",
					Lines: []string{
						"root html",
					}},
				}},
			}},
		},
	}
	start := time.Now()
	fmt.Println("Example configuration from random tutorial\n")
	fmt.Println(r)
	fmt.Println("time took: ",time.Since(start))
}