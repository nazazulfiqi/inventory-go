// handlers/middleware.go
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Custom type untuk context key
type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{w, http.StatusOK}

		// Log informasi dasar request
		// log.Printf(
		// 	"[%s] %s %s - Started",
		// 	start.Format("2006-01-02 15:04:05"),
		// 	r.Method,
		// 	r.URL.Path,
		// )

		next.ServeHTTP(lrw, r)

		// Ambil informasi user dari context jika ada
		var userID interface{}
		var role interface{}
		if ctx := r.Context(); ctx != nil {
			userID = ctx.Value(UserIDKey)
			role = ctx.Value(RoleKey)
		}

		duration := time.Since(start)
		log.Printf(
			"[%s] %s %s - Completed | Status: %d | Duration: %v | UserID: %v | Role: %v",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			lrw.statusCode,
			duration.Round(time.Millisecond),
			userID,
			role,
		)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Ekstrak token dari header
		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString = splitToken[1]

		// Parse dan validasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Token is invalid", http.StatusUnauthorized)
			return
		}

		// Ekstrak claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Validasi dan ekstrak user_id
		userID, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		// Validasi dan ekstrak role
		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Invalid role in token", http.StatusUnauthorized)
			return
		}

		// Validasi role yang diizinkan
		validRoles := map[string]bool{
			"admin":      true,
			"supervisor": true,
			"karyawan":   true,
		}

		if !validRoles[role] {
			http.Error(w, "Unauthorized role", http.StatusForbidden)
			return
		}

		// Tambahkan ke context
		ctx := context.WithValue(r.Context(), UserIDKey, uint(userID))
		ctx = context.WithValue(ctx, RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Middleware untuk role admin
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(RoleKey).(string)
		if !ok || role != "admin" {
			http.Error(w, "Admin access required", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware untuk supervisor dan admin
func SupervisorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(RoleKey).(string)
		if !ok || (role != "supervisor" && role != "admin") {
			http.Error(w, "Supervisor or admin access required", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
