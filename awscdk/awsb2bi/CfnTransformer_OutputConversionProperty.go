package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputConversionProperty := &OutputConversionProperty{
//   	ToFormat: jsii.String("toFormat"),
//
//   	// the properties below are optional
//   	FormatOptions: &FormatOptionsProperty{
//   		X12: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html
//
type CfnTransformer_OutputConversionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html#cfn-b2bi-transformer-outputconversion-toformat
	//
	ToFormat *string `field:"required" json:"toFormat" yaml:"toFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-outputconversion.html#cfn-b2bi-transformer-outputconversion-formatoptions
	//
	FormatOptions interface{} `field:"optional" json:"formatOptions" yaml:"formatOptions"`
}

