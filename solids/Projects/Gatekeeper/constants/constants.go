package constants

// config related constants
const DefaultConfigurationType = "yaml"
const DefaultConfigurationPath = "environment"

// db related constants
const (
	DBDriver      = "postgres"
	ConstantDB    = "db"
	UserTableName = "users"
)

// application related constants
const (
	AppName = "Gatekeeper"
)

// server related constants
const (
	Origin        = "origin"
	ContentLength = "Content-Length"
	ContentType   = "Content-Type"
	Authorization = "Authorization"
)

// user service logging messages
const (
	// validation errors
	LogUserValidationFailed     = "user validation failed with invalid input data"
	LogUserEmailAlreadyExists   = "user with this email already exists in the system"
	LogFailedCheckUserExistence = "failed to check if user exists in database"
	LogUserNameAlreadyExist     = "user with same username exist in database"

	// password processing
	LogFailedHashPassword = "failed to hash password during user creation"

	// database operations
	LogFailedSaveUserToDB = "failed to save user to database"

	// handler logs
	LogValidationError  = "request validation failed - invalid input parameters"
	LogConflictError    = "resource conflict - user already exists"
	LogUserDoesNotExist = "failed to fetch user - user with given username not exist"
	LogInvalidPassword  = "invalid user credentials - provided password is incorrect"
	LogTokenkErr        = "failed to generate token"
)

// Authentication related constant
const (
	MinSecretKeyLen           = 32
	ExipredToken              = "token has expired"
	InvalidToken              = "token is invalid"
	JWTValidationErrorExpired = 512
	UserName                  = "username"
)
