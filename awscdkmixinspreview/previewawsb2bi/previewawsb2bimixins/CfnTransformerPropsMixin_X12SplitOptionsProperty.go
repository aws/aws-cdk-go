package previewawsb2bimixins


// Contains options for splitting X12 EDI files into smaller units.
//
// This is useful for processing large EDI files more efficiently.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12SplitOptionsProperty := &X12SplitOptionsProperty{
//   	SplitBy: jsii.String("splitBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12splitoptions.html
//
type CfnTransformerPropsMixin_X12SplitOptionsProperty struct {
	// Specifies the method used to split X12 EDI files.
	//
	// Valid values include `TRANSACTION` (split by individual transaction sets), or `NONE` (no splitting).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-x12splitoptions.html#cfn-b2bi-transformer-x12splitoptions-splitby
	//
	SplitBy *string `field:"optional" json:"splitBy" yaml:"splitBy"`
}

