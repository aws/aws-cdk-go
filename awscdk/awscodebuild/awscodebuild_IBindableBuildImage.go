package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A variant of {@link IBuildImage} that allows binding to the project.
type IBindableBuildImage interface {
	IBuildImage
	// Function that allows the build image access to the construct tree.
	Bind(scope constructs.Construct, project IProject, options *BuildImageBindOptions) *BuildImageConfig
}

// The jsii proxy for IBindableBuildImage
type jsiiProxy_IBindableBuildImage struct {
	jsiiProxy_IBuildImage
}

func (i *jsiiProxy_IBindableBuildImage) Bind(scope constructs.Construct, project IProject, options *BuildImageBindOptions) *BuildImageConfig {
	if err := i.validateBindParameters(scope, project, options); err != nil {
		panic(err)
	}
	var returns *BuildImageConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project, options},
		&returns,
	)

	return returns
}

