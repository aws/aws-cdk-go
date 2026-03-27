package awscleanroomsml

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscleanroomsml/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnConfiguredModelAlgorithmAssociationPropsMixin := awscdkcfnpropertymixins.Aws_cleanroomsml.NewCfnConfiguredModelAlgorithmAssociationPropsMixin(&CfnConfiguredModelAlgorithmAssociationMixinProps{
//   	ConfiguredModelAlgorithmArn: jsii.String("configuredModelAlgorithmArn"),
//   	Description: jsii.String("description"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	PrivacyConfiguration: &PrivacyConfigurationProperty{
//   		Policies: &PrivacyConfigurationPoliciesProperty{
//   			TrainedModelExports: &TrainedModelExportsConfigurationPolicyProperty{
//   				FilesToExport: []*string{
//   					jsii.String("filesToExport"),
//   				},
//   				MaxSize: &TrainedModelExportsMaxSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			TrainedModelInferenceJobs: &TrainedModelInferenceJobsConfigurationPolicyProperty{
//   				ContainerLogs: []interface{}{
//   					&LogsConfigurationPolicyProperty{
//   						AllowedAccountIds: []*string{
//   							jsii.String("allowedAccountIds"),
//   						},
//   						FilterPattern: jsii.String("filterPattern"),
//   						LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   							CustomEntityConfig: &CustomEntityConfigProperty{
//   								CustomDataIdentifiers: []*string{
//   									jsii.String("customDataIdentifiers"),
//   								},
//   							},
//   							EntitiesToRedact: []*string{
//   								jsii.String("entitiesToRedact"),
//   							},
//   						},
//   						LogType: jsii.String("logType"),
//   					},
//   				},
//   				MaxOutputSize: &TrainedModelInferenceMaxOutputSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   			TrainedModels: &TrainedModelsConfigurationPolicyProperty{
//   				ContainerLogs: []interface{}{
//   					&LogsConfigurationPolicyProperty{
//   						AllowedAccountIds: []*string{
//   							jsii.String("allowedAccountIds"),
//   						},
//   						FilterPattern: jsii.String("filterPattern"),
//   						LogRedactionConfiguration: &LogRedactionConfigurationProperty{
//   							CustomEntityConfig: &CustomEntityConfigProperty{
//   								CustomDataIdentifiers: []*string{
//   									jsii.String("customDataIdentifiers"),
//   								},
//   							},
//   							EntitiesToRedact: []*string{
//   								jsii.String("entitiesToRedact"),
//   							},
//   						},
//   						LogType: jsii.String("logType"),
//   					},
//   				},
//   				ContainerMetrics: &MetricsConfigurationPolicyProperty{
//   					NoiseLevel: jsii.String("noiseLevel"),
//   				},
//   				MaxArtifactSize: &TrainedModelArtifactMaxSizeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-configuredmodelalgorithmassociation.html
//
type CfnConfiguredModelAlgorithmAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnConfiguredModelAlgorithmAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConfiguredModelAlgorithmAssociationPropsMixin
type jsiiProxy_CfnConfiguredModelAlgorithmAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnConfiguredModelAlgorithmAssociationPropsMixin) Props() *CfnConfiguredModelAlgorithmAssociationMixinProps {
	var returns *CfnConfiguredModelAlgorithmAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConfiguredModelAlgorithmAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation`.
func NewCfnConfiguredModelAlgorithmAssociationPropsMixin(props *CfnConfiguredModelAlgorithmAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnConfiguredModelAlgorithmAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConfiguredModelAlgorithmAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConfiguredModelAlgorithmAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation`.
func NewCfnConfiguredModelAlgorithmAssociationPropsMixin_Override(c CfnConfiguredModelAlgorithmAssociationPropsMixin, props *CfnConfiguredModelAlgorithmAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnConfiguredModelAlgorithmAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConfiguredModelAlgorithmAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConfiguredModelAlgorithmAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnConfiguredModelAlgorithmAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConfiguredModelAlgorithmAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnConfiguredModelAlgorithmAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

