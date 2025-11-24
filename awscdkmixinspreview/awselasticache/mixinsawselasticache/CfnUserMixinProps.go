package mixinsawselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnUserPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var authenticationMode interface{}
//
//   cfnUserMixinProps := &CfnUserMixinProps{
//   	AccessString: jsii.String("accessString"),
//   	AuthenticationMode: authenticationMode,
//   	Engine: jsii.String("engine"),
//   	NoPasswordRequired: jsii.Boolean(false),
//   	Passwords: []*string{
//   		jsii.String("passwords"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserId: jsii.String("userId"),
//   	UserName: jsii.String("userName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html
//
type CfnUserMixinProps struct {
	// Access permissions string used for this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-accessstring
	//
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// Specifies the authentication mode to use. Below is an example of the possible JSON values:.
	//
	// ```
	// { Passwords: ["*****", "******"] // If Type is password.
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-authenticationmode
	//
	AuthenticationMode interface{} `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// The current supported values are valkey and redis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Indicates a password is not required for this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-nopasswordrequired
	//
	NoPasswordRequired interface{} `field:"optional" json:"noPasswordRequired" yaml:"noPasswordRequired"`
	// Passwords used for this user.
	//
	// You can create up to two passwords for each user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-passwords
	//
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
	// The list of tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-userid
	//
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
	// The username of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-username
	//
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

