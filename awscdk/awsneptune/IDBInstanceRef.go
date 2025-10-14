package awsneptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBInstance.
// Experimental.
type IDBInstanceRef interface {
	constructs.IConstruct
	// A reference to a DBInstance resource.
	// Experimental.
	DbInstanceRef() *DBInstanceReference
}

// The jsii proxy for IDBInstanceRef
type jsiiProxy_IDBInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBInstanceRef) DbInstanceRef() *DBInstanceReference {
	var returns *DBInstanceReference
	_jsii_.Get(
		j,
		"dbInstanceRef",
		&returns,
	)
	return returns
}

