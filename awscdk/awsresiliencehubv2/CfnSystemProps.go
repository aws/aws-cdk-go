package awsresiliencehubv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSystem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSystemProps := &CfnSystemProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	SharingEnabled: jsii.Boolean(false),
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
type CfnSystemProps struct {
	// The name of the system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key ID for encrypting system data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Whether the system is enabled to be shared with other members of the Organization.
	//
	// Only applicable if the system owner is a management account or delegated admin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-sharingenabled
	//
	// Default: - false.
	//
	SharingEnabled interface{} `field:"optional" json:"sharingEnabled" yaml:"sharingEnabled"`
	// Tags assigned to the system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-system.html#cfn-resiliencehubv2-system-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

