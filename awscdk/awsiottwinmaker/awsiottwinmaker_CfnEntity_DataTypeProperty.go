package awsiottwinmaker


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
//   		type: jsii.String("type"),
//   		unitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	relationship: &relationshipProperty{
//   		relationshipType: jsii.String("relationshipType"),
//   		targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   	},
//   	type: jsii.String("type"),
//   	unitOfMeasure: jsii.String("unitOfMeasure"),
//   }
//
type CfnEntity_DataTypeProperty struct {
	// `CfnEntity.DataTypeProperty.AllowedValues`.
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// `CfnEntity.DataTypeProperty.NestedType`.
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// `CfnEntity.DataTypeProperty.Relationship`.
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// `CfnEntity.DataTypeProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// `CfnEntity.DataTypeProperty.UnitOfMeasure`.
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

