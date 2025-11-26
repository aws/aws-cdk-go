package previewawsobservabilityadminmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsobservabilityadmin/previewawsobservabilityadminmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines how telemetry data should be centralized across an AWS Organization, including source and destination configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationCentralizationRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnOrganizationCentralizationRulePropsMixin(&CfnOrganizationCentralizationRuleMixinProps{
//   	Rule: &CentralizationRuleProperty{
//   		Destination: &CentralizationRuleDestinationProperty{
//   			Account: jsii.String("account"),
//   			DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   				BackupConfiguration: &LogsBackupConfigurationProperty{
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					Region: jsii.String("region"),
//   				},
//   				LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   					EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   					EncryptionStrategy: jsii.String("encryptionStrategy"),
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   			},
//   			Region: jsii.String("region"),
//   		},
//   		Source: &CentralizationRuleSourceProperty{
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   			Scope: jsii.String("scope"),
//   			SourceLogsConfiguration: &SourceLogsConfigurationProperty{
//   				EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   				LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   			},
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html
//
type CfnOrganizationCentralizationRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOrganizationCentralizationRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationCentralizationRulePropsMixin
type jsiiProxy_CfnOrganizationCentralizationRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOrganizationCentralizationRulePropsMixin) Props() *CfnOrganizationCentralizationRuleMixinProps {
	var returns *CfnOrganizationCentralizationRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationCentralizationRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ObservabilityAdmin::OrganizationCentralizationRule`.
func NewCfnOrganizationCentralizationRulePropsMixin(props *CfnOrganizationCentralizationRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOrganizationCentralizationRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationCentralizationRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationCentralizationRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ObservabilityAdmin::OrganizationCentralizationRule`.
func NewCfnOrganizationCentralizationRulePropsMixin_Override(c CfnOrganizationCentralizationRulePropsMixin, props *CfnOrganizationCentralizationRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOrganizationCentralizationRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationCentralizationRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationCentralizationRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationCentralizationRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOrganizationCentralizationRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

