package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Route53::RecordSetProtection.
//
// Protects a record set (all record sets with the given name and type) in a hosted zone against changes until the protection is deleted. The protection must DependsOn (or reference with Ref) the AWS::Route53::RecordSet it protects. Adding or removing a protection and changing that record set must be two separate stack updates: a removal is applied in the cleanup phase after updates, and an addition in the same update blocks the record set's rollback. To recover a rollback blocked this way (UPDATE_ROLLBACK_FAILED), delete the protection, then run ContinueUpdateRollback. While a protection exists, creating any record set with its name and type, including one with a new SetIdentifier, is refused. Changing Comment replaces the protection by deleting it first, so the record set is unprotected for the duration of the replacement. While a protection exists, including one kept by DeletionPolicy Retain or created outside the stack, deleting its AWS::Route53::RecordSet fails with ResourceConflict; delete the protection first. When a create is retried after its outcome was lost, an existing protection with the same Comment is taken as the one it created; without a Comment, an existing protection that has no comment is adopted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRecordSetProtectionPropsMixin := awscdkcfnpropertymixins.Aws_route53.NewCfnRecordSetProtectionPropsMixin(&CfnRecordSetProtectionMixinProps{
//   	Comment: jsii.String("comment"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	RecordName: jsii.String("recordName"),
//   	RecordType: jsii.String("recordType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetprotection.html
//
type CfnRecordSetProtectionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRecordSetProtectionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRecordSetProtectionPropsMixin
type jsiiProxy_CfnRecordSetProtectionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRecordSetProtectionPropsMixin) Props() *CfnRecordSetProtectionMixinProps {
	var returns *CfnRecordSetProtectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetProtectionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::RecordSetProtection`.
func NewCfnRecordSetProtectionPropsMixin(props *CfnRecordSetProtectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRecordSetProtectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRecordSetProtectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecordSetProtectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetProtectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::RecordSetProtection`.
func NewCfnRecordSetProtectionPropsMixin_Override(c CfnRecordSetProtectionPropsMixin, props *CfnRecordSetProtectionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetProtectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRecordSetProtectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSetProtectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetProtectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordSetProtectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetProtectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecordSetProtectionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRecordSetProtectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

