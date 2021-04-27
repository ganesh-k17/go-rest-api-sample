run below command in the terminal.

go to the folder cd simple-route

go run main.go


In this program we have used http/net library for handling routing and added below route.

http.HandleFunc("/", homePage)
http.HandleFunc("/articles", returnAllArticles)