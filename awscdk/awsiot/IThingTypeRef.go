package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingType.
// Experimental.
type IThingTypeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ThingType resource.
	// Experimental.
	ThingTypeRef() *ThingTypeReference
}

// The jsii proxy for IThingTypeRef
type jsiiProxy_IThingTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IThingTypeRef) ThingTypeRef() *ThingTypeReference {
	var returns *ThingTypeReference
	_jsii_.Get(
		j,
		"thingTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThingTypeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThingTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

