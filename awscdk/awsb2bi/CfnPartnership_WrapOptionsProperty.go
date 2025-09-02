package awsb2bi


// Contains options for wrapping (line folding) in X12 EDI files.
//
// Wrapping controls how long lines are handled in the EDI output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wrapOptionsProperty := &WrapOptionsProperty{
//   	LineLength: jsii.Number(123),
//   	LineTerminator: jsii.String("lineTerminator"),
//   	WrapBy: jsii.String("wrapBy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html
//
type CfnPartnership_WrapOptionsProperty struct {
	// Specifies the maximum length of a line before wrapping occurs.
	//
	// This value is used when `wrapBy` is set to `LINE_LENGTH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html#cfn-b2bi-partnership-wrapoptions-linelength
	//
	LineLength *float64 `field:"optional" json:"lineLength" yaml:"lineLength"`
	// Specifies the character sequence used to terminate lines when wrapping. Valid values:.
	//
	// - `CRLF` : carriage return and line feed
	// - `LF` : line feed)
	// - `CR` : carriage return.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html#cfn-b2bi-partnership-wrapoptions-lineterminator
	//
	LineTerminator *string `field:"optional" json:"lineTerminator" yaml:"lineTerminator"`
	// Specifies the method used for wrapping lines in the EDI output. Valid values:.
	//
	// - `SEGMENT` : Wraps by segment.
	// - `ONE_LINE` : Indicates that the entire content is on a single line.
	//
	// > When you specify `ONE_LINE` , do not provide either the line length nor the line terminator value.
	// - `LINE_LENGTH` : Wraps by character count, as specified by `lineLength` value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-wrapoptions.html#cfn-b2bi-partnership-wrapoptions-wrapby
	//
	WrapBy *string `field:"optional" json:"wrapBy" yaml:"wrapBy"`
}

