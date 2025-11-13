package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecurityGroup.
// Experimental.
type ISecurityGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SecurityGroup resource.
	// Experimental.
	SecurityGroupRef() *SecurityGroupReference
}

// The jsii proxy for ISecurityGroupRef
type jsiiProxy_ISecurityGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISecurityGroupRef) SecurityGroupRef() *SecurityGroupReference {
	var returns *SecurityGroupReference
	_jsii_.Get(
		j,
		"securityGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecurityGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

