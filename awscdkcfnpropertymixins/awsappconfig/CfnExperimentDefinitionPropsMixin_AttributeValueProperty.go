package awsappconfig


// A typed attribute value for a treatment flag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   attributeValueProperty := &AttributeValueProperty{
//   	BooleanValue: jsii.Boolean(false),
//   	NumberArray: []interface{}{
//   		jsii.Number(123),
//   	},
//   	NumberValue: jsii.Number(123),
//   	StringArray: []*string{
//   		jsii.String("stringArray"),
//   	},
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentdefinition-attributevalue.html
//
type CfnExperimentDefinitionPropsMixin_AttributeValueProperty struct {
	// A boolean value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentdefinition-attributevalue.html#cfn-appconfig-experimentdefinition-attributevalue-booleanvalue
	//
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// An array of numeric values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentdefinition-attributevalue.html#cfn-appconfig-experimentdefinition-attributevalue-numberarray
	//
	NumberArray interface{} `field:"optional" json:"numberArray" yaml:"numberArray"`
	// A numeric value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentdefinition-attributevalue.html#cfn-appconfig-experimentdefinition-attributevalue-numbervalue
	//
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// An array of string values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentdefinition-attributevalue.html#cfn-appconfig-experimentdefinition-attributevalue-stringarray
	//
	StringArray *[]*string `field:"optional" json:"stringArray" yaml:"stringArray"`
	// A string value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-experimentdefinition-attributevalue.html#cfn-appconfig-experimentdefinition-attributevalue-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

