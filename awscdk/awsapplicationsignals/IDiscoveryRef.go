package awsapplicationsignals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Discovery.
// Experimental.
type IDiscoveryRef interface {
	constructs.IConstruct
	// A reference to a Discovery resource.
	// Experimental.
	DiscoveryRef() *DiscoveryReference
}

// The jsii proxy for IDiscoveryRef
type jsiiProxy_IDiscoveryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDiscoveryRef) DiscoveryRef() *DiscoveryReference {
	var returns *DiscoveryReference
	_jsii_.Get(
		j,
		"discoveryRef",
		&returns,
	)
	return returns
}

