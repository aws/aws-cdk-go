package previewawspinpointmixins


// Specifies behavior-based criteria for the segment, such as how recently users have used your app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   behaviorProperty := &BehaviorProperty{
//   	Recency: &RecencyProperty{
//   		Duration: jsii.String("duration"),
//   		RecencyType: jsii.String("recencyType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-behavior.html
//
type CfnSegmentPropsMixin_BehaviorProperty struct {
	// Specifies how recently segment members were active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-behavior.html#cfn-pinpoint-segment-behavior-recency
	//
	Recency interface{} `field:"optional" json:"recency" yaml:"recency"`
}

