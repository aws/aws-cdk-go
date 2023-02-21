package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
)

// A parameter group.
//
// Represents both a cluster parameter group,
// and an instance parameter group.
type IParameterGroup interface {
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

