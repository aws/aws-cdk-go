package awscdklambdapythonalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Command hooks.
//
// These commands will run in the environment in which bundling occurs: inside
// the container for Docker bundling or on the host OS for local bundling.
//
// Commands are chained with `&&`.
//
// ```text
// {
//   // Run tests prior to bundling
//   beforeBundling(inputDir: string, outputDir: string): string[] {
//     return [`pytest`];
//   }
//   // ...
// }
// ```.
// Experimental.
type ICommandHooks interface {
	// Returns commands to run after bundling.
	//
	// Commands are chained with `&&`.
	// Experimental.
	AfterBundling(inputDir *string, outputDir *string) *[]*string
	// Returns commands to run before bundling.
	//
	// Commands are chained with `&&`.
	// Experimental.
	BeforeBundling(inputDir *string, outputDir *string) *[]*string
}

// The jsii proxy for ICommandHooks
type jsiiProxy_ICommandHooks struct {
	_ byte // padding
}

func (i *jsiiProxy_ICommandHooks) AfterBundling(inputDir *string, outputDir *string) *[]*string {
	if err := i.validateAfterBundlingParameters(inputDir, outputDir); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"afterBundling",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICommandHooks) BeforeBundling(inputDir *string, outputDir *string) *[]*string {
	if err := i.validateBeforeBundlingParameters(inputDir, outputDir); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"beforeBundling",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

