package awscloudwatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnView`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnViewProps := &CfnViewProps{
//   	Definition: jsii.String("definition"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-view.html
//
type CfnViewProps struct {
	// The SQL query that defines the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-view.html#cfn-cloudwatch-view-definition
	//
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// A description of the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-view.html#cfn-cloudwatch-view-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the view.
	//
	// Must begin with the "view." prefix followed by lowercase alphanumeric characters, hyphens, and underscores. View names must be unique within the account and region. If not specified, a name is generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-view.html#cfn-cloudwatch-view-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-view.html#cfn-cloudwatch-view-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

