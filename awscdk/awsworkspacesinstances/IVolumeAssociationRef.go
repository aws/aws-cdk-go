package awsworkspacesinstances

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesinstances/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VolumeAssociation.
// Experimental.
type IVolumeAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVolumeAssociationRef
type jsiiProxy_IVolumeAssociationRef struct {
	internal.Type__constructsIConstruct
}

