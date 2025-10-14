package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Thing.
// Experimental.
type IThingRef interface {
	constructs.IConstruct
	// A reference to a Thing resource.
	// Experimental.
	ThingRef() *ThingReference
}

// The jsii proxy for IThingRef
type jsiiProxy_IThingRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IThingRef) ThingRef() *ThingReference {
	var returns *ThingReference
	_jsii_.Get(
		j,
		"thingRef",
		&returns,
	)
	return returns
}

