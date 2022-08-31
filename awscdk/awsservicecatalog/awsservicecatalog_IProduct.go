package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicecatalog/internal"
)

// A Service Catalog product, currently only supports type CloudFormationProduct.
// Experimental.
type IProduct interface {
	awscdk.IResource
	// Associate Tag Options.
	//
	// A TagOption is a key-value pair managed in AWS Service Catalog.
	// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
	// Experimental.
	AssociateTagOptions(tagOptions TagOptions)
	// The ARN of the product.
	// Experimental.
	ProductArn() *string
	// The id of the product.
	// Experimental.
	ProductId() *string
}

// The jsii proxy for IProduct
type jsiiProxy_IProduct struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProduct) AssociateTagOptions(tagOptions TagOptions) {
	_jsii_.InvokeVoid(
		i,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

func (j *jsiiProxy_IProduct) ProductArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

