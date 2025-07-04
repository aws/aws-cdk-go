package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputConversionProperty := &InputConversionProperty{
//   	FromFormat: jsii.String("fromFormat"),
//
//   	// the properties below are optional
//   	AdvancedOptions: &AdvancedOptionsProperty{
//   		X12: &X12AdvancedOptionsProperty{
//   			SplitOptions: &X12SplitOptionsProperty{
//   				SplitBy: jsii.String("splitBy"),
//   			},
//   		},
//   	},
//   	FormatOptions: &FormatOptionsProperty{
//   		X12: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html
//
type CfnTransformer_InputConversionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html#cfn-b2bi-transformer-inputconversion-fromformat
	//
	FromFormat *string `field:"required" json:"fromFormat" yaml:"fromFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html#cfn-b2bi-transformer-inputconversion-advancedoptions
	//
	AdvancedOptions interface{} `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-inputconversion.html#cfn-b2bi-transformer-inputconversion-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
}

