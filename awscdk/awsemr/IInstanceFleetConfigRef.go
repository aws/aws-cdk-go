package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceFleetConfig.
// Experimental.
type IInstanceFleetConfigRef interface {
	constructs.IConstruct
	// A reference to a InstanceFleetConfig resource.
	// Experimental.
	InstanceFleetConfigRef() *InstanceFleetConfigReference
}

// The jsii proxy for IInstanceFleetConfigRef
type jsiiProxy_IInstanceFleetConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceFleetConfigRef) InstanceFleetConfigRef() *InstanceFleetConfigReference {
	var returns *InstanceFleetConfigReference
	_jsii_.Get(
		j,
		"instanceFleetConfigRef",
		&returns,
	)
	return returns
}

