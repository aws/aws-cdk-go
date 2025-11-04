package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroupIngress.
// Experimental.
type ISecurityGroupIngressRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecurityGroupIngress resource.
	// Experimental.
	SecurityGroupIngressRef() *SecurityGroupIngressReference
}

// The jsii proxy for ISecurityGroupIngressRef
type jsiiProxy_ISecurityGroupIngressRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISecurityGroupIngressRef) SecurityGroupIngressRef() *SecurityGroupIngressReference {
	var returns *SecurityGroupIngressReference
	_jsii_.Get(
		j,
		"securityGroupIngressRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupIngressRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupIngressRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

