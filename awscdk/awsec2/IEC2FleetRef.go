package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EC2Fleet.
// Experimental.
type IEC2FleetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEC2FleetRef
type jsiiProxy_IEC2FleetRef struct {
	internal.Type__constructsIConstruct
}

