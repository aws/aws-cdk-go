package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an OpenAPI definition asset.
//
// Example:
//   var integration integration
//
//
//   api := apigateway.NewSpecRestApi(this, jsii.String("books-api"), &SpecRestApiProps{
//   	ApiDefinition: apigateway.ApiDefinition_FromAsset(jsii.String("path-to-file.json")),
//   })
//
//   booksResource := api.Root.AddResource(jsii.String("books"))
//   booksResource.AddMethod(jsii.String("GET"), integration)
//
type ApiDefinition interface {
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(_scope constructs.Construct, _restApi IRestApi)
}

// The jsii proxy struct for ApiDefinition
type jsiiProxy_ApiDefinition struct {
	_ byte // padding
}

func NewApiDefinition_Override(a ApiDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		nil, // no parameters
		a,
	)
}

// Loads the API specification from a local disk asset.
func ApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	if err := validateApiDefinition_FromAssetParameters(file, options); err != nil {
		panic(err)
	}
	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func ApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	if err := validateApiDefinition_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
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
//   apigateway.ApiDefinition_FromInline(map[string]interface{}{
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
func ApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	if err := validateApiDefinition_FromInlineParameters(definition); err != nil {
		panic(err)
	}
	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.ApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDefinition) Bind(scope constructs.Construct) *ApiDefinitionConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiDefinition) BindAfterCreate(_scope constructs.Construct, _restApi IRestApi) {
	if err := a.validateBindAfterCreateParameters(_scope, _restApi); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindAfterCreate",
		[]interface{}{_scope, _restApi},
	)
}

