package awsecrassets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An asset that represents a Docker image.
//
// The image will loaded from an existing tarball and uploaded to an ECR repository.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewTarballImageAsset(this, jsii.String("MyBuildImage"), &TarballImageAssetProps{
//   	TarballFile: jsii.String("local-image.tar"),
//   })
//
type TarballImageAsset interface {
	constructs.Construct
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	AssetHash() *string
	// The tag of this asset when it is uploaded to ECR.
	//
	// The tag may differ from the assetHash if a stack synthesizer adds a dockerTagPrefix.
	ImageTag() *string
	// The full URI of the image (including a tag).
	//
	// Use this reference to pull
	// the asset.
	ImageUri() *string
	SetImageUri(val *string)
	// The tree node.
	Node() constructs.Node
	// Repository where the image is stored.
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TarballImageAsset
type jsiiProxy_TarballImageAsset struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TarballImageAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) ImageTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}


func NewTarballImageAsset(scope constructs.Construct, id *string, props *TarballImageAssetProps) TarballImageAsset {
	_init_.Initialize()

	if err := validateNewTarballImageAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TarballImageAsset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTarballImageAsset_Override(t TarballImageAsset, scope constructs.Construct, id *string, props *TarballImageAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TarballImageAsset)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_TarballImageAsset)SetRepository(val awsecr.IRepository) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func TarballImageAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTarballImageAsset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TarballImageAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

