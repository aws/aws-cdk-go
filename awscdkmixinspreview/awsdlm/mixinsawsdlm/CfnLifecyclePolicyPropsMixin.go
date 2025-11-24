package mixinsawsdlm

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsdlm/mixinsawsdlm/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a lifecycle policy, which is used to automate operations on Amazon EBS resources.
//
// The properties are required when you add a lifecycle policy and optional when you update a lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var crossRegionCopyTargets interface{}
//   var excludeTags interface{}
//   var excludeVolumeTypes interface{}
//
//   cfnLifecyclePolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnLifecyclePolicyPropsMixin(&CfnLifecyclePolicyMixinProps{
//   	CopyTags: jsii.Boolean(false),
//   	CreateInterval: jsii.Number(123),
//   	CrossRegionCopyTargets: crossRegionCopyTargets,
//   	DefaultPolicy: jsii.String("defaultPolicy"),
//   	Description: jsii.String("description"),
//   	Exclusions: &ExclusionsProperty{
//   		ExcludeBootVolumes: jsii.Boolean(false),
//   		ExcludeTags: excludeTags,
//   		ExcludeVolumeTypes: excludeVolumeTypes,
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	ExtendDeletion: jsii.Boolean(false),
//   	PolicyDetails: &PolicyDetailsProperty{
//   		Actions: []interface{}{
//   			&ActionProperty{
//   				CrossRegionCopy: []interface{}{
//   					&CrossRegionCopyActionProperty{
//   						EncryptionConfiguration: &EncryptionConfigurationProperty{
//   							CmkArn: jsii.String("cmkArn"),
//   							Encrypted: jsii.Boolean(false),
//   						},
//   						RetainRule: &CrossRegionCopyRetainRuleProperty{
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   						Target: jsii.String("target"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		CopyTags: jsii.Boolean(false),
//   		CreateInterval: jsii.Number(123),
//   		CrossRegionCopyTargets: crossRegionCopyTargets,
//   		EventSource: &EventSourceProperty{
//   			Parameters: &EventParametersProperty{
//   				DescriptionRegex: jsii.String("descriptionRegex"),
//   				EventType: jsii.String("eventType"),
//   				SnapshotOwner: []*string{
//   					jsii.String("snapshotOwner"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   		Exclusions: &ExclusionsProperty{
//   			ExcludeBootVolumes: jsii.Boolean(false),
//   			ExcludeTags: excludeTags,
//   			ExcludeVolumeTypes: excludeVolumeTypes,
//   		},
//   		ExtendDeletion: jsii.Boolean(false),
//   		Parameters: &ParametersProperty{
//   			ExcludeBootVolume: jsii.Boolean(false),
//   			ExcludeDataVolumeTags: []interface{}{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			NoReboot: jsii.Boolean(false),
//   		},
//   		PolicyLanguage: jsii.String("policyLanguage"),
//   		PolicyType: jsii.String("policyType"),
//   		ResourceLocations: []*string{
//   			jsii.String("resourceLocations"),
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   		RetainInterval: jsii.Number(123),
//   		Schedules: []interface{}{
//   			&ScheduleProperty{
//   				ArchiveRule: &ArchiveRuleProperty{
//   					RetainRule: &ArchiveRetainRuleProperty{
//   						RetentionArchiveTier: &RetentionArchiveTierProperty{
//   							Count: jsii.Number(123),
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   					},
//   				},
//   				CopyTags: jsii.Boolean(false),
//   				CreateRule: &CreateRuleProperty{
//   					CronExpression: jsii.String("cronExpression"),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   					Location: jsii.String("location"),
//   					Scripts: []interface{}{
//   						&ScriptProperty{
//   							ExecuteOperationOnScriptFailure: jsii.Boolean(false),
//   							ExecutionHandler: jsii.String("executionHandler"),
//   							ExecutionHandlerService: jsii.String("executionHandlerService"),
//   							ExecutionTimeout: jsii.Number(123),
//   							MaximumRetryCount: jsii.Number(123),
//   							Stages: []*string{
//   								jsii.String("stages"),
//   							},
//   						},
//   					},
//   					Times: []*string{
//   						jsii.String("times"),
//   					},
//   				},
//   				CrossRegionCopyRules: []interface{}{
//   					&CrossRegionCopyRuleProperty{
//   						CmkArn: jsii.String("cmkArn"),
//   						CopyTags: jsii.Boolean(false),
//   						DeprecateRule: &CrossRegionCopyDeprecateRuleProperty{
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   						Encrypted: jsii.Boolean(false),
//   						RetainRule: &CrossRegionCopyRetainRuleProperty{
//   							Interval: jsii.Number(123),
//   							IntervalUnit: jsii.String("intervalUnit"),
//   						},
//   						Target: jsii.String("target"),
//   						TargetRegion: jsii.String("targetRegion"),
//   					},
//   				},
//   				DeprecateRule: &DeprecateRuleProperty{
//   					Count: jsii.Number(123),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   				},
//   				FastRestoreRule: &FastRestoreRuleProperty{
//   					AvailabilityZones: []*string{
//   						jsii.String("availabilityZones"),
//   					},
//   					Count: jsii.Number(123),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   				},
//   				Name: jsii.String("name"),
//   				RetainRule: &RetainRuleProperty{
//   					Count: jsii.Number(123),
//   					Interval: jsii.Number(123),
//   					IntervalUnit: jsii.String("intervalUnit"),
//   				},
//   				ShareRules: []interface{}{
//   					&ShareRuleProperty{
//   						TargetAccounts: []*string{
//   							jsii.String("targetAccounts"),
//   						},
//   						UnshareInterval: jsii.Number(123),
//   						UnshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   					},
//   				},
//   				TagsToAdd: []interface{}{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				VariableTags: []interface{}{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		TargetTags: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RetainInterval: jsii.Number(123),
//   	State: jsii.String("state"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html
//
type CfnLifecyclePolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLifecyclePolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLifecyclePolicyPropsMixin
type jsiiProxy_CfnLifecyclePolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLifecyclePolicyPropsMixin) Props() *CfnLifecyclePolicyMixinProps {
	var returns *CfnLifecyclePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecyclePolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicyPropsMixin(props *CfnLifecyclePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLifecyclePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLifecyclePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLifecyclePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DLM::LifecyclePolicy`.
func NewCfnLifecyclePolicyPropsMixin_Override(c CfnLifecyclePolicyPropsMixin, props *CfnLifecyclePolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLifecyclePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLifecyclePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecyclePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dlm.mixins.CfnLifecyclePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLifecyclePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLifecyclePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

