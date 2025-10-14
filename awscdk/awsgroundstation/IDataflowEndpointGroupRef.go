package awsgroundstation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataflowEndpointGroup.
// Experimental.
type IDataflowEndpointGroupRef interface {
	constructs.IConstruct
	// A reference to a DataflowEndpointGroup resource.
	// Experimental.
	DataflowEndpointGroupRef() *DataflowEndpointGroupReference
}

// The jsii proxy for IDataflowEndpointGroupRef
type jsiiProxy_IDataflowEndpointGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataflowEndpointGroupRef) DataflowEndpointGroupRef() *DataflowEndpointGroupReference {
	var returns *DataflowEndpointGroupReference
	_jsii_.Get(
		j,
		"dataflowEndpointGroupRef",
		&returns,
	)
	return returns
}

