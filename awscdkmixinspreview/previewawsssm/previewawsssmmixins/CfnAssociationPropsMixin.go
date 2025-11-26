package previewawsssmmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsssm/previewawsssmmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SSM::Association` resource creates a State Manager association for your managed instances.
//
// A State Manager association defines the state that you want to maintain on your instances. For example, an association can specify that anti-virus software must be installed and running on your instances, or that certain ports must be closed. For static targets, the association specifies a schedule for when the configuration is reapplied. For dynamic targets, such as an AWS Resource Groups or an AWS Auto Scaling Group, State Manager applies the configuration when new instances are added to the group. The association also specifies actions to take when applying the configuration. For example, an association for anti-virus software might run once a day. If the software is not installed, then State Manager installs it. If the software is installed, but the service is not running, then the association might instruct State Manager to start the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnAssociationPropsMixin(&CfnAssociationMixinProps{
//   	ApplyOnlyAtCronInterval: jsii.Boolean(false),
//   	AssociationName: jsii.String("associationName"),
//   	AutomationTargetParameterName: jsii.String("automationTargetParameterName"),
//   	CalendarNames: []*string{
//   		jsii.String("calendarNames"),
//   	},
//   	ComplianceSeverity: jsii.String("complianceSeverity"),
//   	DocumentVersion: jsii.String("documentVersion"),
//   	InstanceId: jsii.String("instanceId"),
//   	MaxConcurrency: jsii.String("maxConcurrency"),
//   	MaxErrors: jsii.String("maxErrors"),
//   	Name: jsii.String("name"),
//   	OutputLocation: &InstanceAssociationOutputLocationProperty{
//   		S3Location: &S3OutputLocationProperty{
//   			OutputS3BucketName: jsii.String("outputS3BucketName"),
//   			OutputS3KeyPrefix: jsii.String("outputS3KeyPrefix"),
//   			OutputS3Region: jsii.String("outputS3Region"),
//   		},
//   	},
//   	Parameters: parameters,
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleOffset: jsii.Number(123),
//   	SyncCompliance: jsii.String("syncCompliance"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	WaitForSuccessTimeoutSeconds: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-association.html
//
type CfnAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssociationPropsMixin
type jsiiProxy_CfnAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAssociationPropsMixin) Props() *CfnAssociationMixinProps {
	var returns *CfnAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SSM::Association`.
func NewCfnAssociationPropsMixin(props *CfnAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SSM::Association`.
func NewCfnAssociationPropsMixin_Override(c CfnAssociationPropsMixin, props *CfnAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ssm.mixins.CfnAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

