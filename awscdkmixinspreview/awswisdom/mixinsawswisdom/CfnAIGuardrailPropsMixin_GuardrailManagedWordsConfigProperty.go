package mixinsawswisdom


// A managed words config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardrailManagedWordsConfigProperty := &GuardrailManagedWordsConfigProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailmanagedwordsconfig.html
//
type CfnAIGuardrailPropsMixin_GuardrailManagedWordsConfigProperty struct {
	// The type of guardrail managed words.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailmanagedwordsconfig.html#cfn-wisdom-aiguardrail-guardrailmanagedwordsconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

