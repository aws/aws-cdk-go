package awsbedrockagentcore


// A root CA certificate configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateProperty := &CertificateProperty{
//   	CertificateLocation: &CertificateLocationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-certificate.html
//
type CfnCodeInterpreterCustom_CertificateProperty struct {
	// Certificate location in Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-certificate.html#cfn-bedrockagentcore-codeinterpretercustom-certificate-certificatelocation
	//
	CertificateLocation interface{} `field:"required" json:"certificateLocation" yaml:"certificateLocation"`
}

