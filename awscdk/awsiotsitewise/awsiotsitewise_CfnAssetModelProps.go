package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAssetModel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetModelProps := &cfnAssetModelProps{
//   	assetModelName: jsii.String("assetModelName"),
//
//   	// the properties below are optional
//   	assetModelCompositeModels: []interface{}{
//   		&assetModelCompositeModelProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			compositeModelProperties: []interface{}{
//   				&assetModelPropertyProperty{
//   					dataType: jsii.String("dataType"),
//   					logicalId: jsii.String("logicalId"),
//   					name: jsii.String("name"),
//   					type: &propertyTypeProperty{
//   						typeName: jsii.String("typeName"),
//
//   						// the properties below are optional
//   						attribute: &attributeProperty{
//   							defaultValue: jsii.String("defaultValue"),
//   						},
//   						metric: &metricProperty{
//   							expression: jsii.String("expression"),
//   							variables: []interface{}{
//   								&expressionVariableProperty{
//   									name: jsii.String("name"),
//   									value: &variableValueProperty{
//   										propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   							window: &metricWindowProperty{
//   								tumbling: &tumblingWindowProperty{
//   									interval: jsii.String("interval"),
//
//   									// the properties below are optional
//   									offset: jsii.String("offset"),
//   								},
//   							},
//   						},
//   						transform: &transformProperty{
//   							expression: jsii.String("expression"),
//   							variables: []interface{}{
//   								&expressionVariableProperty{
//   									name: jsii.String("name"),
//   									value: &variableValueProperty{
//   										propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   										// the properties below are optional
//   										hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					dataTypeSpec: jsii.String("dataTypeSpec"),
//   					unit: jsii.String("unit"),
//   				},
//   			},
//   			description: jsii.String("description"),
//   		},
//   	},
//   	assetModelDescription: jsii.String("assetModelDescription"),
//   	assetModelHierarchies: []interface{}{
//   		&assetModelHierarchyProperty{
//   			childAssetModelId: jsii.String("childAssetModelId"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   		},
//   	},
//   	assetModelProperties: []interface{}{
//   		&assetModelPropertyProperty{
//   			dataType: jsii.String("dataType"),
//   			logicalId: jsii.String("logicalId"),
//   			name: jsii.String("name"),
//   			type: &propertyTypeProperty{
//   				typeName: jsii.String("typeName"),
//
//   				// the properties below are optional
//   				attribute: &attributeProperty{
//   					defaultValue: jsii.String("defaultValue"),
//   				},
//   				metric: &metricProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   					window: &metricWindowProperty{
//   						tumbling: &tumblingWindowProperty{
//   							interval: jsii.String("interval"),
//
//   							// the properties below are optional
//   							offset: jsii.String("offset"),
//   						},
//   					},
//   				},
//   				transform: &transformProperty{
//   					expression: jsii.String("expression"),
//   					variables: []interface{}{
//   						&expressionVariableProperty{
//   							name: jsii.String("name"),
//   							value: &variableValueProperty{
//   								propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   								// the properties below are optional
//   								hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			dataTypeSpec: jsii.String("dataTypeSpec"),
//   			unit: jsii.String("unit"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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

