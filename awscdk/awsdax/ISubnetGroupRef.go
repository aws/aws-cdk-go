package awsdax

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdax/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetGroup.
// Experimental.
type ISubnetGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubnetGroupRef
type jsiiProxy_ISubnetGroupRef struct {
	internal.Type__constructsIConstruct
}

