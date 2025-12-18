package awscdkelasticachealpha


// Properties for defining an ElastiCache base user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import elasticache_alpha "github.com/aws/aws-cdk-go/awscdkelasticachealpha"
//
//   var accessControl AccessControl
//
//   userBaseProps := &UserBaseProps{
//   	AccessControl: accessControl,
//   	UserId: jsii.String("userId"),
//
//   	// the properties below are optional
//   	Engine: elasticache_alpha.UserEngine_VALKEY,
//   }
//
// Experimental.
type UserBaseProps struct {
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
}

