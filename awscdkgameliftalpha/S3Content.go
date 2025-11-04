package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Game content from an S3 archive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   s3Content := gamelift_alpha.NewS3Content(bucket, jsii.String("key"), jsii.String("objectVersion"))
//
// Experimental.
type S3Content interface {
	Content
	// Called when the Build is initialized to allow this object to bind.
	// Experimental.
	Bind(scope constructs.Construct, role awsiam.IRole) *ContentConfig
}

// The jsii proxy struct for S3Content
type jsiiProxy_S3Content struct {
	jsiiProxy_Content
}

// Experimental.
func NewS3Content(bucket awss3.IBucket, key *string, objectVersion *string) S3Content {
	_init_.Initialize()

	if err := validateNewS3ContentParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3Content{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.S3Content",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Content_Override(s S3Content, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.S3Content",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Loads the game content from a local disk path.
// Experimental.
func S3Content_FromAsset(path *string, options *awss3assets.AssetOptions) AssetContent {
	_init_.Initialize()

	if err := validateS3Content_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetContent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.S3Content",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Game content as an S3 object.
// Experimental.
func S3Content_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Content {
	_init_.Initialize()

	if err := validateS3Content_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Content

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.S3Content",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3Content) Bind(scope constructs.Construct, role awsiam.IRole) *ContentConfig {
	if err := s.validateBindParameters(scope, role); err != nil {
		panic(err)
	}
	var returns *ContentConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, role},
		&returns,
	)

	return returns
}

