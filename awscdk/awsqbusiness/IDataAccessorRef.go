package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataAccessor.
// Experimental.
type IDataAccessorRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataAccessorRef
type jsiiProxy_IDataAccessorRef struct {
	internal.Type__constructsIConstruct
}

