package interfacesawslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLakeSettings.
// Experimental.
type IDataLakeSettingsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataLakeSettings resource.
	// Experimental.
	DataLakeSettingsRef() *DataLakeSettingsReference
}

// The jsii proxy for IDataLakeSettingsRef
type jsiiProxy_IDataLakeSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataLakeSettingsRef) DataLakeSettingsRef() *DataLakeSettingsReference {
	var returns *DataLakeSettingsReference
	_jsii_.Get(
		j,
		"dataLakeSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataLakeSettingsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataLakeSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

