package awsaccountaccess

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	IdentitySource: &IdentitySourceProperty{
//   		IdentityCenter: &IdentityCenterProperty{
//   			InstanceArn: jsii.String("instanceArn"),
//
//   			// the properties below are optional
//   			ApplicationArn: jsii.String("applicationArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-application.html
//
type CfnApplicationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-application.html#cfn-accountaccess-application-identitysource
	//
	IdentitySource interface{} `field:"required" json:"identitySource" yaml:"identitySource"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-application.html#cfn-accountaccess-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

