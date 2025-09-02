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
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	QueryLoggingConfiguration: &QueryLoggingConfigurationProperty{
//   		Destinations: []interface{}{
//   			&LoggingDestinationProperty{
//   				CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   					LogGroupArn: jsii.String("logGroupArn"),
//   				},
//   				Filters: &LoggingFilterProperty{
//   					QspThreshold: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceConfiguration: &WorkspaceConfigurationProperty{
//   		LimitsPerLabelSets: []interface{}{
//   			&LimitsPerLabelSetProperty{
//   				LabelSet: []interface{}{
//   					&LabelProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Limits: &LimitsPerLabelSetEntryProperty{
//   					MaxSeries: jsii.Number(123),
//   				},
//   			},
//   		},
//   		RetentionPeriodInDays: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html
//
type CfnWorkspaceProps struct {
	// The alert manager definition, a YAML configuration for the alert manager in your Amazon Managed Service for Prometheus workspace.
	//
	// For details about the alert manager definition, see [Creating an alert manager configuration files](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-alertmanager-config.html) in the *Amazon Managed Service for Prometheus User Guide* .
	//
	// The following example shows part of a CloudFormation YAML file with an embedded alert manager definition (following the `- |-` ).
	//
	// `Workspace: Type: AWS::APS::Workspace .... Properties: .... AlertManagerDefinition: Fn::Sub: - |- alertmanager_config: | templates: - 'default_template' route: receiver: example-sns receivers: - name: example-sns sns_configs: - topic_arn: 'arn:aws:sns:${AWS::Region}:${AWS::AccountId}:${TopicName}' -`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-alertmanagerdefinition
	//
	AlertManagerDefinition *string `field:"optional" json:"alertManagerDefinition" yaml:"alertManagerDefinition"`
	// The alias that is assigned to this workspace to help identify it.
	//
	// It does not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// (optional) The ARN for a customer managed AWS KMS key to use for encrypting data within your workspace.
	//
	// For more information about using your own key in your workspace, see [Encryption at rest](https://docs.aws.amazon.com/prometheus/latest/userguide/encryption-at-rest-Amazon-Service-Prometheus.html) in the *Amazon Managed Service for Prometheus User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Contains information about the logging configuration for the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The definition of logging configuration in an Amazon Managed Service for Prometheus workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-queryloggingconfiguration
	//
	QueryLoggingConfiguration interface{} `field:"optional" json:"queryLoggingConfiguration" yaml:"queryLoggingConfiguration"`
	// The list of tag keys and values that are associated with the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Use this structure to define label sets and the ingestion limits for time series that match label sets, and to specify the retention period of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aps-workspace.html#cfn-aps-workspace-workspaceconfiguration
	//
	WorkspaceConfiguration interface{} `field:"optional" json:"workspaceConfiguration" yaml:"workspaceConfiguration"`
}

