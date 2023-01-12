package awselasticache


// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authenticationMode interface{}
//
//   cfnUserProps := &cfnUserProps{
//   	engine: jsii.String("engine"),
//   	userId: jsii.String("userId"),
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	accessString: jsii.String("accessString"),
//   	authenticationMode: authenticationMode,
//   	noPasswordRequired: jsii.Boolean(false),
//   	passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   }
//
type CfnUserProps struct {
	// The current supported value is redis.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The ID of the user.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// The username of the user.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Access permissions string used for this user.
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// `AWS::ElastiCache::User.AuthenticationMode`.
	AuthenticationMode interface{} `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Indicates a password is not required for this user.
	NoPasswordRequired interface{} `field:"optional" json:"noPasswordRequired" yaml:"noPasswordRequired"`
	// Passwords used for this user.
	//
	// You can create up to two passwords for each user.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

