package cxapi


// Backwards compatibility for when `MissingContext` was defined here.
//
// This is necessary because its used as an input in the stable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var props interface{}
//
//   missingContext := &MissingContext{
//   	Key: jsii.String("key"),
//   	Props: map[string]interface{}{
//   		"propsKey": props,
//   	},
//   	Provider: jsii.String("provider"),
//   }
//
// See: core.Stack.reportMissingContext
//
// Deprecated: moved to package 'cloud-assembly-schema'.
type MissingContext struct {
	// The missing context key.
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Key *string `field:"required" json:"key" yaml:"key"`
	// A set of provider-specific options.
	//
	// (This is the old untyped definition, which is necessary for backwards compatibility.
	// See cxschema for a type definition.)
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Props *map[string]interface{} `field:"required" json:"props" yaml:"props"`
	// The provider from which we expect this context key to be obtained.
	//
	// (This is the old untyped definition, which is necessary for backwards compatibility.
	// See cxschema for a type definition.)
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
}

