package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionGroup.
// Experimental.
type IConnectionGroupRef interface {
	constructs.IConstruct
	// A reference to a ConnectionGroup resource.
	// Experimental.
	ConnectionGroupRef() *ConnectionGroupReference
}

// The jsii proxy for IConnectionGroupRef
type jsiiProxy_IConnectionGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConnectionGroupRef) ConnectionGroupRef() *ConnectionGroupReference {
	var returns *ConnectionGroupReference
	_jsii_.Get(
		j,
		"connectionGroupRef",
		&returns,
	)
	return returns
}

