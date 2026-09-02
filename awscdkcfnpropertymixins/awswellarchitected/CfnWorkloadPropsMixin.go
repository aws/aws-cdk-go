package awswellarchitected

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::WellArchitected::Workload Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkloadPropsMixin := awscdkcfnpropertymixins.Aws_wellarchitected.NewCfnWorkloadPropsMixin(&CfnWorkloadMixinProps{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	ArchitecturalDesign: jsii.String("architecturalDesign"),
//   	AwsRegions: []*string{
//   		jsii.String("awsRegions"),
//   	},
//   	Description: jsii.String("description"),
//   	DiscoveryConfig: &DiscoveryConfigProperty{
//   		TrustedAdvisorIntegrationStatus: jsii.String("trustedAdvisorIntegrationStatus"),
//   		WorkloadResourceDefinition: []*string{
//   			jsii.String("workloadResourceDefinition"),
//   		},
//   	},
//   	Environment: jsii.String("environment"),
//   	Industry: jsii.String("industry"),
//   	IndustryType: jsii.String("industryType"),
//   	Lenses: []*string{
//   		jsii.String("lenses"),
//   	},
//   	NonAwsRegions: []*string{
//   		jsii.String("nonAwsRegions"),
//   	},
//   	Notes: jsii.String("notes"),
//   	ReviewOwner: jsii.String("reviewOwner"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkloadName: jsii.String("workloadName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wellarchitected-workload.html
//
type CfnWorkloadPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkloadMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkloadPropsMixin
type jsiiProxy_CfnWorkloadPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWorkloadPropsMixin) Props() *CfnWorkloadMixinProps {
	var returns *CfnWorkloadMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkloadPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WellArchitected::Workload`.
func NewCfnWorkloadPropsMixin(props *CfnWorkloadMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkloadPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkloadPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkloadPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnWorkloadPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WellArchitected::Workload`.
func NewCfnWorkloadPropsMixin_Override(c CfnWorkloadPropsMixin, props *CfnWorkloadMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnWorkloadPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkloadPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkloadPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnWorkloadPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkloadPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_wellarchitected.CfnWorkloadPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkloadPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkloadPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

