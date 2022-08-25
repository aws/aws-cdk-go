package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Construct that contains a Service Catalog product stack with its previous deployments maintained.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type s3BucketProduct struct {
//   	productStack
//   }
//
//   func newS3BucketProduct(scope cdk.Construct, id *string) *s3BucketProduct {
//   	this := &s3BucketProduct{}
//   	servicecatalog.NewProductStack_Override(this, scope, id)
//
//   	s3.NewBucket(this, jsii.String("BucketProductV2"))
//   	return this
//   }
//
//   productStackHistory := servicecatalog.NewProductStackHistory(this, jsii.String("ProductStackHistory"), &productStackHistoryProps{
//   	productStack: NewS3BucketProduct(this, jsii.String("S3BucketProduct")),
//   	currentVersionName: jsii.String("v2"),
//   	currentVersionLocked: jsii.Boolean(true),
//   })
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("MyFirstProduct"), &cloudFormationProductProps{
//   	productName: jsii.String("My Product"),
//   	owner: jsii.String("Product Owner"),
//   	productVersions: []cloudFormationProductVersion{
//   		productStackHistory.currentVersion(),
//   	},
//   })
//
type ProductStackHistory interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Retains product stack template as a snapshot when deployed and retrieves a CloudFormationProductVersion for the current product version.
	CurrentVersion() *CloudFormationProductVersion
	// Returns a string representation of this construct.
	ToString() *string
	// Retrieves a CloudFormationProductVersion from a previously deployed productVersionName.
	VersionFromSnapshot(productVersionName *string) *CloudFormationProductVersion
}

// The jsii proxy struct for ProductStackHistory
type jsiiProxy_ProductStackHistory struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ProductStackHistory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewProductStackHistory(scope constructs.Construct, id *string, props *ProductStackHistoryProps) ProductStackHistory {
	_init_.Initialize()

	j := jsiiProxy_ProductStackHistory{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.ProductStackHistory",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewProductStackHistory_Override(p ProductStackHistory, scope constructs.Construct, id *string, props *ProductStackHistoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.ProductStackHistory",
		[]interface{}{scope, id, props},
		p,
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
func ProductStackHistory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.ProductStackHistory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProductStackHistory) CurrentVersion() *CloudFormationProductVersion {
	var returns *CloudFormationProductVersion

	_jsii_.Invoke(
		p,
		"currentVersion",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProductStackHistory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProductStackHistory) VersionFromSnapshot(productVersionName *string) *CloudFormationProductVersion {
	var returns *CloudFormationProductVersion

	_jsii_.Invoke(
		p,
		"versionFromSnapshot",
		[]interface{}{productVersionName},
		&returns,
	)

	return returns
}

