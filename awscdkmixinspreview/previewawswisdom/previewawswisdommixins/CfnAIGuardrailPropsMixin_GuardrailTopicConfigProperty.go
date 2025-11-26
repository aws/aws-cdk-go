package previewawswisdommixins


// Topic configuration in topic policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardrailTopicConfigProperty := &GuardrailTopicConfigProperty{
//   	Definition: jsii.String("definition"),
//   	Examples: []*string{
//   		jsii.String("examples"),
//   	},
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html
//
type CfnAIGuardrailPropsMixin_GuardrailTopicConfigProperty struct {
	// Definition of topic in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-definition
	//
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Text example in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-examples
	//
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// Name of topic in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type of topic in a policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

