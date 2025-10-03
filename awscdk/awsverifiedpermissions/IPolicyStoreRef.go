package awsverifiedpermissions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStore.
// Experimental.
type IPolicyStoreRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPolicyStoreRef
type jsiiProxy_IPolicyStoreRef struct {
	internal.Type__constructsIConstruct
}

