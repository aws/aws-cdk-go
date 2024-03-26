package awsiotsitewise


// Contains information about a composite model in an asset model.
//
// This object contains the asset property definitions that you define in the composite model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelCompositeModelProperty := &AssetModelCompositeModelProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	ComposedAssetModelId: jsii.String("composedAssetModelId"),
//   	CompositeModelProperties: []interface{}{
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
//   	Description: jsii.String("description"),
//   	ExternalId: jsii.String("externalId"),
//   	Id: jsii.String("id"),
//   	ParentAssetModelCompositeModelExternalId: jsii.String("parentAssetModelCompositeModelExternalId"),
//   	Path: []*string{
//   		jsii.String("path"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html
//
type CfnAssetModel_AssetModelCompositeModelProperty struct {
	// The name of the composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the composite model.
	//
	// For alarm composite models, this type is `AWS/ALARM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of a component model which is reused to create this composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-composedassetmodelid
	//
	ComposedAssetModelId *string `field:"optional" json:"composedAssetModelId" yaml:"composedAssetModelId"`
	// The asset property definitions for this composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-compositemodelproperties
	//
	CompositeModelProperties interface{} `field:"optional" json:"compositeModelProperties" yaml:"compositeModelProperties"`
	// The description of the composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The external ID of a composite model on this asset model.
	//
	// For more information, see [Using external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The ID of the asset model composite model.
	//
	// > This is a return value and can't be set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The external ID of the parent asset model.
	//
	// For more information, see [Using external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the *AWS IoT SiteWise User Guide* .
	//
	// > If `ParentCompositeModelExternalId` is specified, this value overrides the value of `ExternalId` , if both are included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-parentassetmodelcompositemodelexternalid
	//
	ParentAssetModelCompositeModelExternalId *string `field:"optional" json:"parentAssetModelCompositeModelExternalId" yaml:"parentAssetModelCompositeModelExternalId"`
	// The structured path to the property from the root of the asset using property names.
	//
	// Path is used as the ID if the asset model is a derived composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-path
	//
	Path *[]*string `field:"optional" json:"path" yaml:"path"`
}

