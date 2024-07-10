package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/mejiadev643/app/jwt"
)


func ControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Haz lo que quieras aquí antes de pasar la solicitud al siguiente middleware o controlador
		fmt.Println("Request received: ", r.Method, r.URL.Path)
		token := jwt.GetToken(r.Header.Get("Authorization"))
		// Verificar si el token está presente
        if token == "" {
            fmt.Println("No token found")
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode("No token found")
            return
        }
        
        fmt.Println("Token: ", token)
		payload, err := jwt.GetPayload(token)
		if err != nil {
            // Manejar el error si no se puede obtener el payload, por ejemplo, token inválido
            fmt.Println("Error getting payload: ", err)
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode("Error getting payload")
            return
        }
		fmt.Println(payload)
		//obtener un permiso del token y recorrerlo
		permiso, ok := payload["permisions"]
		if !ok {
            // Si no se encuentran los permisos en el payload
            fmt.Println("Permissions not found in token")
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Unauthorized: Permissions not found"))
            return
        }
		//recorrer con un bucle for para comparar los permisos
		for _, v := range permiso.([]interface{}) {
			//fmt.Println(v)
			if v == "read" {
				fmt.Println("Tiene permiso de lectura")
				next.ServeHTTP(w, r)
				return
			}
		}
		//fmt.Println(permiso)
		// Llama al siguiente middleware o controlador

		//next.ServeHTTP(w, r)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Unauthorized: No read permission")

		// Haz lo que quieras aquí después de que la solicitud haya sido manejada( noe s necesariio en este caso)
		fmt.Println("Request handled: ", r.Method, r.URL.Path)
	})
}