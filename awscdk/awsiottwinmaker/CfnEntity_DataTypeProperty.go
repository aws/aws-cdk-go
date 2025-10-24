package awsiottwinmaker


// The entity data type.
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
//   	Type: jsii.String("type"),
//   	UnitOfMeasure: jsii.String("unitOfMeasure"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html
//
type CfnEntity_DataTypeProperty struct {
	// The allowed values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-allowedvalues
	//
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The nested type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-nestedtype
	//
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// The relationship.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-relationship
	//
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// The entity type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The unit of measure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datatype.html#cfn-iottwinmaker-entity-datatype-unitofmeasure
	//
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

