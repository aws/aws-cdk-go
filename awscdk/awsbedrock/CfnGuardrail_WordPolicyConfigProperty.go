package awsbedrock


// Contains details about the word policy to configured for the guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wordPolicyConfigProperty := &WordPolicyConfigProperty{
//   	ManagedWordListsConfig: []interface{}{
//   		&ManagedWordsConfigProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	WordsConfig: []interface{}{
//   		&WordConfigProperty{
//   			Text: jsii.String("text"),
//
//   			// the properties below are optional
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-wordpolicyconfig.html
//
type CfnGuardrail_WordPolicyConfigProperty struct {
	// A list of managed words to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-wordpolicyconfig.html#cfn-bedrock-guardrail-wordpolicyconfig-managedwordlistsconfig
	//
	ManagedWordListsConfig interface{} `field:"optional" json:"managedWordListsConfig" yaml:"managedWordListsConfig"`
	// A list of words to configure for the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-wordpolicyconfig.html#cfn-bedrock-guardrail-wordpolicyconfig-wordsconfig
	//
	WordsConfig interface{} `field:"optional" json:"wordsConfig" yaml:"wordsConfig"`
}

