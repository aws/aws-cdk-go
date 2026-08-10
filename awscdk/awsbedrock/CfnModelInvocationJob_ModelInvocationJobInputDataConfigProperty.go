package awsbedrock


// Details about the location of the input to the batch inference job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInvocationJobInputDataConfigProperty := &ModelInvocationJobInputDataConfigProperty{
//   	S3InputDataConfig: &ModelInvocationJobS3InputDataConfigProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		S3BucketOwner: jsii.String("s3BucketOwner"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobinputdataconfig.html
//
type CfnModelInvocationJob_ModelInvocationJobInputDataConfigProperty struct {
	// Contains the configuration of the S3 location of the input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobinputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjobinputdataconfig-s3inputdataconfig
	//
	S3InputDataConfig interface{} `field:"required" json:"s3InputDataConfig" yaml:"s3InputDataConfig"`
}

