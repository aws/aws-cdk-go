package awsbedrockagentcore


// Certificate location in Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   certificateLocationProperty := &CertificateLocationProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-certificatelocation.html
//
type CfnCodeInterpreterCustomPropsMixin_CertificateLocationProperty struct {
	// Secrets Manager secret ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-codeinterpretercustom-certificatelocation.html#cfn-bedrockagentcore-codeinterpretercustom-certificatelocation-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

