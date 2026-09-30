package awssagemaker


// The dataset configuration for an AI workload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIDatasetConfigProperty := &AIDatasetConfigProperty{
//   	InputDataConfig: []interface{}{
//   		&AIWorkloadInputDataConfigProperty{
//   			ChannelName: jsii.String("channelName"),
//   			DataSource: &AIWorkloadDataSourceProperty{
//   				S3DataSource: &AIWorkloadS3DataSourceProperty{
//   					S3Uri: jsii.String("s3Uri"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aidatasetconfig.html
//
type CfnAIWorkloadConfig_AIDatasetConfigProperty struct {
	// An array of input data channel configurations for the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aidatasetconfig.html#cfn-sagemaker-aiworkloadconfig-aidatasetconfig-inputdataconfig
	//
	InputDataConfig interface{} `field:"required" json:"inputDataConfig" yaml:"inputDataConfig"`
}

