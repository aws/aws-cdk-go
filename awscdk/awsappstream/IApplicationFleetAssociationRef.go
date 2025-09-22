package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationFleetAssociation.
// Experimental.
type IApplicationFleetAssociationRef interface {
	constructs.IConstruct
	// A reference to a ApplicationFleetAssociation resource.
	// Experimental.
	ApplicationFleetAssociationRef() *ApplicationFleetAssociationReference
}

// The jsii proxy for IApplicationFleetAssociationRef
type jsiiProxy_IApplicationFleetAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationFleetAssociationRef) ApplicationFleetAssociationRef() *ApplicationFleetAssociationReference {
	var returns *ApplicationFleetAssociationReference
	_jsii_.Get(
		j,
		"applicationFleetAssociationRef",
		&returns,
	)
	return returns
}

