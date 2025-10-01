package awscdkelasticachealpha


// Engine type for ElastiCache users and user groups.
//
// Example:
//   user := elasticache.NewIamUser(this, jsii.String("User"), &IamUserProps{
//   	// set user engine
//   	Engine: elasticache.UserEngine_REDIS,
//
//   	// set user id
//   	UserId: jsii.String("my-user"),
//
//   	// set username
//   	UserName: jsii.String("my-user"),
//
//   	// set access string
//   	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
//   })
//
// Experimental.
type UserEngine string

const (
	// Valkey engine.
	// Experimental.
	UserEngine_VALKEY UserEngine = "VALKEY"
	// Redis engine.
	// Experimental.
	UserEngine_REDIS UserEngine = "REDIS"
)

