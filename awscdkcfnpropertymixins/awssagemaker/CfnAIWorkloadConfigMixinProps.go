package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAIWorkloadConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAIWorkloadConfigMixinProps := &CfnAIWorkloadConfigMixinProps{
//   	AiWorkloadConfigName: jsii.String("aiWorkloadConfigName"),
//   	AiWorkloadConfigs: &AIWorkloadConfigsProperty{
//   		WorkloadSpec: &WorkloadSpecProperty{
//   			Inline: jsii.String("inline"),
//   		},
//   	},
//   	DatasetConfig: &AIDatasetConfigProperty{
//   		InputDataConfig: []interface{}{
//   			&AIWorkloadInputDataConfigProperty{
//   				ChannelName: jsii.String("channelName"),
//   				DataSource: &AIWorkloadDataSourceProperty{
//   					S3DataSource: &AIWorkloadS3DataSourceProperty{
//   						S3Uri: jsii.String("s3Uri"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-aiworkloadconfig.html
//
type CfnAIWorkloadConfigMixinProps struct {
	// The name of the AI workload configuration.
	//
	// The name must be unique within your AWS account in the current AWS Region. Only lowercase letters and digits are accepted: DeleteAIWorkloadConfig lowercases the name before looking it up, so a name containing an uppercase letter produces a configuration that can be created and read but never deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-aiworkloadconfig.html#cfn-sagemaker-aiworkloadconfig-aiworkloadconfigname
	//
	AiWorkloadConfigName *string `field:"optional" json:"aiWorkloadConfigName" yaml:"aiWorkloadConfigName"`
	// The benchmark tool configuration for an AI workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-aiworkloadconfig.html#cfn-sagemaker-aiworkloadconfig-aiworkloadconfigs
	//
	AiWorkloadConfigs interface{} `field:"optional" json:"aiWorkloadConfigs" yaml:"aiWorkloadConfigs"`
	// The dataset configuration for an AI workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-aiworkloadconfig.html#cfn-sagemaker-aiworkloadconfig-datasetconfig
	//
	DatasetConfig interface{} `field:"optional" json:"datasetConfig" yaml:"datasetConfig"`
	// The metadata that you apply to the AI workload configuration to help you categorize and organize it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-aiworkloadconfig.html#cfn-sagemaker-aiworkloadconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

