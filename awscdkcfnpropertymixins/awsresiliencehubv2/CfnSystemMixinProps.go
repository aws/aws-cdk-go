package awsresiliencehubv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnSystemPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSystemMixinProps := &CfnSystemMixinProps{
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html
//
type CfnSystemMixinProps struct {
	// The description of the system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key ID for encrypting system data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags assigned to the system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

