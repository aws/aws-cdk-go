package interfacesawswafv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WebACLAssociation.
// Experimental.
type IWebACLAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WebACLAssociation resource.
	// Experimental.
	WebAclAssociationRef() *WebACLAssociationReference
}

// The jsii proxy for IWebACLAssociationRef
type jsiiProxy_IWebACLAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IWebACLAssociationRef) WebAclAssociationRef() *WebACLAssociationReference {
	var returns *WebACLAssociationReference
	_jsii_.Get(
		j,
		"webAclAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebACLAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebACLAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

