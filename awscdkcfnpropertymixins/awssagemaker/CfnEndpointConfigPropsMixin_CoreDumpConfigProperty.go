package awssagemaker


// Specifies where SageMaker writes core dumps from the model container when the process crashes, and how it encrypts them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   coreDumpConfigProperty := &CoreDumpConfigProperty{
//   	DestinationS3Uri: jsii.String("destinationS3Uri"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-coredumpconfig.html
//
type CfnEndpointConfigPropsMixin_CoreDumpConfigProperty struct {
	// The Amazon S3 bucket to send the core dump to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-coredumpconfig.html#cfn-sagemaker-endpointconfig-coredumpconfig-destinations3uri
	//
	DestinationS3Uri *string `field:"optional" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// The AWS Key Management Service (AWS KMS) key that SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption.
	//
	// If you use a KMS key ID or an alias of your KMS key, the SageMaker execution role must include permissions to call kms:Encrypt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-coredumpconfig.html#cfn-sagemaker-endpointconfig-coredumpconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

