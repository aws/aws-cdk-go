package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnDataTableAttributePropsMixin_ValidationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-enum
	//
	Enum interface{} `field:"optional" json:"enum" yaml:"enum"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-exclusivemaximum
	//
	ExclusiveMaximum *float64 `field:"optional" json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-exclusiveminimum
	//
	ExclusiveMinimum *float64 `field:"optional" json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-maximum
	//
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-maxlength
	//
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-maxvalues
	//
	MaxValues *float64 `field:"optional" json:"maxValues" yaml:"maxValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-minimum
	//
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-minlength
	//
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-minvalues
	//
	MinValues *float64 `field:"optional" json:"minValues" yaml:"minValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-validation.html#cfn-connect-datatableattribute-validation-multipleof
	//
	MultipleOf *float64 `field:"optional" json:"multipleOf" yaml:"multipleOf"`
}

