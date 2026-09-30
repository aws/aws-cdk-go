package awsbedrockagentcore


// Validation rules for namespace variable values.
//
// Multiple rules can be specified and all must pass.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceKeyValidationProperty := &NamespaceKeyValidationProperty{
//   	AllowedValues: []*string{
//   		jsii.String("allowedValues"),
//   	},
//   	RegexPattern: jsii.String("regexPattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-namespacekeyvalidation.html
//
type CfnMemory_NamespaceKeyValidationProperty struct {
	// List of allowed values for this namespace variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-namespacekeyvalidation.html#cfn-bedrockagentcore-memory-namespacekeyvalidation-allowedvalues
	//
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// A regex pattern that a namespace variable value must match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-namespacekeyvalidation.html#cfn-bedrockagentcore-memory-namespacekeyvalidation-regexpattern
	//
	RegexPattern *string `field:"optional" json:"regexPattern" yaml:"regexPattern"`
}

