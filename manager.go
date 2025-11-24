package main


import (
	"github.com/gorilla/websocket"
    "net/http"
	"strings"
	"log"
	// "./utils"
)

var allowedOrigins = []string{
    "http://localhost:3000",
    "https://yourdomain.com",
    "https://trustedapp.com",
    "*.yourdomain.com", 
}

// mark as used to avoid "unused" lint error when the variable is referenced only in closures
var _ = allowedOrigins


var websocketUpgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        origin := r.Header.Get("Origin")

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
    },
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}


type Manager struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan []byte
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

func NewManager() *Manager {
	return &Manager{};
	// return &Manager{
	// 	Clients:    make(map[*websocket.Conn]bool),
	// 	Broadcast:  make(chan []byte),
	// 	Register:   make(chan *websocket.Conn),
	// 	Unregister: make(chan *websocket.Conn),
	// }
}

func (m *Manager) serverWS(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket Endpoint Hit");
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err);
		return
	}
	conn.Close();
}



// func (m *Manager) serverWS(w http.ResponseWriter, r *http.Request) {
//     log.Println("WebSocket Endpoint Hit")
//     upgrader := newWebsocketUpgrader()
//     conn, err := upgrader.Upgrade(w, r, nil)
//     if err != nil {
//         log.Println(err)
//         return
//     }

//     // register connection if manager is initialized
//     if m != nil {
//         m.Register <- conn
//     }

//     // read loop in a goroutine; unregister and close on exit
//     go func(c *websocket.Conn) {
//         defer func() {
//             if m != nil {
//                 m.Unregister <- c
//             }
//             c.Close()
//         }()

//         for {
//             _, message, err := c.ReadMessage()
//             if err != nil {
//                 log.Println("read error:", err)
//                 break
//             }
//             if m != nil {
//                 m.Broadcast <- message
//             }
//         }
//     }(conn)
// }




// var websocketUpgrader = websocket.Upgrader{
//     CheckOrigin: func(r *http.Request) bool {
//         origin := r.Header.Get("Origin")

//         // Allow empty origin (e.g., Postman or some native apps)
//         if origin == "" {
//             return false // or true depending on your use-case
//         }

//         for _, allowed := range allowedOrigins {
//             if origin == allowed {
//                 return true
//             }
//         }

//         return false // Reject everything else
//     },
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
// }

// var websocketUpgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// var  (
// 	websocketUpgrader = websocket.Upgrader{
// 		ReadBufferSize:  1024,
// 	    WriteBufferSize: 1024,
// 	}
// )