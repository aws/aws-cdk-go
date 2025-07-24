package awspinpoint


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDimensionProperty := &AttributeDimensionProperty{
//   	AttributeType: jsii.String("attributeType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-attributedimension.html
//
type CfnSegment_AttributeDimensionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-attributedimension.html#cfn-pinpoint-segment-attributedimension-attributetype
	//
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-attributedimension.html#cfn-pinpoint-segment-attributedimension-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

