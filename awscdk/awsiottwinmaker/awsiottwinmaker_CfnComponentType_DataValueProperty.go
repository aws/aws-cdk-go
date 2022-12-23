package awsiottwinmaker


// An object that specifies a value for a property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   dataValueProperty := &dataValueProperty{
//   	booleanValue: jsii.Boolean(false),
//   	doubleValue: jsii.Number(123),
//   	expression: jsii.String("expression"),
//   	integerValue: jsii.Number(123),
//   	listValue: []interface{}{
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
//   	longValue: jsii.Number(123),
//   	mapValue: map[string]interface{}{
//   		"mapValueKey": &dataValueProperty{
//   			"booleanValue": jsii.Boolean(false),
//   			"doubleValue": jsii.Number(123),
//   			"expression": jsii.String("expression"),
//   			"integerValue": jsii.Number(123),
//   			"listValue": []interface{}{
//   				dataValueProperty_,
//   			},
//   			"longValue": jsii.Number(123),
//   			"mapValue": map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			"relationshipValue": relationshipValue,
//   			"stringValue": jsii.String("stringValue"),
//   		},
//   	},
//   	relationshipValue: relationshipValue,
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnComponentType_DataValueProperty struct {
	// A boolean value.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// A double value.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// An expression that produces the value.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// An integer value.
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// A list of multiple values.
	ListValue interface{} `field:"optional" json:"listValue" yaml:"listValue"`
	// A long value.
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// An object that maps strings to multiple `DataValue` objects.
	MapValue interface{} `field:"optional" json:"mapValue" yaml:"mapValue"`
	// A value that relates a component to another component.
	RelationshipValue interface{} `field:"optional" json:"relationshipValue" yaml:"relationshipValue"`
	// A string value.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

