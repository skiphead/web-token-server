# About the Web-token-server

This is a simple microservice for generating tokens and checking their validity.

# How to build and run?

#### cd project folder:

    - Edited config file ./config/server.json 
    (listen port, expired live token (sec), TLS on/off, pki cert and key files )
    
    - go build ./cmd/main.go
     
    - main(windows main.exe) --config-path="/path-config/server.json"
     
    - default settings http://127.0.0.1:8080, expireted  live after 600 sec

# How to use?

#### New generate token


  - Request POST: /new

    {
       "name": string
    }

  
  - Response POST:

    {
       "token": string
    }


#### Check valid token

  
  - Request POST: /check
     
    {
       "token": string
    }

  
  - Response POST:

    {
       "valid": bool
    }

#### Info token


  - Request POST: /info

     {
        "token": string
     }


  - Response POST:

     {
       "token": string
       "name": string
     }



Stable work on GoLang 1.8.3
