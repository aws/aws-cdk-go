package awscdkelasticachealpha


// Properties for defining an ElastiCache user with IAM authentication.
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
type IamUserProps struct {
	// Access control configuration for the user.
	// Experimental.
	AccessControl AccessControl `field:"required" json:"accessControl" yaml:"accessControl"`
	// The ID of the user.
	// Experimental.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// The engine type for the user.
	//
	// Enum options: UserEngine.VALKEY, UserEngine.REDIS.
	// Default: - UserEngine.REDIS for NoPasswordUser, UserEngine.VALKEY for all other user types.
	//
	// Experimental.
	Engine UserEngine `field:"optional" json:"engine" yaml:"engine"`
	// The name of the user.
	// Default: - Same as userId.
	//
	// Experimental.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

