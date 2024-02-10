package awsosis

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
//   	BufferOptions: &BufferOptionsProperty{
//   		PersistentBufferEnabled: jsii.Boolean(false),
//   	},
//   	EncryptionAtRestOptions: &EncryptionAtRestOptionsProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
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
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//
//   		// the properties below are optional
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html
//
type CfnPipelineProps struct {
	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-maxunits
	//
	MaxUnits *float64 `field:"required" json:"maxUnits" yaml:"maxUnits"`
	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-minunits
	//
	MinUnits *float64 `field:"required" json:"minUnits" yaml:"minUnits"`
	// The Data Prepper pipeline configuration in YAML format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-pipelineconfigurationbody
	//
	PipelineConfigurationBody *string `field:"required" json:"pipelineConfigurationBody" yaml:"pipelineConfigurationBody"`
	// The name of the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-pipelinename
	//
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Options that specify the configuration of a persistent buffer.
	//
	// To configure how OpenSearch Ingestion encrypts this data, set the `EncryptionAtRestOptions` . For more information, see [Persistent buffering](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/osis-features-overview.html#persistent-buffering) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-bufferoptions
	//
	BufferOptions interface{} `field:"optional" json:"bufferOptions" yaml:"bufferOptions"`
	// Options to control how OpenSearch encrypts buffer data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-encryptionatrestoptions
	//
	EncryptionAtRestOptions interface{} `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// Key-value pairs that represent log publishing settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-logpublishingoptions
	//
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// List of tags to add to the pipeline upon creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Options that specify the subnets and security groups for an OpenSearch Ingestion VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-osis-pipeline.html#cfn-osis-pipeline-vpcoptions
	//
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

