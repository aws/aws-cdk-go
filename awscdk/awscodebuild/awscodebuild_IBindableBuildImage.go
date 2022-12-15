package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A variant of {@link IBuildImage} that allows binding to the project.
// Experimental.
type IBindableBuildImage interface {
	IBuildImage
	// Function that allows the build image access to the construct tree.
	// Experimental.
	Bind(scope awscdk.Construct, project IProject, options *BuildImageBindOptions) *BuildImageConfig
}

// The jsii proxy for IBindableBuildImage
type jsiiProxy_IBindableBuildImage struct {
	jsiiProxy_IBuildImage
}

func (i *jsiiProxy_IBindableBuildImage) Bind(scope awscdk.Construct, project IProject, options *BuildImageBindOptions) *BuildImageConfig {
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

