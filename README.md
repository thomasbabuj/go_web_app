# go_web_app
Go Web App ( Experiment )

In this repo, trying to follow this video tutorial (https://www.youtube.com/watch?v=Vlie-srOU8c )
to build basic Web App using Go.

    1) Setup a new repo and add basic html file

    2) Since, I have added Go PATH and Go bin directory configured in my .zshrc 
       I can go install or go build to compile a project from my src directory

    3) To run the project, just use go_web_app in the project root directory.
       Eg, by running "go_web_app" command will out Hello world in the 
       terminal.

# Video Comments,

    1) In Go everything is based on congfiguration

    2) Need to create Resource Server

        - listen on a TCP port 
        - handle requests: route a URL to a file

    3) ServeMux = HTTP request router = multiplexor
            http://www.alexedwards.net/blog/a-recap-of-request-handling

        Mutexes are something else
            http://www.alexedwards.net/blog/understanding-mutexes


Reference for Writing, building, installing, and testing Go code :
        
        https://www.youtube.com/watch?v=XCsL89YtqCs