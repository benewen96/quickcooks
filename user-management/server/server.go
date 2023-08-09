package server

func init() {
	router := newRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
