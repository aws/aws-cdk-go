package awscdk


// A custom rule source for the validation engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   validationRuleSource := &ValidationRuleSource{
//   	Content: jsii.String("content"),
//   	Name: jsii.String("name"),
//   }
//
type ValidationRuleSource struct {
	// The rule content (e.g., Rego policy source code).
	Content *string `field:"required" json:"content" yaml:"content"`
	// The name of the rule source.
	Name *string `field:"required" json:"name" yaml:"name"`
}

