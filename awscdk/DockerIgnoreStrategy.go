package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerIgnoreStrategy := cdk.NewDockerIgnoreStrategy(jsii.String("absoluteRootPath"), []*string{
//   	jsii.String("patterns"),
//   })
//
type DockerIgnoreStrategy interface {
	IgnoreStrategy
	// Adds another pattern.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for DockerIgnoreStrategy
type jsiiProxy_DockerIgnoreStrategy struct {
	jsiiProxy_IgnoreStrategy
}

func NewDockerIgnoreStrategy(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	if err := validateNewDockerIgnoreStrategyParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerIgnoreStrategy{}

	_jsii_.Create(
		"aws-cdk-lib.DockerIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		&j,
	)

	return &j
}

func NewDockerIgnoreStrategy_Override(d DockerIgnoreStrategy, absoluteRootPath *string, patterns *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DockerIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		d,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
func DockerIgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	if err := validateDockerIgnoreStrategy_DockerParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
func DockerIgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	if err := validateDockerIgnoreStrategy_FromCopyOptionsParameters(options, absoluteRootPath); err != nil {
		panic(err)
	}
	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
func DockerIgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	if err := validateDockerIgnoreStrategy_GitParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
func DockerIgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	if err := validateDockerIgnoreStrategy_GlobParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerIgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerIgnoreStrategy) Add(pattern *string) {
	if err := d.validateAddParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"add",
		[]interface{}{pattern},
	)
}

func (d *jsiiProxy_DockerIgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	if err := d.validateIgnoresParameters(absoluteFilePath); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		d,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

