package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVerifiedAccessGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVerifiedAccessGroupProps := &CfnVerifiedAccessGroupProps{
//   	VerifiedAccessInstanceId: jsii.String("verifiedAccessInstanceId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyEnabled: jsii.Boolean(false),
//   	SseSpecification: &SseSpecificationProperty{
//   		CustomerManagedKeyEnabled: jsii.Boolean(false),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html
//
type CfnVerifiedAccessGroupProps struct {
	// The ID of the AWS Verified Access instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html#cfn-ec2-verifiedaccessgroup-verifiedaccessinstanceid
	//
	VerifiedAccessInstanceId *string `field:"required" json:"verifiedAccessInstanceId" yaml:"verifiedAccessInstanceId"`
	// A description for the AWS Verified Access group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html#cfn-ec2-verifiedaccessgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Verified Access policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html#cfn-ec2-verifiedaccessgroup-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The status of the Verified Access policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html#cfn-ec2-verifiedaccessgroup-policyenabled
	//
	PolicyEnabled interface{} `field:"optional" json:"policyEnabled" yaml:"policyEnabled"`
	// The options for additional server side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html#cfn-ec2-verifiedaccessgroup-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessgroup.html#cfn-ec2-verifiedaccessgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

