package awslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Logs::LogGroup` resource specifies a log group.
//
// A log group defines common properties for log streams, such as their retention and access control rules. Each log stream must belong to one log group.
//
// You can create up to 1,000,000 log groups per Region per account. You must use the following guidelines when naming a log group:
//
// - Log group names must be unique within a Region for an AWS account.
// - Log group names can be between 1 and 512 characters long.
// - Log group names consist of the following characters: a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), '/' (forward slash), and '.' (period).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataProtectionPolicy interface{}
//   var fieldIndexPolicies interface{}
//   var mergeStrategy IMergeStrategy
//   var resourcePolicyDocument interface{}
//
//   cfnLogGroupPropsMixin := awscdkcfnpropertymixins.Aws_logs.NewCfnLogGroupPropsMixin(&CfnLogGroupMixinProps{
//   	DataProtectionPolicy: dataProtectionPolicy,
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	FieldIndexPolicies: []interface{}{
//   		fieldIndexPolicies,
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogGroupClass: jsii.String("logGroupClass"),
//   	LogGroupName: jsii.String("logGroupName"),
//   	ResourcePolicyDocument: resourcePolicyDocument,
//   	RetentionInDays: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
//
type CfnLogGroupPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnLogGroupMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLogGroupPropsMixin
type jsiiProxy_CfnLogGroupPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnLogGroupPropsMixin) Props() *CfnLogGroupMixinProps {
	var returns *CfnLogGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLogGroupPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Logs::LogGroup`.
func NewCfnLogGroupPropsMixin(props *CfnLogGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnLogGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLogGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLogGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Logs::LogGroup`.
func NewCfnLogGroupPropsMixin_Override(c CfnLogGroupPropsMixin, props *CfnLogGroupMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnLogGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLogGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLogGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_logs.CfnLogGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLogGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLogGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

