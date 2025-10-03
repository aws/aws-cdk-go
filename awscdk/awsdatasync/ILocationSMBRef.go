package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationSMB.
// Experimental.
type ILocationSMBRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationSMBRef
type jsiiProxy_ILocationSMBRef struct {
	internal.Type__constructsIConstruct
}

