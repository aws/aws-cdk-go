package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Theme.
// Experimental.
type IThemeRef interface {
	constructs.IConstruct
	// A reference to a Theme resource.
	// Experimental.
	ThemeRef() *ThemeReference
}

// The jsii proxy for IThemeRef
type jsiiProxy_IThemeRef struct {
	internal.Type__constructsIConstruct
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

