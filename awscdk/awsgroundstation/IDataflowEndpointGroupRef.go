package awsgroundstation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataflowEndpointGroup.
// Experimental.
type IDataflowEndpointGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataflowEndpointGroupRef
type jsiiProxy_IDataflowEndpointGroupRef struct {
	internal.Type__constructsIConstruct
}

