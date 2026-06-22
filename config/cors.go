package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://belatihan-production-7069.up.railway.app",
	"https://my-fe-ten.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}