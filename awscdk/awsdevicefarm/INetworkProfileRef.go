package awsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkProfile.
// Experimental.
type INetworkProfileRef interface {
	constructs.IConstruct
	// A reference to a NetworkProfile resource.
	// Experimental.
	NetworkProfileRef() *NetworkProfileReference
}

// The jsii proxy for INetworkProfileRef
type jsiiProxy_INetworkProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkProfileRef) NetworkProfileRef() *NetworkProfileReference {
	var returns *NetworkProfileReference
	_jsii_.Get(
		j,
		"networkProfileRef",
		&returns,
	)
	return returns
}

