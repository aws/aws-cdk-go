package awscontroltower

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledControl.
// Experimental.
type IEnabledControlRef interface {
	constructs.IConstruct
	// A reference to a EnabledControl resource.
	// Experimental.
	EnabledControlRef() *EnabledControlReference
}

// The jsii proxy for IEnabledControlRef
type jsiiProxy_IEnabledControlRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnabledControlRef) EnabledControlRef() *EnabledControlReference {
	var returns *EnabledControlReference
	_jsii_.Get(
		j,
		"enabledControlRef",
		&returns,
	)
	return returns
}

