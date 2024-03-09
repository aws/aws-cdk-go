package awsiotsitewise


// Contains information about a composite model in an asset model.
//
// This object contains the asset property definitions that you define in the composite model. You can use composite asset models to define alarms on this asset model.
//
// If you use the `AssetModelCompositeModel` property to create an alarm, you must use the following information to define three asset model properties:
//
// - Use an asset model property to specify the alarm type.
//
// - The name must be `AWS/ALARM_TYPE` .
// - The data type must be `STRING` .
// - For the `Type` property, the type name must be `Attribute` and the default value must be `IOT_EVENTS` .
// - Use an asset model property to specify the alarm source.
//
// - The name must be `AWS/ALARM_SOURCE` .
// - The data type must be `STRING` .
// - For the `Type` property, the type name must be `Attribute` and the default value must be the ARN of the alarm model that you created in AWS IoT Events .
//
// > For the ARN of the alarm model, you can use the `Fn::Sub` intrinsic function to substitute the `AWS::Partition` , `AWS::Region` , and `AWS::AccountId` variables in an input string with values that you specify.
// >
// > For example, `Fn::Sub: "arn:${AWS::Partition}:iotevents:${AWS::Region}:${AWS::AccountId}:alarmModel/TestAlarmModel"` .
// >
// > Replace `TestAlarmModel` with the name of your alarm model.
// >
// > For more information about using the `Fn::Sub` intrinsic function, see [Fn::Sub](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-sub.html) .
// - Use an asset model property to specify the state of the alarm.
//
// - The name must be `AWS/ALARM_STATE` .
// - The data type must be `STRUCT` .
// - The `DataTypeSpec` value must be `AWS/ALARM_STATE` .
// - For the `Type` property, the type name must be `Measurement` .
//
// At the bottom of this page, we provide a YAML example that you can modify to create an alarm.
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
	// The component model ID for which the composite model is composed of.
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
	// The External ID of the composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The Actual ID of the composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The parent composite model External ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-parentassetmodelcompositemodelexternalid
	//
	ParentAssetModelCompositeModelExternalId *string `field:"optional" json:"parentAssetModelCompositeModelExternalId" yaml:"parentAssetModelCompositeModelExternalId"`
	// The path of the composite model.
	//
	// This is only for derived composite models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-path
	//
	Path *[]*string `field:"optional" json:"path" yaml:"path"`
}

