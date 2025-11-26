package previewawsbedrockmixins


// The system-defined guardrail profile that you're using with your guardrail.
//
// Guardrail profiles define the destination AWS Regions where guardrail inference requests can be automatically routed. Using guardrail profiles helps maintain guardrail performance and reliability when demand increases.
//
// For more information, see the [Amazon Bedrock User Guide](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-cross-region.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardrailCrossRegionConfigProperty := &GuardrailCrossRegionConfigProperty{
//   	GuardrailProfileArn: jsii.String("guardrailProfileArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-guardrailcrossregionconfig.html
//
type CfnGuardrailPropsMixin_GuardrailCrossRegionConfigProperty struct {
	// The Amazon Resource Name (ARN) of the guardrail profile that your guardrail is using.
	//
	// Guardrail profile availability depends on your current AWS Region . For more information, see the [Amazon Bedrock User Guide](https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-cross-region-support.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-guardrailcrossregionconfig.html#cfn-bedrock-guardrail-guardrailcrossregionconfig-guardrailprofilearn
	//
	GuardrailProfileArn *string `field:"optional" json:"guardrailProfileArn" yaml:"guardrailProfileArn"`
}

