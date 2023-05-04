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
//   propertyProperty := &PropertyProperty{
//   	Definition: definition,
//   	Value: &dataValueProperty{
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
//   }
//
type CfnEntity_PropertyProperty struct {
	// An object that specifies information about a property.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// An object that contains information about a value for a time series property.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

