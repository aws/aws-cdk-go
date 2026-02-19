package awsdocdb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdocdb/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdocdb"
	"github.com/aws/constructs-go/constructs/v10"
)

// A parameter group.
type IClusterParameterGroup interface {
	interfacesawsdocdb.IDBClusterParameterGroupRef
	awscdk.IResource
	// The name of this parameter group.
	ParameterGroupName() *string
}

// The jsii proxy for IClusterParameterGroup
type jsiiProxy_IClusterParameterGroup struct {
	internal.Type__interfacesawsdocdbIDBClusterParameterGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IClusterParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IClusterParameterGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IClusterParameterGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterParameterGroup) DbClusterParameterGroupRef() *interfacesawsdocdb.DBClusterParameterGroupReference {
	var returns *interfacesawsdocdb.DBClusterParameterGroupReference
	_jsii_.Get(
		j,
		"dbClusterParameterGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterParameterGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClusterParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

