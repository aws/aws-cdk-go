package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroupVpcAssociation.
// Experimental.
type ISecurityGroupVpcAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecurityGroupVpcAssociation resource.
	// Experimental.
	SecurityGroupVpcAssociationRef() *SecurityGroupVpcAssociationReference
}

// The jsii proxy for ISecurityGroupVpcAssociationRef
type jsiiProxy_ISecurityGroupVpcAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISecurityGroupVpcAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupVpcAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

