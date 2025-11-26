package previewawssecurityhubmixins


// An object that includes the data type of a security control parameter and its current value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterValueProperty := &ParameterValueProperty{
//   	Boolean: jsii.Boolean(false),
//   	Double: jsii.Number(123),
//   	Enum: jsii.String("enum"),
//   	EnumList: []*string{
//   		jsii.String("enumList"),
//   	},
//   	Integer: jsii.Number(123),
//   	IntegerList: []interface{}{
//   		jsii.Number(123),
//   	},
//   	String: jsii.String("string"),
//   	StringList: []*string{
//   		jsii.String("stringList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html
//
type CfnSecurityControlPropsMixin_ParameterValueProperty struct {
	// A control parameter that is a boolean.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-boolean
	//
	Boolean interface{} `field:"optional" json:"boolean" yaml:"boolean"`
	// A control parameter that is a double.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-double
	//
	Double *float64 `field:"optional" json:"double" yaml:"double"`
	// A control parameter that is an enum.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-enum
	//
	Enum *string `field:"optional" json:"enum" yaml:"enum"`
	// A control parameter that is a list of enums.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-enumlist
	//
	EnumList *[]*string `field:"optional" json:"enumList" yaml:"enumList"`
	// A control parameter that is an integer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-integer
	//
	Integer *float64 `field:"optional" json:"integer" yaml:"integer"`
	// A control parameter that is a list of integers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-integerlist
	//
	IntegerList interface{} `field:"optional" json:"integerList" yaml:"integerList"`
	// A control parameter that is a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-string
	//
	String *string `field:"optional" json:"string" yaml:"string"`
	// A control parameter that is a list of strings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parametervalue.html#cfn-securityhub-securitycontrol-parametervalue-stringlist
	//
	StringList *[]*string `field:"optional" json:"stringList" yaml:"stringList"`
}

