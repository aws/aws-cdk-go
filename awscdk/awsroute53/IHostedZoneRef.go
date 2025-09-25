package awsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostedZone.
// Experimental.
type IHostedZoneRef interface {
	constructs.IConstruct
	// A reference to a HostedZone resource.
	// Experimental.
	HostedZoneRef() *HostedZoneReference
}

// The jsii proxy for IHostedZoneRef
type jsiiProxy_IHostedZoneRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHostedZoneRef) HostedZoneRef() *HostedZoneReference {
	var returns *HostedZoneReference
	_jsii_.Get(
		j,
		"hostedZoneRef",
		&returns,
	)
	return returns
}

