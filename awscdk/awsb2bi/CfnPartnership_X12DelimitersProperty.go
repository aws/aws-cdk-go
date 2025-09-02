package awsb2bi


// In X12 EDI messages, delimiters are used to mark the end of segments or elements, and are defined in the interchange control header.
//
// The delimiters are part of the message's syntax and divide up its different elements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12DelimitersProperty := &X12DelimitersProperty{
//   	ComponentSeparator: jsii.String("componentSeparator"),
//   	DataElementSeparator: jsii.String("dataElementSeparator"),
//   	SegmentTerminator: jsii.String("segmentTerminator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html
//
type CfnPartnership_X12DelimitersProperty struct {
	// The component, or sub-element, separator.
	//
	// The default value is `:` (colon).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html#cfn-b2bi-partnership-x12delimiters-componentseparator
	//
	ComponentSeparator *string `field:"optional" json:"componentSeparator" yaml:"componentSeparator"`
	// The data element separator.
	//
	// The default value is `*` (asterisk).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html#cfn-b2bi-partnership-x12delimiters-dataelementseparator
	//
	DataElementSeparator *string `field:"optional" json:"dataElementSeparator" yaml:"dataElementSeparator"`
	// The segment terminator.
	//
	// The default value is `~` (tilde).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html#cfn-b2bi-partnership-x12delimiters-segmentterminator
	//
	SegmentTerminator *string `field:"optional" json:"segmentTerminator" yaml:"segmentTerminator"`
}

