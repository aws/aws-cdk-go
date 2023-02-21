package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// A receipt rule.
type IReceiptRule interface {
	awscdk.IResource
	// The name of the receipt rule.
	ReceiptRuleName() *string
}

// The jsii proxy for IReceiptRule
type jsiiProxy_IReceiptRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IReceiptRule) ReceiptRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receiptRuleName",
		&returns,
	)
	return returns
}

