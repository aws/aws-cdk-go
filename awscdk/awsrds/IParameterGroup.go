package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds"
	"github.com/aws/constructs-go/constructs/v10"
)

// A parameter group.
//
// Represents both a cluster parameter group,
// and an instance parameter group.
type IParameterGroup interface {
	interfacesawsrds.IDBClusterParameterGroupRef
	interfacesawsrds.IDBParameterGroupRef
	awscdk.IResource
	// Adds a parameter to this group.
	//
	// If this is an imported parameter group,
	// this method does nothing.
	//
	// Returns: true if the parameter was actually added
	// (i.e., this ParameterGroup is not imported),
	// false otherwise.
	AddParameter(key *string, value *string) *bool
	// Method called when this Parameter Group is used when defining a database cluster.
	BindToCluster(options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig
	// Method called when this Parameter Group is used when defining a database instance.
	BindToInstance(options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig
}

// The jsii proxy for IParameterGroup
type jsiiProxy_IParameterGroup struct {
	internal.Type__interfacesawsrdsIDBClusterParameterGroupRef
	internal.Type__interfacesawsrdsIDBParameterGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IParameterGroup) AddParameter(key *string, value *string) *bool {
	if err := i.validateAddParameterParameters(key, value); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"addParameter",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameterGroup) BindToCluster(options *ParameterGroupClusterBindOptions) *ParameterGroupClusterConfig {
	if err := i.validateBindToClusterParameters(options); err != nil {
		panic(err)
	}
	var returns *ParameterGroupClusterConfig

	_jsii_.Invoke(
		i,
		"bindToCluster",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameterGroup) BindToInstance(options *ParameterGroupInstanceBindOptions) *ParameterGroupInstanceConfig {
	if err := i.validateBindToInstanceParameters(options); err != nil {
		panic(err)
	}
	var returns *ParameterGroupInstanceConfig

	_jsii_.Invoke(
		i,
		"bindToInstance",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IParameterGroup) DbClusterParameterGroupRef() *interfacesawsrds.DBClusterParameterGroupReference {
	var returns *interfacesawsrds.DBClusterParameterGroupReference
	_jsii_.Get(
		j,
		"dbClusterParameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterGroup) DbParameterGroupRef() *interfacesawsrds.DBParameterGroupReference {
	var returns *interfacesawsrds.DBParameterGroupReference
	_jsii_.Get(
		j,
		"dbParameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

