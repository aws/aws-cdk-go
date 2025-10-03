package awsssmguiconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmguiconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Preferences.
// Experimental.
type IPreferencesRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPreferencesRef
type jsiiProxy_IPreferencesRef struct {
	internal.Type__constructsIConstruct
}

