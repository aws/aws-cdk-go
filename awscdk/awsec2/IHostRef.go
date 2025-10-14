package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Host.
// Experimental.
type IHostRef interface {
	constructs.IConstruct
	// A reference to a Host resource.
	// Experimental.
	HostRef() *HostReference
}

// The jsii proxy for IHostRef
type jsiiProxy_IHostRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHostRef) HostRef() *HostReference {
	var returns *HostReference
	_jsii_.Get(
		j,
		"hostRef",
		&returns,
	)
	return returns
}

