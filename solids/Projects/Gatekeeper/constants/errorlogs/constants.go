package errorlogs

const (
	ParsingError        = "error while parsing the configuration file: %v"
	ServerError         = "error while starting the server: %v"
	ExtensionError      = "error while enabling extensions: %v"
	MigrationError      = "error while migrating the database: %v"
	SaltGenerationError = "error while generating salt: %v"
	BindJsonError       = "error while binding JSON"
	CreateUserError     = "error while creating user: %v"
	InvalidKeySize      = "invalid key size : must be atleast %d characters long"
)
