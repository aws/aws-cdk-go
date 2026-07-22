package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReservedDBInstance.
// Experimental.
type IReservedDBInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReservedDBInstance resource.
	// Experimental.
	ReservedDbInstanceRef() *ReservedDBInstanceReference
}

// The jsii proxy for IReservedDBInstanceRef
type jsiiProxy_IReservedDBInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReservedDBInstanceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IReservedDBInstanceRef) ReservedDbInstanceRef() *ReservedDBInstanceReference {
	var returns *ReservedDBInstanceReference
	_jsii_.Get(
		j,
		"reservedDbInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReservedDBInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReservedDBInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

