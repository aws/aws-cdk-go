package awsbackupgateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackupgateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Hypervisor.
// Experimental.
type IHypervisorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHypervisorRef
type jsiiProxy_IHypervisorRef struct {
	internal.Type__constructsIConstruct
}

