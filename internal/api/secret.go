package api

type Secret struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	SecretValue string `json:"secret"`
	Version     int    `json:"version"`
}
