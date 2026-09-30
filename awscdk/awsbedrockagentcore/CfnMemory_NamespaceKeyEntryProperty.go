package awsbedrockagentcore


// A namespace variable key definition with optional validation rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceKeyEntryProperty := &NamespaceKeyEntryProperty{
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Validation: &NamespaceKeyValidationProperty{
//   		AllowedValues: []*string{
//   			jsii.String("allowedValues"),
//   		},
//   		RegexPattern: jsii.String("regexPattern"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-namespacekeyentry.html
//
type CfnMemory_NamespaceKeyEntryProperty struct {
	// A namespace variable key name.
	//
	// Must start with a lowercase letter and contain only lowercase alphanumeric characters. Cannot be a built-in variable (memoryStrategyId, sessionId, actorId).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-namespacekeyentry.html#cfn-bedrockagentcore-memory-namespacekeyentry-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// Validation rules for namespace variable values.
	//
	// Multiple rules can be specified and all must pass.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-namespacekeyentry.html#cfn-bedrockagentcore-memory-namespacekeyentry-validation
	//
	Validation interface{} `field:"optional" json:"validation" yaml:"validation"`
}

