package awsosis

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &CfnPipelineProps{
//   	MaxUnits: jsii.Number(123),
//   	MinUnits: jsii.Number(123),
//   	PipelineConfigurationBody: jsii.String("pipelineConfigurationBody"),
//   	PipelineName: jsii.String("pipelineName"),
//
//   	// the properties below are optional
//   	LogPublishingOptions: &LogPublishingOptionsProperty{
//   		CloudWatchLogDestination: &CloudWatchLogDestinationProperty{
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		IsLoggingEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcOptions: &VpcOptionsProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	MaxUnits *float64 `field:"required" json:"maxUnits" yaml:"maxUnits"`
	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	MinUnits *float64 `field:"required" json:"minUnits" yaml:"minUnits"`
	// The Data Prepper pipeline configuration in YAML format.
	PipelineConfigurationBody *string `field:"required" json:"pipelineConfigurationBody" yaml:"pipelineConfigurationBody"`
	// The name of the pipeline.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Key-value pairs that represent log publishing settings.
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// List of tags to add to the pipeline upon creation.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Options that specify the subnets and security groups for an OpenSearch Ingestion VPC endpoint.
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

