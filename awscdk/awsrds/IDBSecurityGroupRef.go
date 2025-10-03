package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBSecurityGroup.
// Experimental.
type IDBSecurityGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBSecurityGroupRef
type jsiiProxy_IDBSecurityGroupRef struct {
	internal.Type__constructsIConstruct
}

