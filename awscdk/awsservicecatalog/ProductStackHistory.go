package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Construct that contains a Service Catalog product stack with its previous deployments maintained.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type s3BucketProduct struct {
//   	productStack
//   }
//
//   func newS3BucketProduct(scope construct, id *string) *s3BucketProduct {
//   	this := &s3BucketProduct{}
//   	servicecatalog.NewProductStack_Override(this, scope, id)
//
//   	s3.NewBucket(this, jsii.String("BucketProductV2"))
//   	return this
//   }
//
//   productStackHistory := servicecatalog.NewProductStackHistory(this, jsii.String("ProductStackHistory"), &ProductStackHistoryProps{
//   	ProductStack: NewS3BucketProduct(this, jsii.String("S3BucketProduct")),
//   	CurrentVersionName: jsii.String("v2"),
//   	CurrentVersionLocked: jsii.Boolean(true),
//   })
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("MyFirstProduct"), &CloudFormationProductProps{
//   	ProductName: jsii.String("My Product"),
//   	Owner: jsii.String("Product Owner"),
//   	ProductVersions: []cloudFormationProductVersion{
//   		productStackHistory.CurrentVersion(),
//   	},
//   })
//
// Experimental.
type ProductStackHistory interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Retains product stack template as a snapshot when deployed and retrieves a CloudFormationProductVersion for the current product version.
	// Experimental.
	CurrentVersion() *CloudFormationProductVersion
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Retrieves a CloudFormationProductVersion from a previously deployed productVersionName.
	// Experimental.
	VersionFromSnapshot(productVersionName *string) *CloudFormationProductVersion
}

// The jsii proxy struct for ProductStackHistory
type jsiiProxy_ProductStackHistory struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_ProductStackHistory) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewProductStackHistory(scope constructs.Construct, id *string, props *ProductStackHistoryProps) ProductStackHistory {
	_init_.Initialize()

	if err := validateNewProductStackHistoryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProductStackHistory{}

	_jsii_.Create(
		"monocdk.aws_servicecatalog.ProductStackHistory",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewProductStackHistory_Override(p ProductStackHistory, scope constructs.Construct, id *string, props *ProductStackHistoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_servicecatalog.ProductStackHistory",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ProductStackHistory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProductStackHistory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_servicecatalog.ProductStackHistory",
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

func (p *jsiiProxy_ProductStackHistory) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProductStackHistory) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_ProductStackHistory) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProductStackHistory) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProductStackHistory) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
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

func (p *jsiiProxy_ProductStackHistory) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProductStackHistory) VersionFromSnapshot(productVersionName *string) *CloudFormationProductVersion {
	if err := p.validateVersionFromSnapshotParameters(productVersionName); err != nil {
		panic(err)
	}
	var returns *CloudFormationProductVersion

	_jsii_.Invoke(
		p,
		"versionFromSnapshot",
		[]interface{}{productVersionName},
		&returns,
	)

	return returns
}

