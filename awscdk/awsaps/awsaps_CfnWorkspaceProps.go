package awsaps

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &cfnWorkspaceProps{
//   	alertManagerDefinition: jsii.String("alertManagerDefinition"),
//   	alias: jsii.String("alias"),
//   	loggingConfiguration: &loggingConfigurationProperty{
//   		logGroupArn: jsii.String("logGroupArn"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWorkspaceProps struct {
	// The alert manager definition for the workspace, as a string.
	//
	// For more information, see [Alert manager and templating](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alert-manager.html) .
	AlertManagerDefinition *string `field:"optional" json:"alertManagerDefinition" yaml:"alertManagerDefinition"`
	// An alias that you assign to this workspace to help you identify it.
	//
	// It does not need to be unique.
	//
	// The alias can be as many as 100 characters and can include any type of characters. Amazon Managed Service for Prometheus automatically strips any blank spaces from the beginning and end of the alias that you specify.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// `AWS::APS::Workspace.LoggingConfiguration`.
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// A list of tag keys and values to associate with the workspace.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

