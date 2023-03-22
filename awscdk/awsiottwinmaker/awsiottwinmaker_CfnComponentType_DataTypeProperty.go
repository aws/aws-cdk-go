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
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	allowedValues: []interface{}{
//   		&dataValueProperty{
//   			booleanValue: jsii.Boolean(false),
//   			doubleValue: jsii.Number(123),
//   			expression: jsii.String("expression"),
//   			integerValue: jsii.Number(123),
//   			listValue: []interface{}{
//   				dataValueProperty_,
//   			},
//   			longValue: jsii.Number(123),
//   			mapValue: map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			relationshipValue: relationshipValue,
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	nestedType: &dataTypeProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		allowedValues: []interface{}{
//   			&dataValueProperty{
//   				booleanValue: jsii.Boolean(false),
//   				doubleValue: jsii.Number(123),
//   				expression: jsii.String("expression"),
//   				integerValue: jsii.Number(123),
//   				listValue: []interface{}{
//   					dataValueProperty_,
//   				},
//   				longValue: jsii.Number(123),
//   				mapValue: map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				relationshipValue: relationshipValue,
//   				stringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		nestedType: dataTypeProperty_,
//   		relationship: &relationshipProperty{
//   			relationshipType: jsii.String("relationshipType"),
//   			targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   		},
//   		unitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	relationship: &relationshipProperty{
//   		relationshipType: jsii.String("relationshipType"),
//   		targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   	},
//   	unitOfMeasure: jsii.String("unitOfMeasure"),
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

