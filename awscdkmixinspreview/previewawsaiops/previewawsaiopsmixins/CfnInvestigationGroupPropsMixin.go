package previewawsaiopsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsaiops/previewawsaiopsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an *investigation group* in your account.
//
// Creating an investigation group is a one-time setup task for each Region in your account. It is a necessary task to be able to perform investigations.
//
// Settings in the investigation group help you centrally manage the common properties of your investigations, such as the following:
//
// - Who can access the investigations
// - Whether investigation data is encrypted with a customer managed AWS Key Management Service key.
// - How long investigations and their data are retained by default.
//
// Currently, you can have one investigation group in each Region in your account. Each investigation in a Region is a part of the investigation group in that Region
//
// To create an investigation group and set up CloudWatch investigations, you must be signed in to an IAM principal that has either the `AIOpsConsoleAdminPolicy` or the `AdministratorAccess` IAM policy attached, or to an account that has similar permissions.
//
// > You can configure CloudWatch alarms to start investigations and add events to investigations. If you create your investigation group with `CreateInvestigationGroup` and you want to enable alarms to do this, you must use `PutInvestigationGroupPolicy` to create a resource policy that grants this permission to CloudWatch alarms.
// >
// > For more information about configuring CloudWatch alarms, see [Using Amazon CloudWatch alarms](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInvestigationGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnInvestigationGroupPropsMixin(&CfnInvestigationGroupMixinProps{
//   	ChatbotNotificationChannels: []interface{}{
//   		&ChatbotNotificationChannelProperty{
//   			ChatConfigurationArns: []*string{
//   				jsii.String("chatConfigurationArns"),
//   			},
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	CrossAccountConfigurations: []interface{}{
//   		&CrossAccountConfigurationProperty{
//   			SourceRoleArn: jsii.String("sourceRoleArn"),
//   		},
//   	},
//   	EncryptionConfig: &EncryptionConfigMapProperty{
//   		EncryptionConfigurationType: jsii.String("encryptionConfigurationType"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	InvestigationGroupPolicy: jsii.String("investigationGroupPolicy"),
//   	IsCloudTrailEventHistoryEnabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	RetentionInDays: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   	TagKeyBoundaries: []*string{
//   		jsii.String("tagKeyBoundaries"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html
//
type CfnInvestigationGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInvestigationGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInvestigationGroupPropsMixin
type jsiiProxy_CfnInvestigationGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInvestigationGroupPropsMixin) Props() *CfnInvestigationGroupMixinProps {
	var returns *CfnInvestigationGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInvestigationGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AIOps::InvestigationGroup`.
func NewCfnInvestigationGroupPropsMixin(props *CfnInvestigationGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInvestigationGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInvestigationGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInvestigationGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AIOps::InvestigationGroup`.
func NewCfnInvestigationGroupPropsMixin_Override(c CfnInvestigationGroupPropsMixin, props *CfnInvestigationGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInvestigationGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInvestigationGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInvestigationGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_aiops.mixins.CfnInvestigationGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInvestigationGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnInvestigationGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

