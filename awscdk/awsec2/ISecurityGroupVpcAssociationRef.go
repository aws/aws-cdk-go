package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroupVpcAssociation.
// Experimental.
type ISecurityGroupVpcAssociationRef interface {
	constructs.IConstruct
	// A reference to a SecurityGroupVpcAssociation resource.
	// Experimental.
	SecurityGroupVpcAssociationRef() *SecurityGroupVpcAssociationReference
}

// The jsii proxy for ISecurityGroupVpcAssociationRef
type jsiiProxy_ISecurityGroupVpcAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecurityGroupVpcAssociationRef) SecurityGroupVpcAssociationRef() *SecurityGroupVpcAssociationReference {
	var returns *SecurityGroupVpcAssociationReference
	_jsii_.Get(
		j,
		"securityGroupVpcAssociationRef",
		&returns,
	)
	return returns
}

