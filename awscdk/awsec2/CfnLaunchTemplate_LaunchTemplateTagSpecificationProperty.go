package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   launchTemplateTagSpecificationProperty := &LaunchTemplateTagSpecificationProperty{
//   	ResourceType: jsii.String("resourceType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatetagspecification.html
//
type CfnLaunchTemplate_LaunchTemplateTagSpecificationProperty struct {
	// The type of resource.
	//
	// To tag the launch template, `ResourceType` must be `launch-template` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatetagspecification.html#cfn-ec2-launchtemplate-launchtemplatetagspecification-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatetagspecification.html#cfn-ec2-launchtemplate-launchtemplatetagspecification-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

