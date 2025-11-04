package awsworkspacesinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesinstances/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VolumeAssociation.
// Experimental.
type IVolumeAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VolumeAssociation resource.
	// Experimental.
	VolumeAssociationRef() *VolumeAssociationReference
}

// The jsii proxy for IVolumeAssociationRef
type jsiiProxy_IVolumeAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IVolumeAssociationRef) VolumeAssociationRef() *VolumeAssociationReference {
	var returns *VolumeAssociationReference
	_jsii_.Get(
		j,
		"volumeAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolumeAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVolumeAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

