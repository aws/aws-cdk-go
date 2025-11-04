package awsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Protection.
// Experimental.
type IProtectionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Protection resource.
	// Experimental.
	ProtectionRef() *ProtectionReference
}

// The jsii proxy for IProtectionRef
type jsiiProxy_IProtectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IProtectionRef) ProtectionRef() *ProtectionReference {
	var returns *ProtectionReference
	_jsii_.Get(
		j,
		"protectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProtectionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProtectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

