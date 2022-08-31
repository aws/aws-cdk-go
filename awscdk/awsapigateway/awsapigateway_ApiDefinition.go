package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Represents an OpenAPI definition asset.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &specRestApiProps{
//   	apiDefinition: apigateway.apiDefinition.fromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.root.addResource(jsii.String("books"))
//   booksResource.addMethod(jsii.String("GET"), integration)
//
// Experimental.
type ApiDefinition interface {
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(scope awscdk.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	// Experimental.
	BindAfterCreate(_scope awscdk.Construct, _restApi IRestApi)
}

// The jsii proxy struct for ApiDefinition
type jsiiProxy_ApiDefinition struct {
	_ byte // padding
}

// Experimental.
func NewApiDefinition_Override(a ApiDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.ApiDefinition",
		nil, // no parameters
		a,
	)
}

// Loads the API specification from a local disk asset.
// Experimental.
func ApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.ApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
// Experimental.
func ApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.ApiDefinition",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Create an API definition from an inline object.
//
// The inline object must follow the
// schema of OpenAPI 2.0 or OpenAPI 3.0
//
// Example:
//   apigateway.apiDefinition.fromInline(map[string]interface{}{
//   	"openapi": jsii.String("3.0.2"),
//   	"paths": map[string]map[string]map[string]map[string]map[string]map[string]map[string]map[string]*string{
//   		"/pets": map[string]map[string]map[string]map[string]map[string]map[string]map[string]*string{
//   			"get": map[string]map[string]map[string]map[string]map[string]map[string]*string{
//   				"responses": map[string]map[string]map[string]map[string]map[string]*string{
//   					jsii.Number(200): map[string]map[string]map[string]map[string]*string{
//   						"content": map[string]map[string]map[string]*string{
//   							"application/json": map[string]map[string]*string{
//   								"schema": map[string]*string{
//   									"$ref": jsii.String("#/components/schemas/Empty"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				"x-amazon-apigateway-integration": map[string]interface{}{
//   					"responses": map[string]map[string]*string{
//   						"default": map[string]*string{
//   							"statusCode": jsii.String("200"),
//   						},
//   					},
//   					"requestTemplates": map[string]*string{
//   						"application/json": jsii.String("{\"statusCode\": 200}"),
//   					},
//   					"passthroughBehavior": jsii.String("when_no_match"),
//   					"type": jsii.String("mock"),
//   				},
//   			},
//   		},
//   	},
//   	"components": map[string]map[string]map[string]*string{
//   		"schemas": map[string]map[string]*string{
//   			"Empty": map[string]*string{
//   				"title": jsii.String("Empty Schema"),
//   				"type": jsii.String("object"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
func ApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.ApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDefinition) Bind(scope awscdk.Construct) *ApiDefinitionConfig {
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDefinition) BindAfterCreate(_scope awscdk.Construct, _restApi IRestApi) {
	_jsii_.InvokeVoid(
		a,
		"bindAfterCreate",
		[]interface{}{_scope, _restApi},
	)
}

