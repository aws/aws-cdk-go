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
//   			CompositeModelProperties: []interface{}{
//   				&AssetModelPropertyProperty{
//   					DataType: jsii.String("dataType"),
//   					LogicalId: jsii.String("logicalId"),
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
//   										PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
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
//   										PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					DataTypeSpec: jsii.String("dataTypeSpec"),
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	AssetModelDescription: jsii.String("assetModelDescription"),
//   	AssetModelHierarchies: []interface{}{
//   		&AssetModelHierarchyProperty{
//   			ChildAssetModelId: jsii.String("childAssetModelId"),
//   			LogicalId: jsii.String("logicalId"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	AssetModelProperties: []interface{}{
//   		&AssetModelPropertyProperty{
//   			DataType: jsii.String("dataType"),
//   			LogicalId: jsii.String("logicalId"),
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
//   								PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
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
//   								PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			DataTypeSpec: jsii.String("dataTypeSpec"),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssetModelProps struct {
	// A unique, friendly name for the asset model.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	AssetModelName *string `field:"required" json:"assetModelName" yaml:"assetModelName"`
	// The composite asset models that are part of this asset model.
	//
	// Composite asset models are asset models that contain specific properties. Each composite model has a type that defines the properties that the composite model supports. You can use composite asset models to define alarms on this asset model.
	AssetModelCompositeModels interface{} `field:"optional" json:"assetModelCompositeModels" yaml:"assetModelCompositeModels"`
	// A description for the asset model.
	AssetModelDescription *string `field:"optional" json:"assetModelDescription" yaml:"assetModelDescription"`
	// The hierarchy definitions of the asset model.
	//
	// Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. For more information, see [Defining relationships between assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 10 hierarchies per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	AssetModelHierarchies interface{} `field:"optional" json:"assetModelHierarchies" yaml:"assetModelHierarchies"`
	// The property definitions of the asset model.
	//
	// For more information, see [Defining data properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-properties.html) in the *AWS IoT SiteWise User Guide* .
	//
	// You can specify up to 200 properties per asset model. For more information, see [Quotas](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the *AWS IoT SiteWise User Guide* .
	AssetModelProperties interface{} `field:"optional" json:"assetModelProperties" yaml:"assetModelProperties"`
	// A list of key-value pairs that contain metadata for the asset.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

