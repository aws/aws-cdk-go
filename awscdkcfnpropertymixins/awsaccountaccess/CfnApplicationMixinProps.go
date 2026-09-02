package awsaccountaccess

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnApplicationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnApplicationMixinProps := &CfnApplicationMixinProps{
//   	IdentitySource: &IdentitySourceProperty{
//   		IdentityCenter: &IdentityCenterProperty{
//   			ApplicationArn: jsii.String("applicationArn"),
//   			InstanceArn: jsii.String("instanceArn"),
//   		},
//   	},
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
type CfnApplicationMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-application.html#cfn-accountaccess-application-identitysource
	//
	IdentitySource interface{} `field:"optional" json:"identitySource" yaml:"identitySource"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-accountaccess-application.html#cfn-accountaccess-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

