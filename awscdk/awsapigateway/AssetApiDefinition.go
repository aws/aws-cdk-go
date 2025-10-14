package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// OpenAPI specification from a local file.
//
// Example:
//   myApiDefinition := apigateway.ApiDefinition_FromAsset(jsii.String("path-to-file.json"))
//   specRestApi := apigateway.NewSpecRestApi(this, jsii.String("my-specrest-api"), &SpecRestApiProps{
//   	Deploy: jsii.Boolean(false),
//   	ApiDefinition: myApiDefinition,
//   })
//
//   // Use `stageName` to deploy to an existing stage
//   deployment := apigateway.NewDeployment(this, jsii.String("my-deployment"), &DeploymentProps{
//   	Api: specRestApi,
//   	StageName: jsii.String("dev"),
//   	RetainDeployments: jsii.Boolean(true),
//   })
//
//   // Trigger a new deployment on OpenAPI definition updates
//   deployment.AddToLogicalId(myApiDefinition)
//
type AssetApiDefinition interface {
	ApiDefinition
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(scope constructs.Construct, restApi IRestApiRef)
}

// The jsii proxy struct for AssetApiDefinition
type jsiiProxy_AssetApiDefinition struct {
	jsiiProxy_ApiDefinition
}

func NewAssetApiDefinition(path *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	if err := validateNewAssetApiDefinitionParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetApiDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetApiDefinition_Override(a AssetApiDefinition, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		[]interface{}{path, options},
		a,
	)
}

// Loads the API specification from a local disk asset.
func AssetApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	if err := validateAssetApiDefinition_FromAssetParameters(file, options); err != nil {
		panic(err)
	}
	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func AssetApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	if err := validateAssetApiDefinition_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
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
//     apigateway.ApiDefinition.fromInline({
//       openapi: '3.0.2',
//       paths: {
//         '/pets': {
//           get: {
//             'responses': {
//               200: {
//                 content: {
//                   'application/json': {
//                     schema: {
//                       $ref: '#/components/schemas/Empty',
//                     },
//                   },
//                 },
//               },
//             },
//             'x-amazon-apigateway-integration': {
//               responses: {
//                 default: {
//                   statusCode: '200',
//                 },
//               },
//               requestTemplates: {
//                 'application/json': '{"statusCode": 200}',
//               },
//               passthroughBehavior: 'when_no_match',
//               type: 'mock',
//             },
//           },
//         },
//       },
//       components: {
//         schemas: {
//           Empty: {
//             title: 'Empty Schema',
//             type: 'object',
//           },
//         },
//       },
//     });
//
func AssetApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	if err := validateAssetApiDefinition_FromInlineParameters(definition); err != nil {
		panic(err)
	}
	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.AssetApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetApiDefinition) Bind(scope constructs.Construct) *ApiDefinitionConfig {
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

func (a *jsiiProxy_AssetApiDefinition) BindAfterCreate(scope constructs.Construct, restApi IRestApiRef) {
	if err := a.validateBindAfterCreateParameters(scope, restApi); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindAfterCreate",
		[]interface{}{scope, restApi},
	)
}

