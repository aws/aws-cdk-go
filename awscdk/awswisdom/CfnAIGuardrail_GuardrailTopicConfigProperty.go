package awswisdom


// Topic config in topic policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailTopicConfigProperty := &GuardrailTopicConfigProperty{
//   	Definition: jsii.String("definition"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Examples: []*string{
//   		jsii.String("examples"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html
//
type CfnAIGuardrail_GuardrailTopicConfigProperty struct {
	// Definition of topic in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-definition
	//
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Name of topic in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of topic in a policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of text examples.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailtopicconfig.html#cfn-wisdom-aiguardrail-guardrailtopicconfig-examples
	//
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
}

