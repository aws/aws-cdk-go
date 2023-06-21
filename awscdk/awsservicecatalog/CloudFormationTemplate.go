package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the Product Provisioning Artifact Template.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type s3BucketProduct struct {
//   	productStack
//   }
//
//   func newS3BucketProduct(scope construct, id *string) *s3BucketProduct {
//   	this := &s3BucketProduct{}
//   	servicecatalog.NewProductStack_Override(this, scope, id)
//
//   	s3.NewBucket(this, jsii.String("BucketProduct"))
//   	return this
//   }
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("Product"), &CloudFormationProductProps{
//   	ProductName: jsii.String("My Product"),
//   	Owner: jsii.String("Product Owner"),
//   	ProductVersions: []cloudFormationProductVersion{
//   		&cloudFormationProductVersion{
//   			ProductVersionName: jsii.String("v1"),
//   			CloudFormationTemplate: servicecatalog.CloudFormationTemplate_FromProductStack(NewS3BucketProduct(this, jsii.String("S3BucketProduct"))),
//   		},
//   	},
//   })
//
type CloudFormationTemplate interface {
	// Called when the product is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *CloudFormationTemplateConfig
}

// The jsii proxy struct for CloudFormationTemplate
type jsiiProxy_CloudFormationTemplate struct {
	_ byte // padding
}

func NewCloudFormationTemplate_Override(c CloudFormationTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.CloudFormationTemplate",
		nil, // no parameters
		c,
	)
}

// Loads the provisioning artifacts template from a local disk path.
func CloudFormationTemplate_FromAsset(path *string, options *awss3assets.AssetOptions) CloudFormationTemplate {
	_init_.Initialize()

	if err := validateCloudFormationTemplate_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CloudFormationTemplate",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Creates a product with the resources defined in the given product stack.
func CloudFormationTemplate_FromProductStack(productStack ProductStack) CloudFormationTemplate {
	_init_.Initialize()

	if err := validateCloudFormationTemplate_FromProductStackParameters(productStack); err != nil {
		panic(err)
	}
	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CloudFormationTemplate",
		"fromProductStack",
		[]interface{}{productStack},
		&returns,
	)

	return returns
}

// Template from URL.
func CloudFormationTemplate_FromUrl(url *string) CloudFormationTemplate {
	_init_.Initialize()

	if err := validateCloudFormationTemplate_FromUrlParameters(url); err != nil {
		panic(err)
	}
	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.CloudFormationTemplate",
		"fromUrl",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationTemplate) Bind(scope constructs.Construct) *CloudFormationTemplateConfig {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CloudFormationTemplateConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

