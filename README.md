# go_web_app
Go Web App ( Experiment )

NICE TO KNOW
www.rawgit.com

## Video Comments :
    
    - http.ListenAndServe
        
        -- listens for, and responds to, http requests
        
        -- handles each request using go routines
        
            --- lightweight concurrency (eg, coroutines - processes --> threads --> coroutines)
            --- this is multiplexing, thus, multiplexor ( = HTTP request router = ServeMux = mux )
            --- blacks main thread (call after configuration of server complete)

    - http.Handle
        
        -- handles a URL request        
        -- maps a URL to any TYPE ("object") implementing the handler interface
            
            --- http://golang.org/pkg/net/http/#Handler
    
    - http.HandleFunc
        
        -- handles a URL request
        -- maps a URL to a FUNCTION

            --- "wrapper" around a function
                ---- turns any function into a handler

    Handle -> handler

    HandleFunc -> handlerFunc
