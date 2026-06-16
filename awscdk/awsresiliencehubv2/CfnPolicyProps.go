package awsresiliencehubv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &CfnPolicyProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AvailabilitySlo: &AvailabilitySloProperty{
//   		Target: jsii.Number(123),
//   	},
//   	DataRecovery: &DataRecoveryTargetsProperty{
//   		TimeBetweenBackupsInMinutes: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MultiAz: &MultiAzTargetsProperty{
//   		DisasterRecoveryApproach: jsii.String("disasterRecoveryApproach"),
//   		RpoInMinutes: jsii.Number(123),
//   		RtoInMinutes: jsii.Number(123),
//   	},
//   	MultiRegion: &MultiRegionTargetsProperty{
//   		DisasterRecoveryApproach: jsii.String("disasterRecoveryApproach"),
//   		RpoInMinutes: jsii.Number(123),
//   		RtoInMinutes: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html
//
type CfnPolicyProps struct {
	// The name of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-availabilityslo
	//
	AvailabilitySlo interface{} `field:"optional" json:"availabilitySlo" yaml:"availabilitySlo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-datarecovery
	//
	DataRecovery interface{} `field:"optional" json:"dataRecovery" yaml:"dataRecovery"`
	// The description of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key ID for encrypting policy data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-multiaz
	//
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-multiregion
	//
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// Tags assigned to the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html#cfn-resiliencehubv2-policy-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

