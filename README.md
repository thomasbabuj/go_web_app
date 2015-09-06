# go_web_app
Go Web App ( Experiment )


## Video Comments :
    
    Processing HTTP requests with Go is primarliy about two things :

        1) ServeMux aka Request Router aka Multiplexor

        2) Handelers

    ## ServeMux :

        ServeMux = HTTP Request router = Multiplexor = Mux

        Compares incoming request against a list of predefined URL paths, and 
        calls the assoicated handler for path whenever a match is found.

    ## Handlers :

        Responsible for writing response headers and bodies.

        Almost any type ("object") can be a handlers, so long as it 
        satisfies the http.Hander Interface.

        In lay man terms, that simply means it must have a ServerHTTP method with
        a following signature :

            ServerHTTP(http.ResponseWriter, *http.Request)




