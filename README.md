# nginx-marshal

Helps making nginx configuration in go

## Usage

Import package 

```go
import "github.com/qwerty2586/nginx-marshal/nginx"
```

Build Configuration structure and print it

```go
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
	
	fmt.Println(r)
```

Output be like
```
daemon               off;
worker_processes     2;
user                 www-data;

http {
    server_tokens    off;
    include          mime.types;
    charset          utf-8;

    access_log       logs/access.log combined;

    server {
        server_name  localhost;
        listen       127.0.0.1:80;
        error_page   500 502 503 504 /50x.html;

        location / {
            root     html;
        }
    }
}
```

## TODO

 - nginx.Parse()
 - Builder style functions