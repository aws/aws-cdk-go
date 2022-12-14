package awsiottwinmaker


// An object that sets information about a property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//
//   propertyProperty := &propertyProperty{
//   	definition: &definitionProperty{
//   		configuration: map[string]*string{
//   			"configurationKey": jsii.String("configuration"),
//   		},
//   		dataType: &dataTypeProperty{
//   			allowedValues: []interface{}{
//   				&dataValueProperty{
//   					booleanValue: jsii.Boolean(false),
//   					doubleValue: jsii.Number(123),
//   					expression: jsii.String("expression"),
//   					integerValue: jsii.Number(123),
//   					listValue: []interface{}{
//   						dataValueProperty_,
//   					},
//   					longValue: jsii.Number(123),
//   					mapValue: map[string]interface{}{
//   						"mapValueKey": dataValueProperty_,
//   					},
//   					relationshipValue: &relationshipValueProperty{
//   						targetComponentName: jsii.String("targetComponentName"),
//   						targetEntityId: jsii.String("targetEntityId"),
//   					},
//   					stringValue: jsii.String("stringValue"),
//   				},
//   			},
//   			nestedType: dataTypeProperty_,
//   			relationship: &relationshipProperty{
//   				relationshipType: jsii.String("relationshipType"),
//   				targetComponentTypeId: jsii.String("targetComponentTypeId"),
//   			},
//   			type: jsii.String("type"),
//   			unitOfMeasure: jsii.String("unitOfMeasure"),
//   		},
//   		defaultValue: &dataValueProperty{
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
//   			relationshipValue: &relationshipValueProperty{
//   				targetComponentName: jsii.String("targetComponentName"),
//   				targetEntityId: jsii.String("targetEntityId"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   		isExternalId: jsii.Boolean(false),
//   		isFinal: jsii.Boolean(false),
//   		isImported: jsii.Boolean(false),
//   		isInherited: jsii.Boolean(false),
//   		isRequiredInEntity: jsii.Boolean(false),
//   		isStoredExternally: jsii.Boolean(false),
//   		isTimeSeries: jsii.Boolean(false),
//   	},
//   	value: &dataValueProperty{
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
//   		relationshipValue: &relationshipValueProperty{
//   			targetComponentName: jsii.String("targetComponentName"),
//   			targetEntityId: jsii.String("targetEntityId"),
//   		},
//   		stringValue: jsii.String("stringValue"),
//   	},
//   }
//
type CfnEntity_PropertyProperty struct {
	// An object that specifies information about a property.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// An object that contains information about a value for a time series property.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

