package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Features that are implemented behind a flag in order to preserve backwards compatibility for existing apps.
//
// The list of flags are available in the
// `aws-cdk-lib/cx-api` module.
//
// The state of the flag for this application is stored as a CDK context variable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   featureFlags := cdk.FeatureFlags_Of(this)
//
type FeatureFlags interface {
	// Check whether a feature flag is enabled.
	//
	// If configured, the flag is present in
	// the construct node context. Falls back to the defaults defined in the `cx-api`
	// module.
	IsEnabled(featureFlag *string) *bool
}

// The jsii proxy struct for FeatureFlags
type jsiiProxy_FeatureFlags struct {
	_ byte // padding
}

// Inspect feature flags on the construct node's context.
func FeatureFlags_Of(scope constructs.IConstruct) FeatureFlags {
	_init_.Initialize()

	if err := validateFeatureFlags_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns FeatureFlags

	_jsii_.StaticInvoke(
		"aws-cdk-lib.FeatureFlags",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FeatureFlags) IsEnabled(featureFlag *string) *bool {
	if err := f.validateIsEnabledParameters(featureFlag); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		f,
		"isEnabled",
		[]interface{}{featureFlag},
		&returns,
	)

	return returns
}

