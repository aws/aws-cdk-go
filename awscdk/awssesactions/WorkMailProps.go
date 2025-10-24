package awssesactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Construction properties for a WorkMail action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic Topic
//
//   workMailProps := &WorkMailProps{
//   	OrganizationArn: jsii.String("organizationArn"),
//
//   	// the properties below are optional
//   	Topic: topic,
//   }
//
type WorkMailProps struct {
	// The WorkMail organization ARN.
	//
	// Amazon WorkMail organization ARNs are in the form
	// `arn:aws:workmail:region:account_ID:organization/organization_ID`.
	//
	// Example:
	//   "arn:aws:workmail:us-east-1:123456789012:organization/m-68755160c4cb4e29a2b2f8fb58f359d7"
	//
	OrganizationArn *string `field:"required" json:"organizationArn" yaml:"organizationArn"`
	// The SNS topic to notify when the WorkMail action is taken.
	// Default: - no topic will be attached to the action.
	//
	Topic awssns.ITopic `field:"optional" json:"topic" yaml:"topic"`
}

