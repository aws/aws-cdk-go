package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityKey.
// Experimental.
type ISecurityKeyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecurityKey resource.
	// Experimental.
	SecurityKeyRef() *SecurityKeyReference
}

// The jsii proxy for ISecurityKeyRef
type jsiiProxy_ISecurityKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISecurityKeyRef) SecurityKeyRef() *SecurityKeyReference {
	var returns *SecurityKeyReference
	_jsii_.Get(
		j,
		"securityKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityKeyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

