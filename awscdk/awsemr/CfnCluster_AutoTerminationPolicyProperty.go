package awsemr


// An auto-termination policy for an Amazon EMR cluster.
//
// An auto-termination policy defines the amount of idle time in seconds after which a cluster automatically terminates. For alternative cluster termination options, see [Control cluster termination](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoTerminationPolicyProperty := &AutoTerminationPolicyProperty{
//   	IdleTimeout: jsii.Number(123),
//   }
//
type CfnCluster_AutoTerminationPolicyProperty struct {
	// Specifies the amount of idle time in seconds after which the cluster automatically terminates.
	//
	// You can specify a minimum of 60 seconds and a maximum of 604800 seconds (seven days).
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

