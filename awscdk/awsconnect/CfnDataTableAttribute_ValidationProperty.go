package awsconnect


// Defines validation rules for data table attribute values.
//
// Based on JSON Schema Draft 2020-12 with additional Connect-specific validations. Validation rules ensure data integrity and consistency across the data table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationProperty := &ValidationProperty{
//   	Enum: &EnumProperty{
//   		Strict: jsii.Boolean(false),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	ExclusiveMaximum: jsii.Number(123),
//   	ExclusiveMinimum: jsii.Number(123),
//   	Maximum: jsii.Number(123),
//   	MaxLength: jsii.Number(123),
//   	MaxValues: jsii.Number(123),
//   	Minimum: jsii.Number(123),
//   	MinLength: jsii.Number(123),
//   	MinValues: jsii.Number(123),
//   	MultipleOf: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html
//
type CfnDataTableAttribute_ValidationProperty struct {
	// Defines enumeration constraints for attribute values.
	//
	// Can specify a list of allowed values and whether custom values are permitted beyond the enumerated list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-enum
	//
	Enum interface{} `field:"optional" json:"enum" yaml:"enum"`
	// The largest exclusive numeric value for NUMBER value type.
	//
	// Can be provided alongside Maximum where both operate independently. Must be greater than ExclusiveMinimum and Minimum. Applies to NUMBER and values within NUMBER_LIST.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-exclusivemaximum
	//
	ExclusiveMaximum *float64 `field:"optional" json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	// The smallest exclusive numeric value for NUMBER value type.
	//
	// Can be provided alongside Minimum where both operate independently. Must be less than ExclusiveMaximum and Maximum. Applies to NUMBER and values within NUMBER_LIST.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-exclusiveminimum
	//
	ExclusiveMinimum *float64 `field:"optional" json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	// The largest inclusive numeric value for NUMBER value type.
	//
	// Can be provided alongside ExclusiveMaximum where both operate independently. Must be greater than or equal to Minimum and greater than ExclusiveMinimum. Applies to NUMBER and values within NUMBER_LIST.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// The maximum number of characters a text value can contain.
	//
	// Applies to TEXT value type and values within a TEXT_LIST. Must be greater than or equal to MinLength.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-maxlength
	//
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// The maximum number of values in a list.
	//
	// Must be an integer greater than or equal to 0 and greater than or equal to MinValues. Applies to all list types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-maxvalues
	//
	MaxValues *float64 `field:"optional" json:"maxValues" yaml:"maxValues"`
	// The smallest inclusive numeric value for NUMBER value type.
	//
	// Cannot be provided when ExclusiveMinimum is also provided. Must be less than or equal to Maximum and less than ExclusiveMaximum. Applies to NUMBER and values within NUMBER_LIST.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// The minimum number of characters a text value can contain.
	//
	// Applies to TEXT value type and values within a TEXT_LIST. Must be less than or equal to MaxLength.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-minlength
	//
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// The minimum number of values in a list.
	//
	// Must be an integer greater than or equal to 0 and less than or equal to MaxValues. Applies to all list types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-minvalues
	//
	MinValues *float64 `field:"optional" json:"minValues" yaml:"minValues"`
	// Specifies that numeric values must be multiples of this number.
	//
	// Must be greater than 0. The result of dividing a value by this multiple must result in an integer. Applies to NUMBER and values within NUMBER_LIST.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-multipleof
	//
	MultipleOf *float64 `field:"optional" json:"multipleOf" yaml:"multipleOf"`
}

