package awsemr


// Managed scaling policy for an Amazon EMR cluster.
//
// The policy specifies the limits for resources that can be added or terminated from a cluster. The policy only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedScalingPolicyProperty := &managedScalingPolicyProperty{
//   	computeLimits: &computeLimitsProperty{
//   		maximumCapacityUnits: jsii.Number(123),
//   		minimumCapacityUnits: jsii.Number(123),
//   		unitType: jsii.String("unitType"),
//
//   		// the properties below are optional
//   		maximumCoreCapacityUnits: jsii.Number(123),
//   		maximumOnDemandCapacityUnits: jsii.Number(123),
//   	},
//   }
//
type CfnCluster_ManagedScalingPolicyProperty struct {
	// The EC2 unit limits for a managed scaling policy.
	//
	// The managed scaling activity of a cluster is not allowed to go above or below these limits. The limit only applies to the core and task nodes. The master node cannot be scaled after initial configuration.
	ComputeLimits interface{} `field:"optional" json:"computeLimits" yaml:"computeLimits"`
}

