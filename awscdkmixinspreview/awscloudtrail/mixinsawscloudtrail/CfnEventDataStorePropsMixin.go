package mixinsawscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscloudtrail/mixinsawscloudtrail/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new event data store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventDataStorePropsMixin := awscdkmixinspreview.Mixins.NewCfnEventDataStorePropsMixin(&CfnEventDataStoreMixinProps{
//   	AdvancedEventSelectors: []interface{}{
//   		&AdvancedEventSelectorProperty{
//   			FieldSelectors: []interface{}{
//   				&AdvancedFieldSelectorProperty{
//   					EndsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					EqualTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					Field: jsii.String("field"),
//   					NotEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					NotEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					NotStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					StartsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	BillingMode: jsii.String("billingMode"),
//   	ContextKeySelectors: []interface{}{
//   		&ContextKeySelectorProperty{
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	FederationEnabled: jsii.Boolean(false),
//   	FederationRoleArn: jsii.String("federationRoleArn"),
//   	IngestionEnabled: jsii.Boolean(false),
//   	InsightsDestination: jsii.String("insightsDestination"),
//   	InsightSelectors: []interface{}{
//   		&InsightSelectorProperty{
//   			InsightType: jsii.String("insightType"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MaxEventSize: jsii.String("maxEventSize"),
//   	MultiRegionEnabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	OrganizationEnabled: jsii.Boolean(false),
//   	RetentionPeriod: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationProtectionEnabled: jsii.Boolean(false),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html
//
type CfnEventDataStorePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEventDataStoreMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventDataStorePropsMixin
type jsiiProxy_CfnEventDataStorePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEventDataStorePropsMixin) Props() *CfnEventDataStoreMixinProps {
	var returns *CfnEventDataStoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventDataStorePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudTrail::EventDataStore`.
func NewCfnEventDataStorePropsMixin(props *CfnEventDataStoreMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEventDataStorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventDataStorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventDataStorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnEventDataStorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudTrail::EventDataStore`.
func NewCfnEventDataStorePropsMixin_Override(c CfnEventDataStorePropsMixin, props *CfnEventDataStoreMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnEventDataStorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEventDataStorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventDataStorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnEventDataStorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventDataStorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudtrail.mixins.CfnEventDataStorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventDataStorePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEventDataStorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

