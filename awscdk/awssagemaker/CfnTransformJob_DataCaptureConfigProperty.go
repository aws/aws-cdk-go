package awssagemaker


// Configuration to control how SageMaker captures inference data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCaptureConfigProperty := &DataCaptureConfigProperty{
//   	DestinationS3Uri: jsii.String("destinationS3Uri"),
//
//   	// the properties below are optional
//   	GenerateInferenceId: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-datacaptureconfig.html
//
type CfnTransformJob_DataCaptureConfigProperty struct {
	// The Amazon S3 location being used to capture the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-datacaptureconfig.html#cfn-sagemaker-transformjob-datacaptureconfig-destinations3uri
	//
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Flag that indicates whether to append inference id to the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-datacaptureconfig.html#cfn-sagemaker-transformjob-datacaptureconfig-generateinferenceid
	//
	GenerateInferenceId interface{} `field:"optional" json:"generateInferenceId" yaml:"generateInferenceId"`
	// The ARN of a KMS key that SageMaker uses to encrypt data on the storage volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-datacaptureconfig.html#cfn-sagemaker-transformjob-datacaptureconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

