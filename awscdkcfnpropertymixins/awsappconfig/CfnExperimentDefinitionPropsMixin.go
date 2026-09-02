package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::AppConfig::ExperimentDefinition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnExperimentDefinitionPropsMixin := awscdkcfnpropertymixins.Aws_appconfig.NewCfnExperimentDefinitionPropsMixin(&CfnExperimentDefinitionMixinProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	AudienceDescription: jsii.String("audienceDescription"),
//   	AudienceRule: jsii.String("audienceRule"),
//   	ConfigurationProfileIdentifier: jsii.String("configurationProfileIdentifier"),
//   	Control: &TreatmentProperty{
//   		AttributeValues: map[string]interface{}{
//   			"attributeValuesKey": &AttributeValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"numberArray": []interface{}{
//   					jsii.Number(123),
//   				},
//   				"numberValue": jsii.Number(123),
//   				"stringArray": []*string{
//   					jsii.String("stringArray"),
//   				},
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		Enabled: jsii.Boolean(false),
//   		Key: jsii.String("key"),
//   		Weight: jsii.Number(123),
//   	},
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	FlagKey: jsii.String("flagKey"),
//   	Hypothesis: jsii.String("hypothesis"),
//   	LaunchCriteria: jsii.String("launchCriteria"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Treatments: []interface{}{
//   		&TreatmentProperty{
//   			AttributeValues: map[string]interface{}{
//   				"attributeValuesKey": &AttributeValueProperty{
//   					"booleanValue": jsii.Boolean(false),
//   					"numberArray": []interface{}{
//   						jsii.Number(123),
//   					},
//   					"numberValue": jsii.Number(123),
//   					"stringArray": []*string{
//   						jsii.String("stringArray"),
//   					},
//   					"stringValue": jsii.String("stringValue"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Enabled: jsii.Boolean(false),
//   			Key: jsii.String("key"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-experimentdefinition.html
//
type CfnExperimentDefinitionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnExperimentDefinitionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExperimentDefinitionPropsMixin
type jsiiProxy_CfnExperimentDefinitionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnExperimentDefinitionPropsMixin) Props() *CfnExperimentDefinitionMixinProps {
	var returns *CfnExperimentDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExperimentDefinitionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppConfig::ExperimentDefinition`.
func NewCfnExperimentDefinitionPropsMixin(props *CfnExperimentDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnExperimentDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExperimentDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExperimentDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExperimentDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppConfig::ExperimentDefinition`.
func NewCfnExperimentDefinitionPropsMixin_Override(c CfnExperimentDefinitionPropsMixin, props *CfnExperimentDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExperimentDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnExperimentDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExperimentDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExperimentDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExperimentDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_appconfig.CfnExperimentDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExperimentDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnExperimentDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

