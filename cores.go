package main;


import(
	"net/http"
	"strings"
	"log"
)

var allowedOrigins = []string{
	"http://localhost",
	"https://yourdomain.com",
	"https://trustedapp.com",
	"*.domain.com",
}
func CheckOrigin(r *http.Request) bool{
		origin := r.Header.Get("Origin")
		log.Println("Origin Received:", origin)
		log.Println("Host:", r.Host)


		if origin == "" {
			return false
		}

		if origin == "http://"+r.Host || origin == "https://"+r.Host {
			return true
		}

		for _, allowed := range allowedOrigins {
			if strings.HasPrefix(allowed, "*.") {
				if matchesWildcard(origin, allowed) {
					return true
				}
			}

			if origin == allowed {
				return true
			}
		}

		return false
}