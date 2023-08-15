package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dummyValue interface{}
//   var props interface{}
//
//   getContextValueOptions := &GetContextValueOptions{
//   	DummyValue: dummyValue,
//   	Provider: jsii.String("provider"),
//
//   	// the properties below are optional
//   	IncludeEnvironment: jsii.Boolean(false),
//   	Props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   }
//
type GetContextValueOptions struct {
	// The context provider to query.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Whether to include the stack's account and region automatically.
	// Default: true.
	//
	IncludeEnvironment *bool `field:"optional" json:"includeEnvironment" yaml:"includeEnvironment"`
	// Provider-specific properties.
	Props *map[string]interface{} `field:"optional" json:"props" yaml:"props"`
	// The value to return if the context value was not found and a missing context is reported.
	//
	// This should be a dummy value that should preferably
	// fail during deployment since it represents an invalid state.
	DummyValue interface{} `field:"required" json:"dummyValue" yaml:"dummyValue"`
}

