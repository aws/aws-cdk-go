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
//   definitionProperty := &definitionProperty{
//   	configuration: map[string]*string{
//   		"configurationKey": jsii.String("configuration"),
//   	},
//   	dataType: &dataTypeProperty{
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
//   	isFinal: jsii.Boolean(false),
//   	isImported: jsii.Boolean(false),
//   	isInherited: jsii.Boolean(false),
//   	isRequiredInEntity: jsii.Boolean(false),
//   	isStoredExternally: jsii.Boolean(false),
//   	isTimeSeries: jsii.Boolean(false),
//   }
//
type CfnEntity_DefinitionProperty struct {
	// `CfnEntity.DefinitionProperty.Configuration`.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// `CfnEntity.DefinitionProperty.DataType`.
	DataType interface{} `field:"optional" json:"dataType" yaml:"dataType"`
	// `CfnEntity.DefinitionProperty.DefaultValue`.
	DefaultValue interface{} `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// `CfnEntity.DefinitionProperty.IsExternalId`.
	IsExternalId interface{} `field:"optional" json:"isExternalId" yaml:"isExternalId"`
	// `CfnEntity.DefinitionProperty.IsFinal`.
	IsFinal interface{} `field:"optional" json:"isFinal" yaml:"isFinal"`
	// `CfnEntity.DefinitionProperty.IsImported`.
	IsImported interface{} `field:"optional" json:"isImported" yaml:"isImported"`
	// `CfnEntity.DefinitionProperty.IsInherited`.
	IsInherited interface{} `field:"optional" json:"isInherited" yaml:"isInherited"`
	// `CfnEntity.DefinitionProperty.IsRequiredInEntity`.
	IsRequiredInEntity interface{} `field:"optional" json:"isRequiredInEntity" yaml:"isRequiredInEntity"`
	// `CfnEntity.DefinitionProperty.IsStoredExternally`.
	IsStoredExternally interface{} `field:"optional" json:"isStoredExternally" yaml:"isStoredExternally"`
	// `CfnEntity.DefinitionProperty.IsTimeSeries`.
	IsTimeSeries interface{} `field:"optional" json:"isTimeSeries" yaml:"isTimeSeries"`
}

