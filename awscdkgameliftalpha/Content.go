package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Before deploying your GameLift-enabled multiplayer game servers for hosting with the GameLift service, you need to upload your game server files.
//
// The class helps you on preparing and uploading custom game server build files or Realtime Servers server script files.
//
// Example:
//   var bucket bucket
//
//   build := gamelift.NewBuild(this, jsii.String("Build"), &BuildProps{
//   	Content: gamelift.Content_FromBucket(bucket, jsii.String("sample-asset-key")),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("BuildArn"), &CfnOutputProps{
//   	Value: build.BuildArn,
//   })
//   awscdk.NewCfnOutput(this, jsii.String("BuildId"), &CfnOutputProps{
//   	Value: build.BuildId,
//   })
//
// Experimental.
type Content interface {
	// Called when the Build is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, role awsiam.IRole) *ContentConfig
}

// The jsii proxy struct for Content
type jsiiProxy_Content struct {
	_ byte // padding
}

// Experimental.
func NewContent_Override(c Content) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Content",
		nil, // no parameters
		c,
	)
}

// Loads the game content from a local disk path.
// Experimental.
func Content_FromAsset(path *string, options *awss3assets.AssetOptions) AssetContent {
	_init_.Initialize()

	if err := validateContent_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Content",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Game content as an S3 object.
// Experimental.
func Content_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Content {
	_init_.Initialize()

	if err := validateContent_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Content

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Content",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Content) Bind(scope constructs.Construct, role awsiam.IRole) *ContentConfig {
	if err := c.validateBindParameters(scope, role); err != nil {
		panic(err)
	}
	var returns *ContentConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, role},
		&returns,
	)

	return returns
}

