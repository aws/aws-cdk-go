// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents file path ignoring behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ignoreStrategy := monocdk.IgnoreStrategy_FromCopyOptions(&CopyOptions{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Follow: monocdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: monocdk.IgnoreMode_GLOB,
//   }, jsii.String("absoluteRootPath"))
//
// Experimental.
type IgnoreStrategy interface {
	// Adds another pattern.
	// Experimental.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	// Experimental.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for IgnoreStrategy
type jsiiProxy_IgnoreStrategy struct {
	_ byte // padding
}

// Experimental.
func NewIgnoreStrategy_Override(i IgnoreStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.IgnoreStrategy",
		nil, // no parameters
		i,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
// Experimental.
func IgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	if err := validateIgnoreStrategy_DockerParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"monocdk.IgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
// Experimental.
func IgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	if err := validateIgnoreStrategy_FromCopyOptionsParameters(options, absoluteRootPath); err != nil {
		panic(err)
	}
	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"monocdk.IgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
// Experimental.
func IgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	if err := validateIgnoreStrategy_GitParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"monocdk.IgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
// Experimental.
func IgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	if err := validateIgnoreStrategy_GlobParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"monocdk.IgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IgnoreStrategy) Add(pattern *string) {
	if err := i.validateAddParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"add",
		[]interface{}{pattern},
	)
}

func (i *jsiiProxy_IgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	if err := i.validateIgnoresParameters(absoluteFilePath); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

