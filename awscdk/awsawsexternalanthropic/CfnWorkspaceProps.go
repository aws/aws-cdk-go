package awsawsexternalanthropic

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &CfnWorkspaceProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DataResidency: &DataResidencyProperty{
//   		AllowedInferenceGeos: []*string{
//   			jsii.String("allowedInferenceGeos"),
//   		},
//   		DefaultInferenceGeo: jsii.String("defaultInferenceGeo"),
//   		WorkspaceGeo: jsii.String("workspaceGeo"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-awsexternalanthropic-workspace.html
//
type CfnWorkspaceProps struct {
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-awsexternalanthropic-workspace.html#cfn-awsexternalanthropic-workspace-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Data residency configuration for the workspace.
	//
	// WorkspaceGeo is immutable after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-awsexternalanthropic-workspace.html#cfn-awsexternalanthropic-workspace-dataresidency
	//
	DataResidency interface{} `field:"optional" json:"dataResidency" yaml:"dataResidency"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-awsexternalanthropic-workspace.html#cfn-awsexternalanthropic-workspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

