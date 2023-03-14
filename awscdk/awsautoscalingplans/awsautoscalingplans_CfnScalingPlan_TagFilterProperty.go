package awsautoscalingplans


// `TagFilter` is a subproperty of [ApplicationSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html) that specifies a tag for an application source to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &tagFilterProperty{
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnScalingPlan_TagFilterProperty struct {
	// The tag key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag values (0 to 20).
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

