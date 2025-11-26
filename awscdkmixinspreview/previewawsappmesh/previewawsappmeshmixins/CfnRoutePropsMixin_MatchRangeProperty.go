package previewawsappmeshmixins


// An object that represents the range of values to match on.
//
// The first character of the range is included in the range, though the last character is not. For example, if the range specified were 1-100, only values 1-99 would be matched.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   matchRangeProperty := &MatchRangeProperty{
//   	End: jsii.Number(123),
//   	Start: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-matchrange.html
//
type CfnRoutePropsMixin_MatchRangeProperty struct {
	// The end of the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-matchrange.html#cfn-appmesh-route-matchrange-end
	//
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// The start of the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-route-matchrange.html#cfn-appmesh-route-matchrange-start
	//
	Start *float64 `field:"optional" json:"start" yaml:"start"`
}

