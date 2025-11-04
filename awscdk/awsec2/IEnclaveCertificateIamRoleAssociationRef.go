package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnclaveCertificateIamRoleAssociation.
// Experimental.
type IEnclaveCertificateIamRoleAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EnclaveCertificateIamRoleAssociation resource.
	// Experimental.
	EnclaveCertificateIamRoleAssociationRef() *EnclaveCertificateIamRoleAssociationReference
}

// The jsii proxy for IEnclaveCertificateIamRoleAssociationRef
type jsiiProxy_IEnclaveCertificateIamRoleAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEnclaveCertificateIamRoleAssociationRef) EnclaveCertificateIamRoleAssociationRef() *EnclaveCertificateIamRoleAssociationReference {
	var returns *EnclaveCertificateIamRoleAssociationReference
	_jsii_.Get(
		j,
		"enclaveCertificateIamRoleAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnclaveCertificateIamRoleAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnclaveCertificateIamRoleAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

