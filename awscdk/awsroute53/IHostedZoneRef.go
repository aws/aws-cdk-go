package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostedZone.
// Experimental.
type IHostedZoneRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHostedZoneRef
type jsiiProxy_IHostedZoneRef struct {
	internal.Type__constructsIConstruct
}

