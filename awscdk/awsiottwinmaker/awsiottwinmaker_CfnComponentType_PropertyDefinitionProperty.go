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
//   propertyDefinitionProperty := &propertyDefinitionProperty{
//   	configurations: map[string]*string{
//   		"configurationsKey": jsii.String("configurations"),
//   	},
//   	dataType: &dataTypeProperty{
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
//   	defaultValue: &dataValueProperty{
//   		booleanValue: jsii.Boolean(false),
//   		doubleValue: jsii.Number(123),
//   		expression: jsii.String("expression"),
//   		integerValue: jsii.Number(123),
//   		listValue: []interface{}{
//   			dataValueProperty_,
//   		},
//   		longValue: jsii.Number(123),
//   		mapValue: map[string]interface{}{
//   			"mapValueKey": dataValueProperty_,
//   		},
//   		relationshipValue: relationshipValue,
//   		stringValue: jsii.String("stringValue"),
//   	},
//   	isExternalId: jsii.Boolean(false),
//   	isRequiredInEntity: jsii.Boolean(false),
//   	isStoredExternally: jsii.Boolean(false),
//   	isTimeSeries: jsii.Boolean(false),
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

