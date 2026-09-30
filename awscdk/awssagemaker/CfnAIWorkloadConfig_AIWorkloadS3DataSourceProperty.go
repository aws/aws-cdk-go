package awssagemaker


// The Amazon S3 data source for an AI workload.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIWorkloadS3DataSourceProperty := &AIWorkloadS3DataSourceProperty{
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloads3datasource.html
//
type CfnAIWorkloadConfig_AIWorkloadS3DataSourceProperty struct {
	// The Amazon S3 URI of the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-aiworkloadconfig-aiworkloads3datasource.html#cfn-sagemaker-aiworkloadconfig-aiworkloads3datasource-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

