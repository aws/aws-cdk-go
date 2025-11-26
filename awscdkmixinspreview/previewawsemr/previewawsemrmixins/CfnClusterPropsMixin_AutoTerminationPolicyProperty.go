package previewawsemrmixins


// An auto-termination policy for an Amazon EMR cluster.
//
// An auto-termination policy defines the amount of idle time in seconds after which a cluster automatically terminates. For alternative cluster termination options, see [Control cluster termination](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoTerminationPolicyProperty := &AutoTerminationPolicyProperty{
//   	IdleTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoterminationpolicy.html
//
type CfnClusterPropsMixin_AutoTerminationPolicyProperty struct {
	// Specifies the amount of idle time in seconds after which the cluster automatically terminates.
	//
	// You can specify a minimum of 60 seconds and a maximum of 604800 seconds (seven days).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-autoterminationpolicy.html#cfn-emr-cluster-autoterminationpolicy-idletimeout
	//
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
}

