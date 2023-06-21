package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUser`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authenticationMode interface{}
//
//   cfnUserProps := &CfnUserProps{
//   	Engine: jsii.String("engine"),
//   	UserId: jsii.String("userId"),
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	AccessString: jsii.String("accessString"),
//   	AuthenticationMode: authenticationMode,
//   	NoPasswordRequired: jsii.Boolean(false),
//   	Passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
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
	// Specifies the authentication mode to use. Below is an example of the possible JSON values:.
	//
	// ```
	// { Type: <iam | no-password-required | password> Passwords: ["*****", "******"] // If Type is password.
	// }
	// ```.
	AuthenticationMode interface{} `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Indicates a password is not required for this user.
	NoPasswordRequired interface{} `field:"optional" json:"noPasswordRequired" yaml:"noPasswordRequired"`
	// Passwords used for this user.
	//
	// You can create up to two passwords for each user.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// `AWS::ElastiCache::User.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

