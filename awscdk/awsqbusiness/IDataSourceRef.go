package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataSource.
// Experimental.
type IDataSourceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DataSource resource.
	// Experimental.
	DataSourceRef() *DataSourceReference
}

// The jsii proxy for IDataSourceRef
type jsiiProxy_IDataSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDataSourceRef) DataSourceRef() *DataSourceReference {
	var returns *DataSourceReference
	_jsii_.Get(
		j,
		"dataSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataSourceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

