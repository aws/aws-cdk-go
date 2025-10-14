package awsrefactorspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrefactorspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Environment.
// Experimental.
type IEnvironmentRef interface {
	constructs.IConstruct
	// A reference to a Environment resource.
	// Experimental.
	EnvironmentRef() *EnvironmentReference
}

// The jsii proxy for IEnvironmentRef
type jsiiProxy_IEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnvironmentRef) EnvironmentRef() *EnvironmentReference {
	var returns *EnvironmentReference
	_jsii_.Get(
		j,
		"environmentRef",
		&returns,
	)
	return returns
}

