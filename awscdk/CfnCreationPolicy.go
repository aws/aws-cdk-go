package awscdk


// Associate the CreationPolicy attribute with a resource to prevent its status from reaching create complete until AWS CloudFormation receives a specified number of success signals or the timeout period is exceeded.
//
// To signal a
// resource, you can use the cfn-signal helper script or SignalResource API. AWS CloudFormation publishes valid signals
// to the stack events so that you track the number of signals sent.
//
// The creation policy is invoked only when AWS CloudFormation creates the associated resource. Currently, the only
// AWS CloudFormation resources that support creation policies are AWS::AutoScaling::AutoScalingGroup, AWS::EC2::Instance,
// AWS::CloudFormation::WaitCondition and AWS::AppStream::Fleet.
//
// Use the CreationPolicy attribute when you want to wait on resource configuration actions before stack creation proceeds.
// For example, if you install and configure software applications on an EC2 instance, you might want those applications to
// be running before proceeding. In such cases, you can add a CreationPolicy attribute to the instance, and then send a success
// signal to the instance after the applications are installed and configured. For a detailed example, see Deploying Applications
// on Amazon EC2 with AWS CloudFormation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCreationPolicy := &CfnCreationPolicy{
//   	AutoScalingCreationPolicy: &CfnResourceAutoScalingCreationPolicy{
//   		MinSuccessfulInstancesPercent: jsii.Number(123),
//   	},
//   	ResourceSignal: &CfnResourceSignal{
//   		Count: jsii.Number(123),
//   		Timeout: jsii.String("timeout"),
//   	},
//   	StartFleet: jsii.Boolean(false),
//   }
//
type CfnCreationPolicy struct {
	// For an Auto Scaling group replacement update, specifies how many instances must signal success for the update to succeed.
	AutoScalingCreationPolicy *CfnResourceAutoScalingCreationPolicy `field:"optional" json:"autoScalingCreationPolicy" yaml:"autoScalingCreationPolicy"`
	// When AWS CloudFormation creates the associated resource, configures the number of required success signals and the length of time that AWS CloudFormation waits for those signals.
	ResourceSignal *CfnResourceSignal `field:"optional" json:"resourceSignal" yaml:"resourceSignal"`
	// For an AppStream Fleet creation, specifies that the fleet is started after creation.
	StartFleet *bool `field:"optional" json:"startFleet" yaml:"startFleet"`
}

