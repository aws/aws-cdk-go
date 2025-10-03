package awsverifiedpermissions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsverifiedpermissions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Policy.
// Experimental.
type IPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPolicyRef
type jsiiProxy_IPolicyRef struct {
	internal.Type__constructsIConstruct
}

