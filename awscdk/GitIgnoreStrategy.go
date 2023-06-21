package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   gitIgnoreStrategy := cdk.NewGitIgnoreStrategy(jsii.String("absoluteRootPath"), []*string{
//   	jsii.String("patterns"),
//   })
//
type GitIgnoreStrategy interface {
	IgnoreStrategy
	// Adds another pattern.
	Add(pattern *string)
	// Determines whether a given file path should be ignored or not.
	//
	// Returns: `true` if the file should be ignored.
	Ignores(absoluteFilePath *string) *bool
}

// The jsii proxy struct for GitIgnoreStrategy
type jsiiProxy_GitIgnoreStrategy struct {
	jsiiProxy_IgnoreStrategy
}

func NewGitIgnoreStrategy(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	if err := validateNewGitIgnoreStrategyParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitIgnoreStrategy{}

	_jsii_.Create(
		"aws-cdk-lib.GitIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		&j,
	)

	return &j
}

func NewGitIgnoreStrategy_Override(g GitIgnoreStrategy, absoluteRootPath *string, patterns *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.GitIgnoreStrategy",
		[]interface{}{absoluteRootPath, patterns},
		g,
	)
}

// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
//
// Returns: `DockerIgnorePattern` associated with the given patterns.
func GitIgnoreStrategy_Docker(absoluteRootPath *string, patterns *[]*string) DockerIgnoreStrategy {
	_init_.Initialize()

	if err := validateGitIgnoreStrategy_DockerParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns DockerIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"docker",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Creates an IgnoreStrategy based on the `ignoreMode` and `exclude` in a `CopyOptions`.
//
// Returns: `IgnoreStrategy` based on the `CopyOptions`.
func GitIgnoreStrategy_FromCopyOptions(options *CopyOptions, absoluteRootPath *string) IgnoreStrategy {
	_init_.Initialize()

	if err := validateGitIgnoreStrategy_FromCopyOptionsParameters(options, absoluteRootPath); err != nil {
		panic(err)
	}
	var returns IgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"fromCopyOptions",
		[]interface{}{options, absoluteRootPath},
		&returns,
	)

	return returns
}

// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
//
// Returns: `GitIgnorePattern` associated with the given patterns.
func GitIgnoreStrategy_Git(absoluteRootPath *string, patterns *[]*string) GitIgnoreStrategy {
	_init_.Initialize()

	if err := validateGitIgnoreStrategy_GitParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns GitIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"git",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

// Ignores file paths based on simple glob patterns.
//
// Returns: `GlobIgnorePattern` associated with the given patterns.
func GitIgnoreStrategy_Glob(absoluteRootPath *string, patterns *[]*string) GlobIgnoreStrategy {
	_init_.Initialize()

	if err := validateGitIgnoreStrategy_GlobParameters(absoluteRootPath, patterns); err != nil {
		panic(err)
	}
	var returns GlobIgnoreStrategy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.GitIgnoreStrategy",
		"glob",
		[]interface{}{absoluteRootPath, patterns},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitIgnoreStrategy) Add(pattern *string) {
	if err := g.validateAddParameters(pattern); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"add",
		[]interface{}{pattern},
	)
}

func (g *jsiiProxy_GitIgnoreStrategy) Ignores(absoluteFilePath *string) *bool {
	if err := g.validateIgnoresParameters(absoluteFilePath); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		g,
		"ignores",
		[]interface{}{absoluteFilePath},
		&returns,
	)

	return returns
}

