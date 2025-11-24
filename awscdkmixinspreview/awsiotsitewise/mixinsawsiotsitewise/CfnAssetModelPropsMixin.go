package mixinsawsiotsitewise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsiotsitewise/mixinsawsiotsitewise/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an asset model from specified property and hierarchy definitions.
//
// You create assets from asset models. With asset models, you can easily create assets of the same type that have standardized definitions. Each asset created from a model inherits the asset model's property and hierarchy definitions. For more information, see [Defining asset models](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html) in the *AWS IoT SiteWise User Guide* .
//
// You can create three types of asset models, `ASSET_MODEL` , `COMPONENT_MODEL` , or an `INTERFACE` .
//
// - *ASSET_MODEL* – (default) An asset model that you can use to create assets. Can't be included as a component in another asset model.
// - *COMPONENT_MODEL* – A reusable component that you can include in the composite models of other asset models. You can't create assets directly from this type of asset model.
// - *INTERFACE* – An interface is a type of model that defines a standard structure that can be applied to different asset models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssetModelPropsMixin := awscdkmixinspreview.Mixins.NewCfnAssetModelPropsMixin(&CfnAssetModelMixinProps{
//   	AssetModelCompositeModels: []interface{}{
//   		&AssetModelCompositeModelProperty{
//   			ComposedAssetModelId: jsii.String("composedAssetModelId"),
//   			CompositeModelProperties: []interface{}{
//   				&AssetModelPropertyProperty{
//   					DataType: jsii.String("dataType"),
//   					DataTypeSpec: jsii.String("dataTypeSpec"),
//   					ExternalId: jsii.String("externalId"),
//   					Id: jsii.String("id"),
//   					LogicalId: jsii.String("logicalId"),
//   					Name: jsii.String("name"),
//   					Type: &PropertyTypeProperty{
//   						Attribute: &AttributeProperty{
//   							DefaultValue: jsii.String("defaultValue"),
//   						},
//   						Metric: &MetricProperty{
//   							Expression: jsii.String("expression"),
//   							Variables: []interface{}{
//   								&ExpressionVariableProperty{
//   									Name: jsii.String("name"),
//   									Value: &VariableValueProperty{
//   										HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   										HierarchyId: jsii.String("hierarchyId"),
//   										HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   										PropertyExternalId: jsii.String("propertyExternalId"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyLogicalId: jsii.String("propertyLogicalId"),
//   										PropertyPath: []interface{}{
//   											&PropertyPathDefinitionProperty{
//   												Name: jsii.String("name"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Window: &MetricWindowProperty{
//   								Tumbling: &TumblingWindowProperty{
//   									Interval: jsii.String("interval"),
//   									Offset: jsii.String("offset"),
//   								},
//   							},
//   						},
//   						Transform: &TransformProperty{
//   							Expression: jsii.String("expression"),
//   							Variables: []interface{}{
//   								&ExpressionVariableProperty{
//   									Name: jsii.String("name"),
//   									Value: &VariableValueProperty{
//   										HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   										HierarchyId: jsii.String("hierarchyId"),
//   										HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   										PropertyExternalId: jsii.String("propertyExternalId"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyLogicalId: jsii.String("propertyLogicalId"),
//   										PropertyPath: []interface{}{
//   											&PropertyPathDefinitionProperty{
//   												Name: jsii.String("name"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   						},
//   						TypeName: jsii.String("typeName"),
//   					},
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			Name: jsii.String("name"),
//   			ParentAssetModelCompositeModelExternalId: jsii.String("parentAssetModelCompositeModelExternalId"),
//   			Path: []*string{
//   				jsii.String("path"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	AssetModelDescription: jsii.String("assetModelDescription"),
//   	AssetModelExternalId: jsii.String("assetModelExternalId"),
//   	AssetModelHierarchies: []interface{}{
//   		&AssetModelHierarchyProperty{
//   			ChildAssetModelId: jsii.String("childAssetModelId"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			LogicalId: jsii.String("logicalId"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	AssetModelName: jsii.String("assetModelName"),
//   	AssetModelProperties: []interface{}{
//   		&AssetModelPropertyProperty{
//   			DataType: jsii.String("dataType"),
//   			DataTypeSpec: jsii.String("dataTypeSpec"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			LogicalId: jsii.String("logicalId"),
//   			Name: jsii.String("name"),
//   			Type: &PropertyTypeProperty{
//   				Attribute: &AttributeProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   				},
//   				Metric: &MetricProperty{
//   					Expression: jsii.String("expression"),
//   					Variables: []interface{}{
//   						&ExpressionVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: &VariableValueProperty{
//   								HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   								HierarchyId: jsii.String("hierarchyId"),
//   								HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   								PropertyExternalId: jsii.String("propertyExternalId"),
//   								PropertyId: jsii.String("propertyId"),
//   								PropertyLogicalId: jsii.String("propertyLogicalId"),
//   								PropertyPath: []interface{}{
//   									&PropertyPathDefinitionProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					Window: &MetricWindowProperty{
//   						Tumbling: &TumblingWindowProperty{
//   							Interval: jsii.String("interval"),
//   							Offset: jsii.String("offset"),
//   						},
//   					},
//   				},
//   				Transform: &TransformProperty{
//   					Expression: jsii.String("expression"),
//   					Variables: []interface{}{
//   						&ExpressionVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: &VariableValueProperty{
//   								HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   								HierarchyId: jsii.String("hierarchyId"),
//   								HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   								PropertyExternalId: jsii.String("propertyExternalId"),
//   								PropertyId: jsii.String("propertyId"),
//   								PropertyLogicalId: jsii.String("propertyLogicalId"),
//   								PropertyPath: []interface{}{
//   									&PropertyPathDefinitionProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				TypeName: jsii.String("typeName"),
//   			},
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   	AssetModelType: jsii.String("assetModelType"),
//   	EnforcedAssetModelInterfaceRelationships: []interface{}{
//   		&EnforcedAssetModelInterfaceRelationshipProperty{
//   			InterfaceAssetModelId: jsii.String("interfaceAssetModelId"),
//   			PropertyMappings: []interface{}{
//   				&EnforcedAssetModelInterfacePropertyMappingProperty{
//   					AssetModelPropertyExternalId: jsii.String("assetModelPropertyExternalId"),
//   					AssetModelPropertyLogicalId: jsii.String("assetModelPropertyLogicalId"),
//   					InterfaceAssetModelPropertyExternalId: jsii.String("interfaceAssetModelPropertyExternalId"),
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html
//
type CfnAssetModelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAssetModelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssetModelPropsMixin
type jsiiProxy_CfnAssetModelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAssetModelPropsMixin) Props() *CfnAssetModelMixinProps {
	var returns *CfnAssetModelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTSiteWise::AssetModel`.
func NewCfnAssetModelPropsMixin(props *CfnAssetModelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAssetModelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssetModelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssetModelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAssetModelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTSiteWise::AssetModel`.
func NewCfnAssetModelPropsMixin_Override(c CfnAssetModelPropsMixin, props *CfnAssetModelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAssetModelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAssetModelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetModelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAssetModelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssetModelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotsitewise.mixins.CfnAssetModelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssetModelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAssetModelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

