package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingGroup.
// Experimental.
type IThingGroupRef interface {
	constructs.IConstruct
	// A reference to a ThingGroup resource.
	// Experimental.
	ThingGroupRef() *ThingGroupReference
}

// The jsii proxy for IThingGroupRef
type jsiiProxy_IThingGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IThingGroupRef) ThingGroupRef() *ThingGroupReference {
	var returns *ThingGroupReference
	_jsii_.Get(
		j,
		"thingGroupRef",
		&returns,
	)
	return returns
}

