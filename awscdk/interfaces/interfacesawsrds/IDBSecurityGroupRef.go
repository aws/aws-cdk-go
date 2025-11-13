package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSecurityGroup.
// Experimental.
type IDBSecurityGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DBSecurityGroup resource.
	// Experimental.
	DbSecurityGroupRef() *DBSecurityGroupReference
}

// The jsii proxy for IDBSecurityGroupRef
type jsiiProxy_IDBSecurityGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDBSecurityGroupRef) DbSecurityGroupRef() *DBSecurityGroupReference {
	var returns *DBSecurityGroupReference
	_jsii_.Get(
		j,
		"dbSecurityGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSecurityGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDBSecurityGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

