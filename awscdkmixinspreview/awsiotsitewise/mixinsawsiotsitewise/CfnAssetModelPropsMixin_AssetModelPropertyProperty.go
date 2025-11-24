package mixinsawsiotsitewise


// Contains information about an asset model property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetModelPropertyProperty := &AssetModelPropertyProperty{
//   	DataType: jsii.String("dataType"),
//   	DataTypeSpec: jsii.String("dataTypeSpec"),
//   	ExternalId: jsii.String("externalId"),
//   	Id: jsii.String("id"),
//   	LogicalId: jsii.String("logicalId"),
//   	Name: jsii.String("name"),
//   	Type: &PropertyTypeProperty{
//   		Attribute: &AttributeProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   		},
//   		Metric: &MetricProperty{
//   			Expression: jsii.String("expression"),
//   			Variables: []interface{}{
//   				&ExpressionVariableProperty{
//   					Name: jsii.String("name"),
//   					Value: &VariableValueProperty{
//   						HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   						HierarchyId: jsii.String("hierarchyId"),
//   						HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   						PropertyExternalId: jsii.String("propertyExternalId"),
//   						PropertyId: jsii.String("propertyId"),
//   						PropertyLogicalId: jsii.String("propertyLogicalId"),
//   						PropertyPath: []interface{}{
//   							&PropertyPathDefinitionProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Window: &MetricWindowProperty{
//   				Tumbling: &TumblingWindowProperty{
//   					Interval: jsii.String("interval"),
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
//   						HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   						HierarchyId: jsii.String("hierarchyId"),
//   						HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   						PropertyExternalId: jsii.String("propertyExternalId"),
//   						PropertyId: jsii.String("propertyId"),
//   						PropertyLogicalId: jsii.String("propertyLogicalId"),
//   						PropertyPath: []interface{}{
//   							&PropertyPathDefinitionProperty{
//   								Name: jsii.String("name"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		TypeName: jsii.String("typeName"),
//   	},
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html
//
type CfnAssetModelPropsMixin_AssetModelPropertyProperty struct {
	// The data type of the asset model property.
	//
	// If you specify `STRUCT` , you must also specify `dataTypeSpec` to identify the type of the structure for this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The data type of the structure for this property.
	//
	// This parameter exists on properties that have the `STRUCT` data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-datatypespec
	//
	DataTypeSpec *string `field:"optional" json:"dataTypeSpec" yaml:"dataTypeSpec"`
	// The external ID of the asset property.
	//
	// For more information, see [Using external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the *AWS IoT SiteWise User Guide* .
	//
	// > One of `ExternalId` or `LogicalId` must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The ID of the property.
	//
	// > This is a return value and can't be set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The `LogicalID` of the asset model property.
	//
	// > One of `ExternalId` or `LogicalId` must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-logicalid
	//
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
	// The name of the asset model property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contains a property type, which can be one of `attribute` , `measurement` , `metric` , or `transform` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-type
	//
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// The unit of the asset model property, such as `Newtons` or `RPM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelproperty.html#cfn-iotsitewise-assetmodel-assetmodelproperty-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

