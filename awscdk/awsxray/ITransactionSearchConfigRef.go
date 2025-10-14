package awsxray

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsxray/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransactionSearchConfig.
// Experimental.
type ITransactionSearchConfigRef interface {
	constructs.IConstruct
	// A reference to a TransactionSearchConfig resource.
	// Experimental.
	TransactionSearchConfigRef() *TransactionSearchConfigReference
}

// The jsii proxy for ITransactionSearchConfigRef
type jsiiProxy_ITransactionSearchConfigRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransactionSearchConfigRef) TransactionSearchConfigRef() *TransactionSearchConfigReference {
	var returns *TransactionSearchConfigReference
	_jsii_.Get(
		j,
		"transactionSearchConfigRef",
		&returns,
	)
	return returns
}

