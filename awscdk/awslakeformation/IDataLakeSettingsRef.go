package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataLakeSettings.
// Experimental.
type IDataLakeSettingsRef interface {
	constructs.IConstruct
	// A reference to a DataLakeSettings resource.
	// Experimental.
	DataLakeSettingsRef() *DataLakeSettingsReference
}

// The jsii proxy for IDataLakeSettingsRef
type jsiiProxy_IDataLakeSettingsRef struct {
	internal.Type__constructsIConstruct
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

