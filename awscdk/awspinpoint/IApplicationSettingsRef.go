package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationSettings.
// Experimental.
type IApplicationSettingsRef interface {
	constructs.IConstruct
	// A reference to a ApplicationSettings resource.
	// Experimental.
	ApplicationSettingsRef() *ApplicationSettingsReference
}

// The jsii proxy for IApplicationSettingsRef
type jsiiProxy_IApplicationSettingsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationSettingsRef) ApplicationSettingsRef() *ApplicationSettingsReference {
	var returns *ApplicationSettingsReference
	_jsii_.Get(
		j,
		"applicationSettingsRef",
		&returns,
	)
	return returns
}

