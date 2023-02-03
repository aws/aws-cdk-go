// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Local bundling.
type ILocalBundling interface {
	// This method is called before attempting docker bundling to allow the bundler to be executed locally.
	//
	// If the local bundler exists, and bundling
	// was performed locally, return `true`. Otherwise, return `false`.
	TryBundle(outputDir *string, options *BundlingOptions) *bool
}

// The jsii proxy for ILocalBundling
type jsiiProxy_ILocalBundling struct {
	_ byte // padding
}

func (i *jsiiProxy_ILocalBundling) TryBundle(outputDir *string, options *BundlingOptions) *bool {
	if err := i.validateTryBundleParameters(outputDir, options); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"tryBundle",
		[]interface{}{outputDir, options},
		&returns,
	)

	return returns
}

