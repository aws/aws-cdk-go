package awsemr


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoTerminationPolicyProperty := &AutoTerminationPolicyProperty{
//   	IdleTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoterminationpolicy.html
//
type CfnCluster_AutoTerminationPolicyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoterminationpolicy.html#cfn-emr-cluster-autoterminationpolicy-idletimeout
	//
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

