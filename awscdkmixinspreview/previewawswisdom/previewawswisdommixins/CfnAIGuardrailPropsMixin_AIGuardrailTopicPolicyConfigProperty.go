package previewawswisdommixins


// Topic policy configuration for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aIGuardrailTopicPolicyConfigProperty := &AIGuardrailTopicPolicyConfigProperty{
//   	TopicsConfig: []interface{}{
//   		&GuardrailTopicConfigProperty{
//   			Definition: jsii.String("definition"),
//   			Examples: []*string{
//   				jsii.String("examples"),
//   			},
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailtopicpolicyconfig.html
//
type CfnAIGuardrailPropsMixin_AIGuardrailTopicPolicyConfigProperty struct {
	// List of topic configs in topic policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailtopicpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailtopicpolicyconfig-topicsconfig
	//
	TopicsConfig interface{} `field:"optional" json:"topicsConfig" yaml:"topicsConfig"`
}

