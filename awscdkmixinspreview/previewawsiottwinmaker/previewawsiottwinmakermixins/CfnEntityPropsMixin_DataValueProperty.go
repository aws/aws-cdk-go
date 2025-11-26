package previewawsiottwinmakermixins


// An object that specifies a value for a property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var dataValueProperty_ DataValueProperty
//   var relationshipValue interface{}
//
//   dataValueProperty := &DataValueProperty{
//   	BooleanValue: jsii.Boolean(false),
//   	DoubleValue: jsii.Number(123),
//   	Expression: jsii.String("expression"),
//   	IntegerValue: jsii.Number(123),
//   	ListValue: []interface{}{
//   		dataValueProperty_,
//   	},
//   	LongValue: jsii.Number(123),
//   	MapValue: map[string]interface{}{
//   		"mapValueKey": dataValueProperty_,
//   	},
//   	RelationshipValue: relationshipValue,
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html
//
type CfnEntityPropsMixin_DataValueProperty struct {
	// A boolean value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-booleanvalue
	//
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// A double value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-doublevalue
	//
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// An expression that produces the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// An integer value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-integervalue
	//
	IntegerValue *float64 `field:"optional" json:"integerValue" yaml:"integerValue"`
	// A list of multiple values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-listvalue
	//
	ListValue interface{} `field:"optional" json:"listValue" yaml:"listValue"`
	// A long value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-longvalue
	//
	LongValue *float64 `field:"optional" json:"longValue" yaml:"longValue"`
	// An object that maps strings to multiple DataValue objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-mapvalue
	//
	MapValue interface{} `field:"optional" json:"mapValue" yaml:"mapValue"`
	// A value that relates a component to another component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-relationshipvalue
	//
	RelationshipValue interface{} `field:"optional" json:"relationshipValue" yaml:"relationshipValue"`
	// A string value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iottwinmaker-entity-datavalue.html#cfn-iottwinmaker-entity-datavalue-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

