package awscontroltower

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LandingZone.
// Experimental.
type ILandingZoneRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILandingZoneRef
type jsiiProxy_ILandingZoneRef struct {
	internal.Type__constructsIConstruct
}

