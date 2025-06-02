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
//   	IgnoreErrorOnMissingContext: jsii.Boolean(false),
//   	IncludeEnvironment: jsii.Boolean(false),
//   	MustExist: jsii.Boolean(false),
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
	// The value to return if the lookup has not yet been performed.
	//
	// Upon first synthesis, the lookups has not yet been performed. The
	// `getValue()` operation returns this value instead, so that synthesis can
	// proceed. After synthesis completes the first time, the actual lookup will
	// be performed and synthesis will run again with the *real* value.
	//
	// Dummy values should preferably have valid shapes so that downstream
	// consumers of lookup values don't throw validation exceptions if they
	// encounter a dummy value (or all possible downstream consumers need to
	// effectively check for the well-known shape of the dummy value); throwing an
	// exception would error out the synthesis operation and prevent the lookup
	// and the second, real, synthesis from happening.
	//
	// ## Connection to mustExist
	//
	// `dummyValue` is also used as the official value to return if the lookup has
	// failed and `mustExist == false`.
	DummyValue interface{} `field:"required" json:"dummyValue" yaml:"dummyValue"`
	// Ignore a lookup failure and return the `dummyValue` instead.
	//
	// `mustExist` is the recommended alias for this deprecated
	// property (note that its value is reversed).
	// Default: false.
	//
	// Deprecated: Use mustExist instead.
	IgnoreErrorOnMissingContext *bool `field:"optional" json:"ignoreErrorOnMissingContext" yaml:"ignoreErrorOnMissingContext"`
	// Whether the resource must exist.
	//
	// If this is set (the default), the query fails if the value or resource we
	// tried to look up doesn't exist.
	//
	// If this is `false` and the value we tried to look up could not be found, the
	// failure is suppressed and `dummyValue` is officially returned instead.
	//
	// When this happens, `dummyValue` is encoded into cached context and it will
	// never be refreshed anymore until the user runs `cdk context --reset <key>`.
	//
	// Note that it is not possible for the CDK app code to make a distinction
	// between "the lookup has not been performed yet" and "the lookup didn't
	// find anything and we returned a default value instead".
	//
	// ## Context providers
	//
	// This feature must explicitly be supported by context providers. It is
	// currently supported by:
	//
	// - KMS key provider
	// - SSM parameter provider
	//
	// ## Note to implementors
	//
	// The dummy value should not be returned for all SDK lookup failures. For
	// example, "no network" or "no credentials" or "malformed query" should
	// not lead to the dummy value being returned. Only the case of "no such
	// resource" should.
	// Default: true.
	//
	MustExist *bool `field:"optional" json:"mustExist" yaml:"mustExist"`
}

