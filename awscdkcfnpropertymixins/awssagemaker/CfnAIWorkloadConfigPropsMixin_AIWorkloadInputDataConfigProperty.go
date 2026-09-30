package awssagemaker


// A channel of input data for an AI workload configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   aIWorkloadInputDataConfigProperty := &AIWorkloadInputDataConfigProperty{
//   	ChannelName: jsii.String("channelName"),
//   	DataSource: &AIWorkloadDataSourceProperty{
//   		S3DataSource: &AIWorkloadS3DataSourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloadinputdataconfig.html
//
type CfnAIWorkloadConfigPropsMixin_AIWorkloadInputDataConfigProperty struct {
	// The logical name for the data channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloadinputdataconfig.html#cfn-sagemaker-aiworkloadconfig-aiworkloadinputdataconfig-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The data source for an AI workload input data channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloadinputdataconfig.html#cfn-sagemaker-aiworkloadconfig-aiworkloadinputdataconfig-datasource
	//
	DataSource interface{} `field:"optional" json:"dataSource" yaml:"dataSource"`
}

