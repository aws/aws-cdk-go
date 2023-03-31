// An experiment to bundle the entire CDK into a single module
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//
//   getContextKeyOptions := &getContextKeyOptions{
//   	provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	includeEnvironment: jsii.Boolean(false),
//   	props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
// Experimental.
type GetContextKeyOptions struct {
	// The context provider to query.
	// Experimental.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Whether to include the stack's account and region automatically.
	// Experimental.
	IncludeEnvironment *bool `field:"optional" json:"includeEnvironment" yaml:"includeEnvironment"`
	// Provider-specific properties.
	// Experimental.
	Props *map[string]interface{} `field:"optional" json:"props" yaml:"props"`
}

