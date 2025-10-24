package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAssetModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetModelProps := &CfnAssetModelProps{
//   	AssetModelName: jsii.String("assetModelName"),
//
//   	// the properties below are optional
//   	AssetModelCompositeModels: []interface{}{
//   		&AssetModelCompositeModelProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			ComposedAssetModelId: jsii.String("composedAssetModelId"),
//   			CompositeModelProperties: []interface{}{
//   				&AssetModelPropertyProperty{
//   					DataType: jsii.String("dataType"),
//   					Name: jsii.String("name"),
//   					Type: &PropertyTypeProperty{
//   						TypeName: jsii.String("typeName"),
//
//   						// the properties below are optional
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
//
//   									// the properties below are optional
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
//   					},
//
//   					// the properties below are optional
//   					DataTypeSpec: jsii.String("dataTypeSpec"),
//   					ExternalId: jsii.String("externalId"),
//   					Id: jsii.String("id"),
//   					LogicalId: jsii.String("logicalId"),
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			ParentAssetModelCompositeModelExternalId: jsii.String("parentAssetModelCompositeModelExternalId"),
//   			Path: []*string{
//   				jsii.String("path"),
//   			},
//   		},
//   	},
//   	AssetModelDescription: jsii.String("assetModelDescription"),
//   	AssetModelExternalId: jsii.String("assetModelExternalId"),
//   	AssetModelHierarchies: []interface{}{
//   		&AssetModelHierarchyProperty{
//   			ChildAssetModelId: jsii.String("childAssetModelId"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			LogicalId: jsii.String("logicalId"),
//   		},
//   	},
//   	AssetModelProperties: []interface{}{
//   		&AssetModelPropertyProperty{
//   			DataType: jsii.String("dataType"),
//   			Name: jsii.String("name"),
//   			Type: &PropertyTypeProperty{
//   				TypeName: jsii.String("typeName"),
//
//   				// the properties below are optional
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
//
//   							// the properties below are optional
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
//   			},
//
//   			// the properties below are optional
//   			DataTypeSpec: jsii.String("dataTypeSpec"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			LogicalId: jsii.String("logicalId"),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   	AssetModelType: jsii.String("assetModelType"),
//   	EnforcedAssetModelInterfaceRelationships: []interface{}{
//   		&EnforcedAssetModelInterfaceRelationshipProperty{
//   			InterfaceAssetModelId: jsii.String("interfaceAssetModelId"),
//   			PropertyMappings: []interface{}{
//   				&EnforcedAssetModelInterfacePropertyMappingProperty{
//   					InterfaceAssetModelPropertyExternalId: jsii.String("interfaceAssetModelPropertyExternalId"),
//
//   					// the properties below are optional
//   					AssetModelPropertyExternalId: jsii.String("assetModelPropertyExternalId"),
//   					AssetModelPropertyLogicalId: jsii.String("assetModelPropertyLogicalId"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html
//
type CfnAssetModelProps struct {
	// A unique name for the asset model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
	//
	AssetModelName *string `field:"required" json:"assetModelName" yaml:"assetModelName"`
	// The composite models that are part of this asset model.
	//
	// It groups properties (such as attributes, measurements, transforms, and metrics) and child composite models that model parts of your industrial equipment. Each composite model has a type that defines the properties that the composite model supports. Use composite models to define alarms on this asset model.
	//
	// > When creating custom composite models, you need to use [CreateAssetModelCompositeModel](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModelCompositeModel.html) . For more information, see [Creating custom composite models (Components)](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-custom-composite-models.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodels
	//
	AssetModelCompositeModels interface{} `field:"optional" json:"assetModelCompositeModels" yaml:"assetModelCompositeModels"`
	// A description for the asset model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
	//
	AssetModelDescription *string `field:"optional" json:"assetModelDescription" yaml:"assetModelDescription"`
	// The external ID of the asset model.
	//
	// For more information, see [Using external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelexternalid
	//
	AssetModelExternalId *string `field:"optional" json:"assetModelExternalId" yaml:"assetModelExternalId"`
	// The hierarchy definitions of the asset model.
	//
	// Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. For more information, see [Asset hierarchies](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 10 hierarchies per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
	//
	AssetModelHierarchies interface{} `field:"optional" json:"assetModelHierarchies" yaml:"assetModelHierarchies"`
	// The property definitions of the asset model.
	//
	// For more information, see [Asset properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 200 properties per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
	//
	AssetModelProperties interface{} `field:"optional" json:"assetModelProperties" yaml:"assetModelProperties"`
	// The type of asset model.
	//
	// - *ASSET_MODEL* – (default) An asset model that you can use to create assets. Can't be included as a component in another asset model.
	// - *COMPONENT_MODEL* – A reusable component that you can include in the composite models of other asset models. You can't create assets directly from this type of asset model.
	// - *INTERFACE* – An interface is a type of model that defines a standard structure that can be applied to different asset models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeltype
	//
	AssetModelType *string `field:"optional" json:"assetModelType" yaml:"assetModelType"`
	// a list of asset model and interface relationships.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationships
	//
	EnforcedAssetModelInterfaceRelationships interface{} `field:"optional" json:"enforcedAssetModelInterfaceRelationships" yaml:"enforcedAssetModelInterfaceRelationships"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

