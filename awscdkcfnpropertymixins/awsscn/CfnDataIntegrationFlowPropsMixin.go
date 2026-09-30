package awsscn

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsscn/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an AWS Supply Chain data integration flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDataIntegrationFlowPropsMixin := awscdkcfnpropertymixins.Aws_scn.NewCfnDataIntegrationFlowPropsMixin(&CfnDataIntegrationFlowMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	Name: jsii.String("name"),
//   	Sources: []interface{}{
//   		&DataIntegrationFlowSourceProperty{
//   			DatasetSource: &DatasetSourceConfigurationProperty{
//   				DatasetIdentifier: jsii.String("datasetIdentifier"),
//   				Options: &DatasetOptionsProperty{
//   					DedupeRecords: jsii.Boolean(false),
//   					DedupeStrategy: &DedupeStrategyProperty{
//   						FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   							Fields: []interface{}{
//   								&FieldPriorityDedupeFieldProperty{
//   									Name: jsii.String("name"),
//   									SortOrder: jsii.String("sortOrder"),
//   								},
//   							},
//   						},
//   						Type: jsii.String("type"),
//   					},
//   					LoadType: jsii.String("loadType"),
//   				},
//   			},
//   			S3Source: &S3SourceConfigurationProperty{
//   				BucketName: jsii.String("bucketName"),
//   				Options: &S3OptionsProperty{
//   					FileType: jsii.String("fileType"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   			},
//   			SourceName: jsii.String("sourceName"),
//   			SourceType: jsii.String("sourceType"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Target: &DataIntegrationFlowTargetProperty{
//   		DatasetTarget: &DatasetTargetConfigurationProperty{
//   			DatasetIdentifier: jsii.String("datasetIdentifier"),
//   			Options: &DatasetOptionsProperty{
//   				DedupeRecords: jsii.Boolean(false),
//   				DedupeStrategy: &DedupeStrategyProperty{
//   					FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   						Fields: []interface{}{
//   							&FieldPriorityDedupeFieldProperty{
//   								Name: jsii.String("name"),
//   								SortOrder: jsii.String("sortOrder"),
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   				LoadType: jsii.String("loadType"),
//   			},
//   		},
//   		TargetType: jsii.String("targetType"),
//   	},
//   	Transformation: &DataIntegrationFlowTransformationProperty{
//   		SqlTransformation: &SqlTransformationConfigurationProperty{
//   			Query: jsii.String("query"),
//   		},
//   		TransformationType: jsii.String("transformationType"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-scn-dataintegrationflow.html
//
type CfnDataIntegrationFlowPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDataIntegrationFlowMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataIntegrationFlowPropsMixin
type jsiiProxy_CfnDataIntegrationFlowPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDataIntegrationFlowPropsMixin) Props() *CfnDataIntegrationFlowMixinProps {
	var returns *CfnDataIntegrationFlowMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataIntegrationFlowPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SCN::DataIntegrationFlow`.
func NewCfnDataIntegrationFlowPropsMixin(props *CfnDataIntegrationFlowMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDataIntegrationFlowPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataIntegrationFlowPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataIntegrationFlowPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SCN::DataIntegrationFlow`.
func NewCfnDataIntegrationFlowPropsMixin_Override(c CfnDataIntegrationFlowPropsMixin, props *CfnDataIntegrationFlowMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDataIntegrationFlowPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataIntegrationFlowPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataIntegrationFlowPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataIntegrationFlowPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataIntegrationFlowPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

