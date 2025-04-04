package awsb2bi


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html#cfn-b2bi-partnership-x12delimiters-componentseparator
	//
	ComponentSeparator *string `field:"optional" json:"componentSeparator" yaml:"componentSeparator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html#cfn-b2bi-partnership-x12delimiters-dataelementseparator
	//
	DataElementSeparator *string `field:"optional" json:"dataElementSeparator" yaml:"dataElementSeparator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12delimiters.html#cfn-b2bi-partnership-x12delimiters-segmentterminator
	//
	SegmentTerminator *string `field:"optional" json:"segmentTerminator" yaml:"segmentTerminator"`
}

