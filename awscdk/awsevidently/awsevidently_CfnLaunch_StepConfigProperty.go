package awsevidently


// A structure that defines when each step of the launch is to start, and how much launch traffic is to be allocated to each variation during each step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stepConfigProperty := &stepConfigProperty{
//   	groupWeights: []interface{}{
//   		&groupToWeightProperty{
//   			groupName: jsii.String("groupName"),
//   			splitWeight: jsii.Number(123),
//   		},
//   	},
//   	startTime: jsii.String("startTime"),
//
//   	// the properties below are optional
//   	segmentOverrides: []interface{}{
//   		&segmentOverrideProperty{
//   			evaluationOrder: jsii.Number(123),
//   			segment: jsii.String("segment"),
//   			weights: []interface{}{
//   				&groupToWeightProperty{
//   					groupName: jsii.String("groupName"),
//   					splitWeight: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnLaunch_StepConfigProperty struct {
	// An array of structures that define how much launch traffic to allocate to each launch group during this step of the launch.
	GroupWeights interface{} `field:"required" json:"groupWeights" yaml:"groupWeights"`
	// The date and time to start this step of the launch.
	//
	// Use UTC format, `yyyy-MM-ddTHH:mm:ssZ` . For example, `2025-11-25T23:59:59Z`
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// `CfnLaunch.StepConfigProperty.SegmentOverrides`.
	SegmentOverrides interface{} `field:"optional" json:"segmentOverrides" yaml:"segmentOverrides"`
}

