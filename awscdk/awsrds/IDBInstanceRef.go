package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBInstance.
// Experimental.
type IDBInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBInstanceRef
type jsiiProxy_IDBInstanceRef struct {
	internal.Type__constructsIConstruct
}

