package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Represents the Product Provisioning Artifact Template.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("Product"), &cloudFormationProductProps{
//   	productName: jsii.String("My Product"),
//   	owner: jsii.String("Product Owner"),
//   	productVersions: []cloudFormationProductVersion{
//   		&cloudFormationProductVersion{
//   			productVersionName: jsii.String("v1"),
//   			cloudFormationTemplate: servicecatalog.cloudFormationTemplate.fromUrl(jsii.String("https://raw.githubusercontent.com/awslabs/aws-cloudformation-templates/master/aws/services/ServiceCatalog/Product.yaml")),
//   		},
//   		&cloudFormationProductVersion{
//   			productVersionName: jsii.String("v2"),
//   			cloudFormationTemplate: servicecatalog.*cloudFormationTemplate.fromAsset(path.join(__dirname, jsii.String("development-environment.template.json"))),
//   		},
//   	},
//   })
//
// Experimental.
type CloudFormationTemplate interface {
	// Called when the product is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(scope awscdk.Construct) *CloudFormationTemplateConfig
}

// The jsii proxy struct for CloudFormationTemplate
type jsiiProxy_CloudFormationTemplate struct {
	_ byte // padding
}

// Experimental.
func NewCloudFormationTemplate_Override(c CloudFormationTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_servicecatalog.CloudFormationTemplate",
		nil, // no parameters
		c,
	)
}

// Loads the provisioning artifacts template from a local disk path.
// Experimental.
func CloudFormationTemplate_FromAsset(path *string, options *awss3assets.AssetOptions) CloudFormationTemplate {
	_init_.Initialize()

	if err := validateCloudFormationTemplate_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_servicecatalog.CloudFormationTemplate",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Creates a product with the resources defined in the given product stack.
// Experimental.
func CloudFormationTemplate_FromProductStack(productStack ProductStack) CloudFormationTemplate {
	_init_.Initialize()

	if err := validateCloudFormationTemplate_FromProductStackParameters(productStack); err != nil {
		panic(err)
	}
	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_servicecatalog.CloudFormationTemplate",
		"fromProductStack",
		[]interface{}{productStack},
		&returns,
	)

	return returns
}

// Template from URL.
// Experimental.
func CloudFormationTemplate_FromUrl(url *string) CloudFormationTemplate {
	_init_.Initialize()

	if err := validateCloudFormationTemplate_FromUrlParameters(url); err != nil {
		panic(err)
	}
	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_servicecatalog.CloudFormationTemplate",
		"fromUrl",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationTemplate) Bind(scope awscdk.Construct) *CloudFormationTemplateConfig {
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

