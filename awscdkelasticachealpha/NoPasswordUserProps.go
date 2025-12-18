package awscdkelasticachealpha


// Properties for defining an ElastiCache user with no password authentication.
//
// Example:
//   // use the original `default` user by using import method
//   defaultUser := elasticache.NoPasswordUser_FromUserAttributes(this, jsii.String("DefaultUser"), &UserBaseAttributes{
//   	// userId and userName must be 'default'
//   	UserId: jsii.String("default"),
//   })
//
//   // create a new default user
//   newDefaultUser := elasticache.NewNoPasswordUser(this, jsii.String("NewDefaultUser"), &NoPasswordUserProps{
//   	// new default user id must not be 'default'
//   	UserId: jsii.String("new-default"),
//   	// new default username must be 'default'
//   	UserName: jsii.String("default"),
//   	// set access string
//   	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
//   })
//
// Experimental.
type NoPasswordUserProps struct {
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

