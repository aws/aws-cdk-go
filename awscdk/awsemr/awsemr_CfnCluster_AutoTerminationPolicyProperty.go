package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoTerminationPolicyProperty := &autoTerminationPolicyProperty{
//   	idleTimeout: jsii.Number(123),
//   }
//
type CfnCluster_AutoTerminationPolicyProperty struct {
	// `CfnCluster.AutoTerminationPolicyProperty.IdleTimeout`.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

