package main

import "net/http"

func home(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("are you lost? try visiting /flag you might find something useful"))
}

func fake(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("what kind of flag? there is no flag here, i heard it's at /fl"))
}

func half(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("wait so you really are looking for a flag? you should find more at /ag . here is the first half, flag: pai{p12aK7i"))
}

func flag(w http.ResponseWriter, _ *http.Request) { // flag: pai{p12aK7ikUm_w1r3sH4rk}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("here is the other half, flag: kUm_w1r3sH4rk}"))
}

func main() {
	router := http.NewServeMux()
	{
		router.HandleFunc("GET /", home)
		router.HandleFunc("GET /flag", fake)
		router.HandleFunc("GET /fl", half)
		router.HandleFunc("GET /ag", flag)
	}

	server := http.Server{
		Addr:    "0.0.0.0:1337",
		Handler: router,
	}

	server.ListenAndServe()
}
