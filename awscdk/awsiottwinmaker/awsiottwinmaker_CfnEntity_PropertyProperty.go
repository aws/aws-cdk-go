package awsiottwinmaker


// An object that sets information about a property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var relationshipValue interface{}
//
//   propertyProperty := &propertyProperty{
//   	definition: definition,
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
//   		relationshipValue: relationshipValue,
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

