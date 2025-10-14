package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSecurityGroup.
// Experimental.
type IDBSecurityGroupRef interface {
	constructs.IConstruct
	// A reference to a DBSecurityGroup resource.
	// Experimental.
	DbSecurityGroupRef() *DBSecurityGroupReference
}

// The jsii proxy for IDBSecurityGroupRef
type jsiiProxy_IDBSecurityGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDBSecurityGroupRef) DbSecurityGroupRef() *DBSecurityGroupReference {
	var returns *DBSecurityGroupReference
	_jsii_.Get(
		j,
		"dbSecurityGroupRef",
		&returns,
	)
	return returns
}

