package awscdkelasticachealpha


// Attributes for importing an existing ElastiCache user.
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
type UserBaseAttributes struct {
	// The engine type for the user.
	// Default: - engine type is unknown.
	//
	// Experimental.
	Engine UserEngine `field:"optional" json:"engine" yaml:"engine"`
	// The ARN of the user.
	//
	// One of `userId` or `userArn` is required.
	// Default: - derived from userId.
	//
	// Experimental.
	UserArn *string `field:"optional" json:"userArn" yaml:"userArn"`
	// The ID of the user.
	//
	// One of `userId` or `userArn` is required.
	// Default: - derived from userArn.
	//
	// Experimental.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// The user's name.
	// Default: - name is unknown.
	//
	// Experimental.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

