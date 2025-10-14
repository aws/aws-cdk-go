package awsssmguiconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmguiconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Preferences.
// Experimental.
type IPreferencesRef interface {
	constructs.IConstruct
	// A reference to a Preferences resource.
	// Experimental.
	PreferencesRef() *PreferencesReference
}

// The jsii proxy for IPreferencesRef
type jsiiProxy_IPreferencesRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPreferencesRef) PreferencesRef() *PreferencesReference {
	var returns *PreferencesReference
	_jsii_.Get(
		j,
		"preferencesRef",
		&returns,
	)
	return returns
}

