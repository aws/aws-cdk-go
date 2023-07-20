package awsaps

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
//   	AlertManagerDefinition: jsii.String("alertManagerDefinition"),
//   	Alias: jsii.String("alias"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html
//
type CfnWorkspaceProps struct {
	// The alert manager definition for the workspace, as a string.
	//
	// For more information, see [Alert manager and templating](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-alertmanagerdefinition
	//
	AlertManagerDefinition *string `field:"optional" json:"alertManagerDefinition" yaml:"alertManagerDefinition"`
	// An alias that you assign to this workspace to help you identify it.
	//
	// It does not need to be unique.
	//
	// The alias can be as many as 100 characters and can include any type of characters. Amazon Managed Service for Prometheus automatically strips any blank spaces from the beginning and end of the alias that you specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The LoggingConfiguration attribute is used to set the logging configuration for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// A list of tag keys and values to associate with the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

