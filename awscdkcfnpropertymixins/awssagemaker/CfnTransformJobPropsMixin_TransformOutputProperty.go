package awssagemaker


// Describes the results of the transform job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transformOutputProperty := &TransformOutputProperty{
//   	Accept: jsii.String("accept"),
//   	AssembleWith: jsii.String("assembleWith"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	S3OutputPath: jsii.String("s3OutputPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformoutput.html
//
type CfnTransformJobPropsMixin_TransformOutputProperty struct {
	// The MIME type used to specify the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformoutput.html#cfn-sagemaker-transformjob-transformoutput-accept
	//
	Accept *string `field:"optional" json:"accept" yaml:"accept"`
	// Defines how to assemble the results of the transform job as a single S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformoutput.html#cfn-sagemaker-transformjob-transformoutput-assemblewith
	//
	AssembleWith *string `field:"optional" json:"assembleWith" yaml:"assembleWith"`
	// The AWS KMS key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformoutput.html#cfn-sagemaker-transformjob-transformoutput-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The Amazon S3 path where you want Amazon SageMaker to store the results of the transform job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-transformjob-transformoutput.html#cfn-sagemaker-transformjob-transformoutput-s3outputpath
	//
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

