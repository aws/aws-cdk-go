package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Farm.
// Experimental.
type IFarmRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFarmRef
type jsiiProxy_IFarmRef struct {
	internal.Type__constructsIConstruct
}

