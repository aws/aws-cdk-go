package awssagemaker


// The data source for an AI workload input data channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIWorkloadDataSourceProperty := &AIWorkloadDataSourceProperty{
//   	S3DataSource: &AIWorkloadS3DataSourceProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloaddatasource.html
//
type CfnAIWorkloadConfig_AIWorkloadDataSourceProperty struct {
	// The Amazon S3 data source for an AI workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloaddatasource.html#cfn-sagemaker-aiworkloadconfig-aiworkloaddatasource-s3datasource
	//
	S3DataSource interface{} `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

