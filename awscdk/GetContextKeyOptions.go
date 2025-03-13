package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//
//   getContextKeyOptions := &GetContextKeyOptions{
//   	Provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	IncludeEnvironment: jsii.Boolean(false),
//   	Props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
type GetContextKeyOptions struct {
	// The context provider to query.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Whether to include the stack's account and region automatically.
	// Default: true.
	//
	IncludeEnvironment *bool `field:"optional" json:"includeEnvironment" yaml:"includeEnvironment"`
	// Provider-specific properties.
	Props *map[string]interface{} `field:"optional" json:"props" yaml:"props"`
}

