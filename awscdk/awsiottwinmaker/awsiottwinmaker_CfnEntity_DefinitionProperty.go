package awsiottwinmaker


// The entity definition.
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
//   definitionProperty := &DefinitionProperty{
//   	Configuration: map[string]*string{
//   		"configurationKey": jsii.String("configuration"),
//   	},
//   	DataType: &dataTypeProperty{
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
//   		Type: jsii.String("type"),
//   		UnitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	DefaultValue: &dataValueProperty{
//   		BooleanValue: jsii.Boolean(false),
//   		DoubleValue: jsii.Number(123),
//   		Expression: jsii.String("expression"),
//   		IntegerValue: jsii.Number(123),
//   		ListValue: []interface{}{
//   			dataValueProperty_,
//   		},
//   		LongValue: jsii.Number(123),
//   		MapValue: map[string]interface{}{
//   			"mapValueKey": dataValueProperty_,
//   		},
//   		RelationshipValue: relationshipValue,
//   		StringValue: jsii.String("stringValue"),
//   	},
//   	IsExternalId: jsii.Boolean(false),
//   	IsFinal: jsii.Boolean(false),
//   	IsImported: jsii.Boolean(false),
//   	IsInherited: jsii.Boolean(false),
//   	IsRequiredInEntity: jsii.Boolean(false),
//   	IsStoredExternally: jsii.Boolean(false),
//   	IsTimeSeries: jsii.Boolean(false),
//   }
//
type CfnEntity_DefinitionProperty struct {
	// The configuration.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The data type.
	DataType interface{} `field:"optional" json:"dataType" yaml:"dataType"`
	// The default value.
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Displays if the entity has a external Id.
	IsExternalId interface{} `field:"optional" json:"isExternalId" yaml:"isExternalId"`
	// Displays if the entity is final.
	IsFinal interface{} `field:"optional" json:"isFinal" yaml:"isFinal"`
	// Displays if the entity is imported.
	IsImported interface{} `field:"optional" json:"isImported" yaml:"isImported"`
	// Displays if the entity is inherited.
	IsInherited interface{} `field:"optional" json:"isInherited" yaml:"isInherited"`
	// Displays if the entity is a required entity.
	IsRequiredInEntity interface{} `field:"optional" json:"isRequiredInEntity" yaml:"isRequiredInEntity"`
	// Displays if the entity is tored externally.
	IsStoredExternally interface{} `field:"optional" json:"isStoredExternally" yaml:"isStoredExternally"`
	// Displays if the entity.
	IsTimeSeries interface{} `field:"optional" json:"isTimeSeries" yaml:"isTimeSeries"`
}

