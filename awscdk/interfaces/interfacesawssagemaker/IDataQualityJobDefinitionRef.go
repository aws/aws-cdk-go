package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataQualityJobDefinition.
// Experimental.
type IDataQualityJobDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataQualityJobDefinition resource.
	// Experimental.
	DataQualityJobDefinitionRef() *DataQualityJobDefinitionReference
}

// The jsii proxy for IDataQualityJobDefinitionRef
type jsiiProxy_IDataQualityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataQualityJobDefinitionRef) DataQualityJobDefinitionRef() *DataQualityJobDefinitionReference {
	var returns *DataQualityJobDefinitionReference
	_jsii_.Get(
		j,
		"dataQualityJobDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataQualityJobDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataQualityJobDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

