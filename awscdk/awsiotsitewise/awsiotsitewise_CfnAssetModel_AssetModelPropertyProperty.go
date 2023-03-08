package awsiotsitewise


// Contains information about an asset model property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelPropertyProperty := &assetModelPropertyProperty{
//   	dataType: jsii.String("dataType"),
//   	logicalId: jsii.String("logicalId"),
//   	name: jsii.String("name"),
//   	type: &propertyTypeProperty{
//   		typeName: jsii.String("typeName"),
//
//   		// the properties below are optional
//   		attribute: &attributeProperty{
//   			defaultValue: jsii.String("defaultValue"),
//   		},
//   		metric: &metricProperty{
//   			expression: jsii.String("expression"),
//   			variables: []interface{}{
//   				&expressionVariableProperty{
//   					name: jsii.String("name"),
//   					value: &variableValueProperty{
//   						propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   						// the properties below are optional
//   						hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					},
//   				},
//   			},
//   			window: &metricWindowProperty{
//   				tumbling: &tumblingWindowProperty{
//   					interval: jsii.String("interval"),
//
//   					// the properties below are optional
//   					offset: jsii.String("offset"),
//   				},
//   			},
//   		},
//   		transform: &transformProperty{
//   			expression: jsii.String("expression"),
//   			variables: []interface{}{
//   				&expressionVariableProperty{
//   					name: jsii.String("name"),
//   					value: &variableValueProperty{
//   						propertyLogicalId: jsii.String("propertyLogicalId"),
//
//   						// the properties below are optional
//   						hierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   					},
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	dataTypeSpec: jsii.String("dataTypeSpec"),
//   	unit: jsii.String("unit"),
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

