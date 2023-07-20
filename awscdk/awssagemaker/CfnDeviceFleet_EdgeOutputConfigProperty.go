package awssagemaker


// The output configuration for storing sample data collected by the fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   edgeOutputConfigProperty := &EdgeOutputConfigProperty{
//   	S3OutputLocation: jsii.String("s3OutputLocation"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-devicefleet-edgeoutputconfig.html
//
type CfnDeviceFleet_EdgeOutputConfigProperty struct {
	// The Amazon Simple Storage (S3) bucket URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-devicefleet-edgeoutputconfig.html#cfn-sagemaker-devicefleet-edgeoutputconfig-s3outputlocation
	//
	S3OutputLocation *string `field:"required" json:"s3OutputLocation" yaml:"s3OutputLocation"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume after compilation job.
	//
	// If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your role's account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-devicefleet-edgeoutputconfig.html#cfn-sagemaker-devicefleet-edgeoutputconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

