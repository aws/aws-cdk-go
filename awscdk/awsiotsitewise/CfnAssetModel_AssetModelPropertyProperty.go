package awsiotsitewise


// Contains information about an asset model property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelPropertyProperty := &AssetModelPropertyProperty{
//   	DataType: jsii.String("dataType"),
//   	LogicalId: jsii.String("logicalId"),
//   	Name: jsii.String("name"),
//   	Type: &PropertyTypeProperty{
//   		TypeName: jsii.String("typeName"),
//
//   		// the properties below are optional
//   		Attribute: &AttributeProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   		},
//   		Metric: &MetricProperty{
//   			Expression: jsii.String("expression"),
//   			Variables: []interface{}{
//   				&ExpressionVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: &VariableValueProperty{
//   						PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   						// the properties below are optional
//   						HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					},
//   				},
//   			},
//   			Window: &MetricWindowProperty{
//   				Tumbling: &TumblingWindowProperty{
//   					Interval: jsii.String("interval"),
//
//   					// the properties below are optional
//   					Offset: jsii.String("offset"),
//   				},
//   			},
//   		},
//   		Transform: &TransformProperty{
//   			Expression: jsii.String("expression"),
//   			Variables: []interface{}{
//   				&ExpressionVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: &VariableValueProperty{
//   						PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   						// the properties below are optional
//   						HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	DataTypeSpec: jsii.String("dataTypeSpec"),
//   	Unit: jsii.String("unit"),
//   }
//
type CfnAssetModel_AssetModelPropertyProperty struct {
	// The data type of the asset model property.
	//
	// The value can be `STRING` , `INTEGER` , `DOUBLE` , `BOOLEAN` , or `STRUCT` .
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The `LogicalID` of the asset model property.
	//
	// The maximum length is 256 characters, with the pattern `[^\\ u0000-\\ u001F\\ u007F]+` .
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model property.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains a property type, which can be one of `Attribute` , `Measurement` , `Metric` , or `Transform` .
	Type interface{} `field:"required" json:"type" yaml:"type"`
	// The data type of the structure for this property.
	//
	// This parameter exists on properties that have the `STRUCT` data type.
	DataTypeSpec *string `field:"optional" json:"dataTypeSpec" yaml:"dataTypeSpec"`
	// The unit of the asset model property, such as `Newtons` or `RPM` .
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

