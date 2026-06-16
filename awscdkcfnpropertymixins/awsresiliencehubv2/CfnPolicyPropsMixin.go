package awsresiliencehubv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsresiliencehubv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a resilience policy that defines availability and disaster recovery requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPolicyPropsMixin := awscdkcfnpropertymixins.Aws_resiliencehubv2.NewCfnPolicyPropsMixin(&CfnPolicyMixinProps{
//   	AvailabilitySlo: &AvailabilitySloProperty{
//   		Target: jsii.Number(123),
//   	},
//   	DataRecovery: &DataRecoveryTargetsProperty{
//   		TimeBetweenBackupsInMinutes: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MultiAz: &MultiAzTargetsProperty{
//   		DisasterRecoveryApproach: jsii.String("disasterRecoveryApproach"),
//   		RpoInMinutes: jsii.Number(123),
//   		RtoInMinutes: jsii.Number(123),
//   	},
//   	MultiRegion: &MultiRegionTargetsProperty{
//   		DisasterRecoveryApproach: jsii.String("disasterRecoveryApproach"),
//   		RpoInMinutes: jsii.Number(123),
//   		RtoInMinutes: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-policy.html
//
type CfnPolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPolicyPropsMixin
type jsiiProxy_CfnPolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPolicyPropsMixin) Props() *CfnPolicyMixinProps {
	var returns *CfnPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ResilienceHubV2::Policy`.
func NewCfnPolicyPropsMixin(props *CfnPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ResilienceHubV2::Policy`.
func NewCfnPolicyPropsMixin_Override(c CfnPolicyPropsMixin, props *CfnPolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_resiliencehubv2.CfnPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

