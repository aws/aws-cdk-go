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
//   	CompositeModelProperties: []interface{}{
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
//   	Description: jsii.String("description"),
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
	// The asset property definitions for this composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-compositemodelproperties
	//
	CompositeModelProperties interface{} `field:"optional" json:"compositeModelProperties" yaml:"compositeModelProperties"`
	// The description of the composite model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelcompositemodel.html#cfn-iotsitewise-assetmodel-assetmodelcompositemodel-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

