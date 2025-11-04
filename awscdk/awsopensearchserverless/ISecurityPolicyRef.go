package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityPolicy.
// Experimental.
type ISecurityPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecurityPolicy resource.
	// Experimental.
	SecurityPolicyRef() *SecurityPolicyReference
}

// The jsii proxy for ISecurityPolicyRef
type jsiiProxy_ISecurityPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISecurityPolicyRef) SecurityPolicyRef() *SecurityPolicyReference {
	var returns *SecurityPolicyReference
	_jsii_.Get(
		j,
		"securityPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

