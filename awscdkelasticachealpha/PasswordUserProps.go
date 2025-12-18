package awscdkelasticachealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining an ElastiCache user with password authentication.
//
// Example:
//   user := elasticache.NewPasswordUser(this, jsii.String("User"), &PasswordUserProps{
//   	// set user engine
//   	Engine: elasticache.UserEngine_VALKEY,
//
//   	// set user id
//   	UserId: jsii.String("my-user-id"),
//
//   	// set access string
//   	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
//
//   	// set username
//   	UserName: jsii.String("my-user-name"),
//
//   	// set up to two passwords
//   	Passwords: []SecretValue{
//   		awscdk.SecretValue_SecretsManager(jsii.String("SecretIdForPassword")),
//   		awscdk.SecretValue_*SecretsManager(jsii.String("AnotherSecretIdForPassword")),
//   	},
//   })
//
// Experimental.
type PasswordUserProps struct {
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
	// The passwords for the user.
	//
	// Password authentication requires using 1-2 passwords.
	// Experimental.
	Passwords *[]awscdk.SecretValue `field:"required" json:"passwords" yaml:"passwords"`
	// The name of the user.
	// Default: - Same as userId.
	//
	// Experimental.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

