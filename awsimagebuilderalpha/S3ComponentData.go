package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for S3-based component data references, containing additional permission grant methods on the S3 object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage DockerImage
//   var grantable IGrantable
//   var keyRef IKeyRef
//   var localBundling ILocalBundling
//
//   s3ComponentData := imagebuilder_alpha.S3ComponentData_FromAsset(this, jsii.String("MyS3ComponentData"), jsii.String("path"), &AssetOptions{
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
type S3ComponentData interface {
	ComponentData
	// Experimental.
	Bucket() awss3.IBucket
	// Experimental.
	Key() *string
	// Grant put permissions to the given grantee for the component data in S3.
	// Experimental.
	GrantPut(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read permissions to the given grantee for the component data in S3.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The rendered component data text, for use in CloudFormation.
	// Experimental.
	Render() *ComponentDataConfig
}

// The jsii proxy struct for S3ComponentData
type jsiiProxy_S3ComponentData struct {
	jsiiProxy_ComponentData
}

func (j *jsiiProxy_S3ComponentData) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ComponentData) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3ComponentData_Override(s S3ComponentData, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		[]interface{}{bucket, key},
		s,
	)
}

// Uploads component data from a local file to S3 to use as the component data.
// Experimental.
func S3ComponentData_FromAsset(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) S3ComponentData {
	_init_.Initialize()

	if err := validateS3ComponentData_FromAssetParameters(scope, id, path, options); err != nil {
		panic(err)
	}
	var returns S3ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		"fromAsset",
		[]interface{}{scope, id, path, options},
		&returns,
	)

	return returns
}

// Uses an inline JSON object as the component data, using the ComponentDocument interface.
// Experimental.
func S3ComponentData_FromComponentDocumentJsonObject(data *ComponentDocument) ComponentData {
	_init_.Initialize()

	if err := validateS3ComponentData_FromComponentDocumentJsonObjectParameters(data); err != nil {
		panic(err)
	}
	var returns ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		"fromComponentDocumentJsonObject",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// Uses an inline JSON/YAML string as the component data.
// Experimental.
func S3ComponentData_FromInline(data *string) ComponentData {
	_init_.Initialize()

	if err := validateS3ComponentData_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// Uses an inline JSON object as the component data.
// Experimental.
func S3ComponentData_FromJsonObject(data *map[string]interface{}) ComponentData {
	_init_.Initialize()

	if err := validateS3ComponentData_FromJsonObjectParameters(data); err != nil {
		panic(err)
	}
	var returns ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		"fromJsonObject",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// References component data from a pre-existing S3 object.
// Experimental.
func S3ComponentData_FromS3(bucket awss3.IBucket, key *string) S3ComponentData {
	_init_.Initialize()

	if err := validateS3ComponentData_FromS3Parameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ComponentData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ComponentData) GrantPut(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantPutParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantPut",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ComponentData) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ComponentData) Render() *ComponentDataConfig {
	var returns *ComponentDataConfig

	_jsii_.Invoke(
		s,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

