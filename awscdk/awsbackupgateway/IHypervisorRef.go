package awsbackupgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackupgateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Hypervisor.
// Experimental.
type IHypervisorRef interface {
	constructs.IConstruct
	// A reference to a Hypervisor resource.
	// Experimental.
	HypervisorRef() *HypervisorReference
}

// The jsii proxy for IHypervisorRef
type jsiiProxy_IHypervisorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHypervisorRef) HypervisorRef() *HypervisorReference {
	var returns *HypervisorReference
	_jsii_.Get(
		j,
		"hypervisorRef",
		&returns,
	)
	return returns
}

