package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StaticIp.
// Experimental.
type IStaticIpRef interface {
	constructs.IConstruct
	// A reference to a StaticIp resource.
	// Experimental.
	StaticIpRef() *StaticIpReference
}

// The jsii proxy for IStaticIpRef
type jsiiProxy_IStaticIpRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStaticIpRef) StaticIpRef() *StaticIpReference {
	var returns *StaticIpReference
	_jsii_.Get(
		j,
		"staticIpRef",
		&returns,
	)
	return returns
}

