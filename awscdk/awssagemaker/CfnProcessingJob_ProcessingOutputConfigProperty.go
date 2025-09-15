package awssagemaker


// Configuration for uploading output from the processing container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processingOutputConfigProperty := &ProcessingOutputConfigProperty{
//   	Outputs: []interface{}{
//   		&ProcessingOutputsObjectProperty{
//   			OutputName: jsii.String("outputName"),
//
//   			// the properties below are optional
//   			AppManaged: jsii.Boolean(false),
//   			FeatureStoreOutput: &FeatureStoreOutputProperty{
//   				FeatureGroupName: jsii.String("featureGroupName"),
//   			},
//   			S3Output: &S3OutputProperty{
//   				S3UploadMode: jsii.String("s3UploadMode"),
//   				S3Uri: jsii.String("s3Uri"),
//
//   				// the properties below are optional
//   				LocalPath: jsii.String("localPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputconfig.html
//
type CfnProcessingJob_ProcessingOutputConfigProperty struct {
	// An array of outputs configuring the data to upload from the processing container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputconfig.html#cfn-sagemaker-processingjob-processingoutputconfig-outputs
	//
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt the processing job output.
	//
	// `KmsKeyId` can be an ID of a KMS key, ARN of a KMS key, or alias of a KMS key. The `KmsKeyId` is applied to all outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-processingoutputconfig.html#cfn-sagemaker-processingjob-processingoutputconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

