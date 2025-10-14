package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DedicatedIpPool.
// Experimental.
type IDedicatedIpPoolRef interface {
	constructs.IConstruct
	// A reference to a DedicatedIpPool resource.
	// Experimental.
	DedicatedIpPoolRef() *DedicatedIpPoolReference
}

// The jsii proxy for IDedicatedIpPoolRef
type jsiiProxy_IDedicatedIpPoolRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDedicatedIpPoolRef) DedicatedIpPoolRef() *DedicatedIpPoolReference {
	var returns *DedicatedIpPoolReference
	_jsii_.Get(
		j,
		"dedicatedIpPoolRef",
		&returns,
	)
	return returns
}

