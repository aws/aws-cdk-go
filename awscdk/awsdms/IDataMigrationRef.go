package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataMigration.
// Experimental.
type IDataMigrationRef interface {
	constructs.IConstruct
	// A reference to a DataMigration resource.
	// Experimental.
	DataMigrationRef() *DataMigrationReference
}

// The jsii proxy for IDataMigrationRef
type jsiiProxy_IDataMigrationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDataMigrationRef) DataMigrationRef() *DataMigrationReference {
	var returns *DataMigrationReference
	_jsii_.Get(
		j,
		"dataMigrationRef",
		&returns,
	)
	return returns
}

