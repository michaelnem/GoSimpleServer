package main

import (
	"fmt"
	"nroute/core"
	"nroute/core/n_http"
)

func handler1(path string) {
	fmt.Println(1, path)
}

func handler2(path string) {
	fmt.Println(2, path)
}

func handler3(path string) {
	fmt.Println(3, path)
}

func main() {
	router := core.NewRouter()

	router.GET("/", handler1)
	router.GET("a/b", handler2)
	router.GET("a/c", handler3)
	router.GET("a/c/d", handler1)
	router.GET("a/c/d", handler2)

	router.Use(n_http.GET, "a/c")
	router.Use(n_http.GET, "a/cd")
	router.Use(n_http.POST, "a/c")

	// router.UseRoute("a/c")
	// router.UseRoute("a/c/d")

	fmt.Println(router)
}
