package previewawsb2bimixins


// A structure that contains the X12 transaction set and version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   formatOptionsProperty := &FormatOptionsProperty{
//   	X12: &X12DetailsProperty{
//   		TransactionSet: jsii.String("transactionSet"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-formatoptions.html
//
type CfnTransformerPropsMixin_FormatOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-formatoptions.html#cfn-b2bi-transformer-formatoptions-x12
	//
	X12 interface{} `field:"optional" json:"x12" yaml:"x12"`
}

