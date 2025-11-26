package previewawsroute53recoveryreadinessmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53recoveryreadiness/previewawsroute53recoveryreadinessmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a cell in recovery group in Amazon Route 53 Application Recovery Controller.
//
// A cell in Route 53 ARC represents replicas or independent units of failover in your application. It groups within it all the AWS resources that are necessary for your application to run independently. Typically, you would have define one set of resources in a primary cell and another set in a standby cell in your recovery group.
//
// After you set up the cells for your application, you can create readiness checks in Route 53 ARC to continually audit readiness for AWS resource quotas, capacity, network routing policies, and other predefined rules.
//
// You can set up notifications about changes that would affect your ability to fail over to a replica and recover. However, you should make decisions about whether to fail away from or to a replica based on your monitoring and health check systems. You should consider readiness checks as a complementary service to those systems.
//
// Route 53 ARC Readiness supports us-east-1 and us-west-2 AWS Regions only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCellPropsMixin := awscdkmixinspreview.Mixins.NewCfnCellPropsMixin(&CfnCellMixinProps{
//   	CellName: jsii.String("cellName"),
//   	Cells: []*string{
//   		jsii.String("cells"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoveryreadiness-cell.html
//
type CfnCellPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCellMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCellPropsMixin
type jsiiProxy_CfnCellPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCellPropsMixin) Props() *CfnCellMixinProps {
	var returns *CfnCellMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCellPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53RecoveryReadiness::Cell`.
func NewCfnCellPropsMixin(props *CfnCellMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCellPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCellPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCellPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnCellPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53RecoveryReadiness::Cell`.
func NewCfnCellPropsMixin_Override(c CfnCellPropsMixin, props *CfnCellMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnCellPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCellPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCellPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnCellPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCellPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.mixins.CfnCellPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCellPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCellPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

