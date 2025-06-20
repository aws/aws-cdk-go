package awsautoscalingplans


// `TagFilter` is a subproperty of [ApplicationSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html) that specifies a tag for an application source to use with a scaling plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &TagFilterProperty{
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-tagfilter.html
//
type CfnScalingPlan_TagFilterProperty struct {
	// The tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-tagfilter.html#cfn-autoscalingplans-scalingplan-tagfilter-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag values (0 to 20).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-tagfilter.html#cfn-autoscalingplans-scalingplan-tagfilter-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

