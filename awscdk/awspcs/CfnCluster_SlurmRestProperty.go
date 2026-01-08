package awspcs


// The Slurm REST API configuration includes settings for enabling and configuring the Slurm REST API.
//
// It's a property of the [ClusterSlurmConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmconfiguration.html) object.
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
	// The default value for `mode` is `NONE` .
	//
	// A value of `STANDARD` means the Slurm REST API is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-slurmrest.html#cfn-pcs-cluster-slurmrest-mode
	//
	// Default: - "NONE".
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

