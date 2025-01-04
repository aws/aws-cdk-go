package awswisdom


// Word policy config for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIGuardrailWordPolicyConfigProperty := &AIGuardrailWordPolicyConfigProperty{
//   	ManagedWordListsConfig: []interface{}{
//   		&GuardrailManagedWordsConfigProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	WordsConfig: []interface{}{
//   		&GuardrailWordConfigProperty{
//   			Text: jsii.String("text"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailwordpolicyconfig.html
//
type CfnAIGuardrail_AIGuardrailWordPolicyConfigProperty struct {
	// A config for the list of managed words.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailwordpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailwordpolicyconfig-managedwordlistsconfig
	//
	ManagedWordListsConfig interface{} `field:"optional" json:"managedWordListsConfig" yaml:"managedWordListsConfig"`
	// List of custom word configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailwordpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailwordpolicyconfig-wordsconfig
	//
	WordsConfig interface{} `field:"optional" json:"wordsConfig" yaml:"wordsConfig"`
}

