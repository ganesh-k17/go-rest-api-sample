run below command in the terminal.

go to the folder cd simple-route-mux

go run main.go


In this program we have used gorilla/mux library for handling routing and added below route.

myRouter.HandleFunc("/", homePage)
myRouter.HandleFunc("/articles", returnAllArticles)
myRouter.HandleFunc("/article/{id}", getArticle)