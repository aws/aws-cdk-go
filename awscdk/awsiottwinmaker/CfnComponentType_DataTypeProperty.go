package awsiottwinmaker


// An object that specifies the data type of a property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ DataTypeProperty
//   var dataValueProperty_ DataValueProperty
//   var relationshipValue interface{}
//
//   dataTypeProperty := &DataTypeProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AllowedValues: []interface{}{
//   		&DataValueProperty{
//   			BooleanValue: jsii.Boolean(false),
//   			DoubleValue: jsii.Number(123),
//   			Expression: jsii.String("expression"),
//   			IntegerValue: jsii.Number(123),
//   			ListValue: []interface{}{
//   				dataValueProperty_,
//   			},
//   			LongValue: jsii.Number(123),
//   			MapValue: map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			RelationshipValue: relationshipValue,
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	NestedType: dataTypeProperty_,
//   	Relationship: &RelationshipProperty{
//   		RelationshipType: jsii.String("relationshipType"),
//   		TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   	},
//   	UnitOfMeasure: jsii.String("unitOfMeasure"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-datatype.html
//
type CfnComponentType_DataTypeProperty struct {
	// The underlying type of the data type.
	//
	// Valid Values: `RELATIONSHIP | STRING | LONG | BOOLEAN | INTEGER | DOUBLE | LIST | MAP`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-datatype.html#cfn-iottwinmaker-componenttype-datatype-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The allowed values for this data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-datatype.html#cfn-iottwinmaker-componenttype-datatype-allowedvalues
	//
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The nested type in the data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-datatype.html#cfn-iottwinmaker-componenttype-datatype-nestedtype
	//
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// A relationship that associates a component with another component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-datatype.html#cfn-iottwinmaker-componenttype-datatype-relationship
	//
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// The unit of measure used in this data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-componenttype-datatype.html#cfn-iottwinmaker-componenttype-datatype-unitofmeasure
	//
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

