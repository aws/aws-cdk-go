package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentProfile.
// Experimental.
type IEnvironmentProfileRef interface {
	constructs.IConstruct
	// A reference to a EnvironmentProfile resource.
	// Experimental.
	EnvironmentProfileRef() *EnvironmentProfileReference
}

// The jsii proxy for IEnvironmentProfileRef
type jsiiProxy_IEnvironmentProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnvironmentProfileRef) EnvironmentProfileRef() *EnvironmentProfileReference {
	var returns *EnvironmentProfileReference
	_jsii_.Get(
		j,
		"environmentProfileRef",
		&returns,
	)
	return returns
}

