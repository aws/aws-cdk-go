package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Tool Schema from a local asset.
//
// The asset is uploaded to an S3 staging bucket, then moved to its final location
// by CloudFormation during deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage DockerImage
//   var grantable IGrantable
//   var keyRef IKeyRef
//   var localBundling ILocalBundling
//
//   assetToolSchema := bedrock_agentcore_alpha.NewAssetToolSchema(jsii.String("path"), &AssetOptions{
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
//   		Volumes: []DockerVolume{
//   			&DockerVolume{
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
//   	DisplayName: jsii.String("displayName"),
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	FollowSymlinks: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   	Readers: []IGrantable{
//   		grantable,
//   	},
//   	SourceKMSKey: keyRef,
//   })
//
// Experimental.
type AssetToolSchema interface {
	ToolSchema
	// The account ID of the S3 bucket owner for cross-account access.
	// Experimental.
	BucketOwnerAccountId() *string
	// The inline tool schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	// Experimental.
	InlineSchema() *[]*ToolDefinition
	// The S3 location of the tool schema file, if using an S3-based schema.
	//
	// Contains the bucket name and object key information.
	// Experimental.
	S3File() *awss3.Location
	// Binds this tool schema to a construct scope.
	//
	// This method initializes the S3 asset if it hasn't been initialized yet.
	// Must be called before rendering the schema as CFN properties.
	// Experimental.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	// Experimental.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for AssetToolSchema
type jsiiProxy_AssetToolSchema struct {
	jsiiProxy_ToolSchema
}

func (j *jsiiProxy_AssetToolSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetToolSchema) InlineSchema() *[]*ToolDefinition {
	var returns *[]*ToolDefinition
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetToolSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetToolSchema(path *string, options *awss3assets.AssetOptions) AssetToolSchema {
	_init_.Initialize()

	if err := validateNewAssetToolSchemaParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetToolSchema{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetToolSchema",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetToolSchema_Override(a AssetToolSchema, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetToolSchema",
		[]interface{}{path, options},
		a,
	)
}

// Creates a Tool Schema from an inline string.
// Experimental.
func AssetToolSchema_FromInline(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateAssetToolSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetToolSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates a tool Schema from a local file.
// Experimental.
func AssetToolSchema_FromLocalAsset(path *string) ToolSchema {
	_init_.Initialize()

	if err := validateAssetToolSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns ToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetToolSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates a Tool Schema from an S3 File.
// Experimental.
func AssetToolSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ToolSchema {
	_init_.Initialize()

	if err := validateAssetToolSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AssetToolSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetToolSchema) Bind(scope constructs.Construct) {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope},
	)
}

func (a *jsiiProxy_AssetToolSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := a.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}

