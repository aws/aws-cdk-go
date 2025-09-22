package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMPool.
// Experimental.
type IIPAMPoolRef interface {
	constructs.IConstruct
	// A reference to a IPAMPool resource.
	// Experimental.
	IpamPoolRef() *IPAMPoolReference
}

// The jsii proxy for IIPAMPoolRef
type jsiiProxy_IIPAMPoolRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPAMPoolRef) IpamPoolRef() *IPAMPoolReference {
	var returns *IPAMPoolReference
	_jsii_.Get(
		j,
		"ipamPoolRef",
		&returns,
	)
	return returns
}

