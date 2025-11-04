package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataSet.
// Experimental.
type IDataSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DataSet resource.
	// Experimental.
	DataSetRef() *DataSetReference
}

// The jsii proxy for IDataSetRef
type jsiiProxy_IDataSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDataSetRef) DataSetRef() *DataSetReference {
	var returns *DataSetReference
	_jsii_.Get(
		j,
		"dataSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

