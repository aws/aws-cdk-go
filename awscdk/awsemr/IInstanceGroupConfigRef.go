package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceGroupConfig.
// Experimental.
type IInstanceGroupConfigRef interface {
	constructs.IConstruct
	// A reference to a InstanceGroupConfig resource.
	// Experimental.
	InstanceGroupConfigRef() *InstanceGroupConfigReference
}

// The jsii proxy for IInstanceGroupConfigRef
type jsiiProxy_IInstanceGroupConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceGroupConfigRef) InstanceGroupConfigRef() *InstanceGroupConfigReference {
	var returns *InstanceGroupConfigReference
	_jsii_.Get(
		j,
		"instanceGroupConfigRef",
		&returns,
	)
	return returns
}

