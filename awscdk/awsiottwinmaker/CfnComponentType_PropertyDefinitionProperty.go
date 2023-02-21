package awsiottwinmaker


// PropertyDefinition is an object that maps strings to the property definitions in the component type.
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
//   propertyDefinitionProperty := &PropertyDefinitionProperty{
//   	Configurations: map[string]*string{
//   		"configurationsKey": jsii.String("configurations"),
//   	},
//   	DataType: &dataTypeProperty{
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
//   	IsRequiredInEntity: jsii.Boolean(false),
//   	IsStoredExternally: jsii.Boolean(false),
//   	IsTimeSeries: jsii.Boolean(false),
//   }
//
type CfnComponentType_PropertyDefinitionProperty struct {
	// A mapping that specifies configuration information about the property.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// `CfnComponentType.PropertyDefinitionProperty.DataType`.
	DataType interface{} `field:"optional" json:"dataType" yaml:"dataType"`
	// A boolean value that specifies whether the property ID comes from an external data store.
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A boolean value that specifies whether the property ID comes from an external data store.
	IsExternalId interface{} `field:"optional" json:"isExternalId" yaml:"isExternalId"`
	// A boolean value that specifies whether the property is required in an entity.
	IsRequiredInEntity interface{} `field:"optional" json:"isRequiredInEntity" yaml:"isRequiredInEntity"`
	// A boolean value that specifies whether the property is stored externally.
	IsStoredExternally interface{} `field:"optional" json:"isStoredExternally" yaml:"isStoredExternally"`
	// A boolean value that specifies whether the property consists of time series data.
	IsTimeSeries interface{} `field:"optional" json:"isTimeSeries" yaml:"isTimeSeries"`
}

