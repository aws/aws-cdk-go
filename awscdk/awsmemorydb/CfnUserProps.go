package awsmemorydb

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
//   	UserName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	AccessString: jsii.String("accessString"),
//   	AuthenticationMode: authenticationMode,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserProps struct {
	// The name of the user.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Access permissions string used for this user.
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// Denotes whether the user requires a password to authenticate.
	//
	// *Example:*
	//
	// `mynewdbuser: Type: AWS::MemoryDB::User Properties: AccessString: on ~* &* +@all AuthenticationMode: Passwords: '1234567890123456' Type: password UserName: mynewdbuser AuthenticationMode: { "Passwords": ["1234567890123456"], "Type": "Password" }`.
	AuthenticationMode interface{} `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

