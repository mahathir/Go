package utilities

// DIConstants is Dependency Injection constants to help DI.
type DIConstants struct {
	IUsrServ string
	IUsrDa   string
}

// NewDIConstants is to initialized DIConstants
func NewDIConstants() *DIConstants {
	return &DIConstants{
		IUsrServ: "IUsrServ",
		IUsrDa:   "IUsrDa",
	}
}
