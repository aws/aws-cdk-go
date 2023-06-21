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
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVerifiedAccessGroupProps struct {
	// The ID of the AWS Verified Access instance.
	VerifiedAccessInstanceId *string `field:"required" json:"verifiedAccessInstanceId" yaml:"verifiedAccessInstanceId"`
	// A description for the AWS Verified Access group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Verified Access policy document.
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The status of the Verified Access policy.
	PolicyEnabled interface{} `field:"optional" json:"policyEnabled" yaml:"policyEnabled"`
	// The tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

