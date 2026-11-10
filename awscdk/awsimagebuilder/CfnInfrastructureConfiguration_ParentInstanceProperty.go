package awsimagebuilder


// Settings for building on an instance that is already running, instead of an instance that Image Builder launches for the build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parentInstanceProperty := &ParentInstanceProperty{
//   	AutoScalingGroup: &AutoScalingGroupProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		CapacityWaitTimeoutSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-parentinstance.html
//
type CfnInfrastructureConfiguration_ParentInstanceProperty struct {
	// Identifies the Amazon EC2 Auto Scaling group that Image Builder takes the build instance from, and how long Image Builder waits when the group has no instance available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-parentinstance.html#cfn-imagebuilder-infrastructureconfiguration-parentinstance-autoscalinggroup
	//
	AutoScalingGroup interface{} `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

