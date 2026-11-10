package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOmniDashboard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOmniDashboardProps := &CfnOmniDashboardProps{
//   	Body: jsii.String("body"),
//   	Name: jsii.String("name"),
//   	SpaceId: jsii.String("spaceId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-omnidashboard.html
//
type CfnOmniDashboardProps struct {
	// The dashboard definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-omnidashboard.html#cfn-cloudwatch-omnidashboard-body
	//
	Body *string `field:"required" json:"body" yaml:"body"`
	// A name that identifies the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-omnidashboard.html#cfn-cloudwatch-omnidashboard-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique ID of the space to create the dashboard in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-omnidashboard.html#cfn-cloudwatch-omnidashboard-spaceid
	//
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
	// An optional description of the dashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-omnidashboard.html#cfn-cloudwatch-omnidashboard-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-omnidashboard.html#cfn-cloudwatch-omnidashboard-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

