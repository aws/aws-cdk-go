package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroupEgress.
// Experimental.
type ISecurityGroupEgressRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecurityGroupEgress resource.
	// Experimental.
	SecurityGroupEgressRef() *SecurityGroupEgressReference
}

// The jsii proxy for ISecurityGroupEgressRef
type jsiiProxy_ISecurityGroupEgressRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISecurityGroupEgressRef) SecurityGroupEgressRef() *SecurityGroupEgressReference {
	var returns *SecurityGroupEgressReference
	_jsii_.Get(
		j,
		"securityGroupEgressRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupEgressRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupEgressRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

