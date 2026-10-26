package interfacesawsmediaconvert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAssociation.
// Experimental.
type ICertificateAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CertificateAssociation resource.
	// Experimental.
	CertificateAssociationRef() *CertificateAssociationReference
}

// The jsii proxy for ICertificateAssociationRef
type jsiiProxy_ICertificateAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICertificateAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICertificateAssociationRef) CertificateAssociationRef() *CertificateAssociationReference {
	var returns *CertificateAssociationReference
	_jsii_.Get(
		j,
		"certificateAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICertificateAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

