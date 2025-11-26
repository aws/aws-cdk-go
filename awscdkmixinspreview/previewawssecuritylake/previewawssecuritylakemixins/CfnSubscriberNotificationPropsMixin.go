package previewawssecuritylakemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssecuritylake/previewawssecuritylakemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Notifies the subscriber when new data is written to the data lake for the sources that the subscriber consumes in Security Lake.
//
// You can create only one subscriber notification per subscriber.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var sqsNotificationConfiguration interface{}
//
//   cfnSubscriberNotificationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSubscriberNotificationPropsMixin(&CfnSubscriberNotificationMixinProps{
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		HttpsNotificationConfiguration: &HttpsNotificationConfigurationProperty{
//   			AuthorizationApiKeyName: jsii.String("authorizationApiKeyName"),
//   			AuthorizationApiKeyValue: jsii.String("authorizationApiKeyValue"),
//   			Endpoint: jsii.String("endpoint"),
//   			HttpMethod: jsii.String("httpMethod"),
//   			TargetRoleArn: jsii.String("targetRoleArn"),
//   		},
//   		SqsNotificationConfiguration: sqsNotificationConfiguration,
//   	},
//   	SubscriberArn: jsii.String("subscriberArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securitylake-subscribernotification.html
//
type CfnSubscriberNotificationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSubscriberNotificationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSubscriberNotificationPropsMixin
type jsiiProxy_CfnSubscriberNotificationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSubscriberNotificationPropsMixin) Props() *CfnSubscriberNotificationMixinProps {
	var returns *CfnSubscriberNotificationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSubscriberNotificationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SecurityLake::SubscriberNotification`.
func NewCfnSubscriberNotificationPropsMixin(props *CfnSubscriberNotificationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSubscriberNotificationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSubscriberNotificationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSubscriberNotificationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SecurityLake::SubscriberNotification`.
func NewCfnSubscriberNotificationPropsMixin_Override(c CfnSubscriberNotificationPropsMixin, props *CfnSubscriberNotificationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSubscriberNotificationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSubscriberNotificationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSubscriberNotificationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_securitylake.mixins.CfnSubscriberNotificationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSubscriberNotificationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSubscriberNotificationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

