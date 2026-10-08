package awsmediatailor


// Clip range configuration for the VOD source associated with the program.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clipRangeProperty := &ClipRangeProperty{
//   	EndOffsetMillis: jsii.Number(123),
//   	StartOffsetMillis: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-cliprange.html
//
type CfnProgram_ClipRangeProperty struct {
	// The end offset of the clip range, in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-cliprange.html#cfn-mediatailor-program-cliprange-endoffsetmillis
	//
	EndOffsetMillis *float64 `field:"optional" json:"endOffsetMillis" yaml:"endOffsetMillis"`
	// The start offset of the clip range, in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-cliprange.html#cfn-mediatailor-program-cliprange-startoffsetmillis
	//
	StartOffsetMillis *float64 `field:"optional" json:"startOffsetMillis" yaml:"startOffsetMillis"`
}

