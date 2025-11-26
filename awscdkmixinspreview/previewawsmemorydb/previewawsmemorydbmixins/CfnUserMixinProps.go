package previewawsmemorydbmixins

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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserName: jsii.String("userName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-user.html
//
type CfnUserMixinProps struct {
	// Access permissions string used for this user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-user.html#cfn-memorydb-user-accessstring
	//
	AccessString *string `field:"optional" json:"accessString" yaml:"accessString"`
	// Denotes whether the user requires a password to authenticate.
	//
	// *Example:*
	//
	// `mynewdbuser: Type: AWS::MemoryDB::User Properties: AccessString: on ~* &* +@all AuthenticationMode: Passwords: '1234567890123456' Type: password UserName: mynewdbuser AuthenticationMode: { "Passwords": ["1234567890123456"], "Type": "Password" }`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-user.html#cfn-memorydb-user-authenticationmode
	//
	AuthenticationMode interface{} `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-user.html#cfn-memorydb-user-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-user.html#cfn-memorydb-user-username
	//
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

