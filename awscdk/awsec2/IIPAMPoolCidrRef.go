package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMPoolCidr.
// Experimental.
type IIPAMPoolCidrRef interface {
	constructs.IConstruct
	// A reference to a IPAMPoolCidr resource.
	// Experimental.
	IpamPoolCidrRef() *IPAMPoolCidrReference
}

// The jsii proxy for IIPAMPoolCidrRef
type jsiiProxy_IIPAMPoolCidrRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPAMPoolCidrRef) IpamPoolCidrRef() *IPAMPoolCidrReference {
	var returns *IPAMPoolCidrReference
	_jsii_.Get(
		j,
		"ipamPoolCidrRef",
		&returns,
	)
	return returns
}

