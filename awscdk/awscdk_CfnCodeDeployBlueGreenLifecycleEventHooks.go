// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Lifecycle events for blue-green deployments.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.lifecycleEventHooks} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenLifecycleEventHooks := &cfnCodeDeployBlueGreenLifecycleEventHooks{
//   	afterAllowTestTraffic: jsii.String("afterAllowTestTraffic"),
//   	afterAllowTraffic: jsii.String("afterAllowTraffic"),
//   	afterInstall: jsii.String("afterInstall"),
//   	beforeAllowTraffic: jsii.String("beforeAllowTraffic"),
//   	beforeInstall: jsii.String("beforeInstall"),
//   }
//
type CfnCodeDeployBlueGreenLifecycleEventHooks struct {
	// Function to use to run tasks after the test listener serves traffic to the replacement task set.
	AfterAllowTestTraffic *string `field:"optional" json:"afterAllowTestTraffic" yaml:"afterAllowTestTraffic"`
	// Function to use to run tasks after the second target group serves traffic to the replacement task set.
	AfterAllowTraffic *string `field:"optional" json:"afterAllowTraffic" yaml:"afterAllowTraffic"`
	// Function to use to run tasks after the replacement task set is created and one of the target groups is associated with it.
	AfterInstall *string `field:"optional" json:"afterInstall" yaml:"afterInstall"`
	// Function to use to run tasks after the second target group is associated with the replacement task set, but before traffic is shifted to the replacement task set.
	BeforeAllowTraffic *string `field:"optional" json:"beforeAllowTraffic" yaml:"beforeAllowTraffic"`
	// Function to use to run tasks before the replacement task set is created.
	BeforeInstall *string `field:"optional" json:"beforeInstall" yaml:"beforeInstall"`
}

