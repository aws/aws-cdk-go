package awsiottwinmaker


// An object that specifies the data type of a property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   dataTypeProperty := &dataTypeProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AllowedValues: []interface{}{
//   		&dataValueProperty{
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
//   	NestedType: &dataTypeProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		AllowedValues: []interface{}{
//   			&dataValueProperty{
//   				BooleanValue: jsii.Boolean(false),
//   				DoubleValue: jsii.Number(123),
//   				Expression: jsii.String("expression"),
//   				IntegerValue: jsii.Number(123),
//   				ListValue: []interface{}{
//   					dataValueProperty_,
//   				},
//   				LongValue: jsii.Number(123),
//   				MapValue: map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				RelationshipValue: relationshipValue,
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		NestedType: dataTypeProperty_,
//   		Relationship: &RelationshipProperty{
//   			RelationshipType: jsii.String("relationshipType"),
//   			TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   		},
//   		UnitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	Relationship: &RelationshipProperty{
//   		RelationshipType: jsii.String("relationshipType"),
//   		TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   	},
//   	UnitOfMeasure: jsii.String("unitOfMeasure"),
//   }
//
type CfnComponentType_DataTypeProperty struct {
	// The underlying type of the data type.
	//
	// Valid Values: `RELATIONSHIP | STRING | LONG | BOOLEAN | INTEGER | DOUBLE | LIST | MAP`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The allowed values for this data type.
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The nested type in the data type.
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// A relationship that associates a component with another component.
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// The unit of measure used in this data type.
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

