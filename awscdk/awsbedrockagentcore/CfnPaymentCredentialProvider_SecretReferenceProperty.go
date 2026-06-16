package awsbedrockagentcore


// A reference to a customer-provided secret stored in AWS Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretReferenceProperty := &SecretReferenceProperty{
//   	JsonKey: jsii.String("jsonKey"),
//   	SecretId: jsii.String("secretId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-secretreference.html
//
type CfnPaymentCredentialProvider_SecretReferenceProperty struct {
	// The JSON key within the secret that contains the credential value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-secretreference.html#cfn-bedrockagentcore-paymentcredentialprovider-secretreference-jsonkey
	//
	JsonKey *string `field:"required" json:"jsonKey" yaml:"jsonKey"`
	// The ID or ARN of the secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentcredentialprovider-secretreference.html#cfn-bedrockagentcore-paymentcredentialprovider-secretreference-secretid
	//
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
}

