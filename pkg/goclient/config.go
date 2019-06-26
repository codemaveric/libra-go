package goclient

type LibraNetwork int

const (
	TestNet LibraNetwork = iota
	MainNet
)

type LibraClientConfig struct {
	Host    string
	Port    string
	Network LibraNetwork
}
