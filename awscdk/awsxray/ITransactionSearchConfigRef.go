package awsxray

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsxray/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransactionSearchConfig.
// Experimental.
type ITransactionSearchConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITransactionSearchConfigRef
type jsiiProxy_ITransactionSearchConfigRef struct {
	internal.Type__constructsIConstruct
}

