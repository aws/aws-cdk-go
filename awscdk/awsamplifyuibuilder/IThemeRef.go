package awsamplifyuibuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Theme.
// Experimental.
type IThemeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Theme resource.
	// Experimental.
	ThemeRef() *ThemeReference
}

// The jsii proxy for IThemeRef
type jsiiProxy_IThemeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IThemeRef) ThemeRef() *ThemeReference {
	var returns *ThemeReference
	_jsii_.Get(
		j,
		"themeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThemeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThemeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

