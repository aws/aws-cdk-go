package awspcs


// The SlurmRest configuration includes configurable settings for Slurm Rest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slurmRestProperty := &SlurmRestProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmrest.html
//
type CfnCluster_SlurmRestProperty struct {
	// The default value is `STANDARD`.
	//
	// A value of `STANDARD` means that Slurm Rest is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmrest.html#cfn-pcs-cluster-slurmrest-mode
	//
	// Default: - "NONE".
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

