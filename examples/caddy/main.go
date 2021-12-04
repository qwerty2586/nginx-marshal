package main

import (
	"fmt"
	caddyfile "github.com/qwerty2586/nginx-marshal/nginx"
	"time"
)

func main() {
	caddyfile.TrailingCharacterEncoding = ""
	r := caddyfile.Root{
		Submodules: []caddyfile.Submodule{{
			Lines: []string{
				"debug",
				"http_port 80",
				"https_port 443",
			},
			Submodules: []caddyfile.Submodule{{
				Name: "servers :443",
				Submodules: []caddyfile.Submodule{{
					Name: "protocol",
					Lines: []string{
						"experimental_http3",
					},
				}},
			}},
		}},
	}
	start := time.Now()
	fmt.Println("Caddyfile support yay\n")
	fmt.Println(r)
	fmt.Println("time took: ", time.Since(start))
}
