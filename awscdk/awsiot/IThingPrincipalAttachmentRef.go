package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThingPrincipalAttachment.
// Experimental.
type IThingPrincipalAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ThingPrincipalAttachment resource.
	// Experimental.
	ThingPrincipalAttachmentRef() *ThingPrincipalAttachmentReference
}

// The jsii proxy for IThingPrincipalAttachmentRef
type jsiiProxy_IThingPrincipalAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IThingPrincipalAttachmentRef) ThingPrincipalAttachmentRef() *ThingPrincipalAttachmentReference {
	var returns *ThingPrincipalAttachmentReference
	_jsii_.Get(
		j,
		"thingPrincipalAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThingPrincipalAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThingPrincipalAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

