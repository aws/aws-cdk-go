package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnUserProps := &cfnUserProps{
//   	accessString: jsii.String("accessString"),
//   	authenticationMode: authenticationMode,
//   	userName: jsii.String("userName"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnUserProps struct {
	// Access permissions string used for this user.
	AccessString *string `field:"required" json:"accessString" yaml:"accessString"`
	// Denotes whether the user requires a password to authenticate.
	//
	// *Example:*
	//
	// `mynewdbuser: Type: AWS::MemoryDB::User Properties: AccessString: on ~* &* +@all AuthenticationMode: Passwords: '1234567890123456' Type: password UserName: mynewdbuser AuthenticationMode: { "Passwords": ["1234567890123456"], "Type": "Password" }`.
	AuthenticationMode interface{} `field:"required" json:"authenticationMode" yaml:"authenticationMode"`
	// The name of the user.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

