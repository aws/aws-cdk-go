package awsinvoicing

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinvoicing/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InvoiceUnit.
// Experimental.
type IInvoiceUnitRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InvoiceUnit resource.
	// Experimental.
	InvoiceUnitRef() *InvoiceUnitReference
}

// The jsii proxy for IInvoiceUnitRef
type jsiiProxy_IInvoiceUnitRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInvoiceUnitRef) InvoiceUnitRef() *InvoiceUnitReference {
	var returns *InvoiceUnitReference
	_jsii_.Get(
		j,
		"invoiceUnitRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInvoiceUnitRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInvoiceUnitRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

