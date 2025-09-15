package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnclaveCertificateIamRoleAssociation.
// Experimental.
type IEnclaveCertificateIamRoleAssociationRef interface {
	constructs.IConstruct
	// A reference to a EnclaveCertificateIamRoleAssociation resource.
	// Experimental.
	EnclaveCertificateIamRoleAssociationRef() *EnclaveCertificateIamRoleAssociationReference
}

// The jsii proxy for IEnclaveCertificateIamRoleAssociationRef
type jsiiProxy_IEnclaveCertificateIamRoleAssociationRef struct {
	internal.Type__constructsIConstruct
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

