package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FuotaTask.
// Experimental.
type IFuotaTaskRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFuotaTaskRef
type jsiiProxy_IFuotaTaskRef struct {
	internal.Type__constructsIConstruct
}

