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

