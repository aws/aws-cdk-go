package previewawsevidentlymixins


// A structure that defines when each step of the launch is to start, and how much launch traffic is to be allocated to each variation during each step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stepConfigProperty := &StepConfigProperty{
//   	GroupWeights: []interface{}{
//   		&GroupToWeightProperty{
//   			GroupName: jsii.String("groupName"),
//   			SplitWeight: jsii.Number(123),
//   		},
//   	},
//   	SegmentOverrides: []interface{}{
//   		&SegmentOverrideProperty{
//   			EvaluationOrder: jsii.Number(123),
//   			Segment: jsii.String("segment"),
//   			Weights: []interface{}{
//   				&GroupToWeightProperty{
//   					GroupName: jsii.String("groupName"),
//   					SplitWeight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-stepconfig.html
//
type CfnLaunchPropsMixin_StepConfigProperty struct {
	// An array of structures that define how much launch traffic to allocate to each launch group during this step of the launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-stepconfig.html#cfn-evidently-launch-stepconfig-groupweights
	//
	GroupWeights interface{} `field:"optional" json:"groupWeights" yaml:"groupWeights"`
	// An array of structures that you can use to specify different traffic splits for one or more audience *segments* .
	//
	// A segment is a portion of your audience that share one or more characteristics. Examples could be Chrome browser users, users in Europe, or Firefox browser users in Europe who also fit other criteria that your application collects, such as age.
	//
	// For more information, see [Use segments to focus your audience](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-stepconfig.html#cfn-evidently-launch-stepconfig-segmentoverrides
	//
	SegmentOverrides interface{} `field:"optional" json:"segmentOverrides" yaml:"segmentOverrides"`
	// The date and time to start this step of the launch.
	//
	// Use UTC format, `yyyy-MM-ddTHH:mm:ssZ` . For example, `2025-11-25T23:59:59Z`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-stepconfig.html#cfn-evidently-launch-stepconfig-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

