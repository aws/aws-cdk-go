package awssagemaker


// The security configuration used to protect model card data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityConfigProperty := &SecurityConfigProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-securityconfig.html
//
type CfnModelCard_SecurityConfigProperty struct {
	// A AWS Key Management Service [key ID](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-id) used to encrypt a model card.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-securityconfig.html#cfn-sagemaker-modelcard-securityconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

