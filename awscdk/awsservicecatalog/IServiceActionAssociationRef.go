package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceActionAssociation.
// Experimental.
type IServiceActionAssociationRef interface {
	constructs.IConstruct
	// A reference to a ServiceActionAssociation resource.
	// Experimental.
	ServiceActionAssociationRef() *ServiceActionAssociationReference
}

// The jsii proxy for IServiceActionAssociationRef
type jsiiProxy_IServiceActionAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceActionAssociationRef) ServiceActionAssociationRef() *ServiceActionAssociationReference {
	var returns *ServiceActionAssociationReference
	_jsii_.Get(
		j,
		"serviceActionAssociationRef",
		&returns,
	)
	return returns
}

