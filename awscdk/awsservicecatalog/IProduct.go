package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog product, currently only supports type CloudFormationProduct.
type IProduct interface {
	interfacesawsservicecatalog.ICloudFormationProductRef
	awscdk.IResource
	// Associate Tag Options.
	//
	// A TagOption is a key-value pair managed in AWS Service Catalog.
	// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
	AssociateTagOptions(tagOptions TagOptions)
	// The asset buckets of a product created via product stack.
	AssetBuckets() *[]awss3.IBucket
	// The ARN of the product.
	ProductArn() *string
	// The id of the product.
	ProductId() *string
}

// The jsii proxy for IProduct
type jsiiProxy_IProduct struct {
	internal.Type__interfacesawsservicecatalogICloudFormationProductRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProduct) AssociateTagOptions(tagOptions TagOptions) {
	if err := i.validateAssociateTagOptionsParameters(tagOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

func (i *jsiiProxy_IProduct) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IProduct) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IProduct) AssetBuckets() *[]awss3.IBucket {
	var returns *[]awss3.IBucket
	_jsii_.Get(
		j,
		"assetBuckets",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_IProduct) CloudFormationProductRef() *interfacesawsservicecatalog.CloudFormationProductReference {
	var returns *interfacesawsservicecatalog.CloudFormationProductReference
	_jsii_.Get(
		j,
		"cloudFormationProductRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProduct) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProduct) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

