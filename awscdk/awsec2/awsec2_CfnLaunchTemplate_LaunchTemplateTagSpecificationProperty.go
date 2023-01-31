package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Specifies the tags to apply to the launch template during creation.
//
// `LaunchTemplateTagSpecification` is a property of [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateTagSpecificationProperty := &launchTemplateTagSpecificationProperty{
//   	resourceType: jsii.String("resourceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLaunchTemplate_LaunchTemplateTagSpecificationProperty struct {
	// The type of resource.
	//
	// To tag the launch template, `ResourceType` must be `launch-template` .
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags for the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

