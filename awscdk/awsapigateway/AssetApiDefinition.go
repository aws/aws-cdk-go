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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var grantable iGrantable
//   var localBundling iLocalBundling
//
//   assetApiDefinition := awscdk.Aws_apigateway.NewAssetApiDefinition(jsii.String("path"), &AssetOptions{
//   	AssetHash: jsii.String("assetHash"),
//   	AssetHashType: cdk.AssetHashType_SOURCE,
//   	Bundling: &BundlingOptions{
//   		Image: dockerImage,
//
//   		// the properties below are optional
//   		BundlingFileAccess: cdk.BundlingFileAccess_VOLUME_COPY,
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Local: localBundling,
//   		Network: jsii.String("network"),
//   		OutputType: cdk.BundlingOutput_ARCHIVED,
//   		Platform: jsii.String("platform"),
//   		SecurityOpt: jsii.String("securityOpt"),
//   		User: jsii.String("user"),
//   		Volumes: []dockerVolume{
//   			&dockerVolume{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				Consistency: cdk.DockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		VolumesFrom: []*string{
//   			jsii.String("volumesFrom"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	DeployTime: jsii.Boolean(false),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Readers: []*iGrantable{
//   		grantable,
//   	},
//   })
//
type AssetApiDefinition interface {
	ApiDefinition
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(scope constructs.Construct, restApi IRestApi)
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

func (a *jsiiProxy_AssetApiDefinition) BindAfterCreate(scope constructs.Construct, restApi IRestApi) {
	if err := a.validateBindAfterCreateParameters(scope, restApi); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindAfterCreate",
		[]interface{}{scope, restApi},
	)
}

