package awspcs


// The accounting configuration includes configurable settings for Slurm accounting.
//
// It's a property of the `ClusterSlurmConfiguration` object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountingProperty := &AccountingProperty{
//   	Mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	DefaultPurgeTimeInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-accounting.html
//
type CfnCluster_AccountingProperty struct {
	// The default value for `mode` is `STANDARD` .
	//
	// A value of `STANDARD` means Slurm accounting is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-accounting.html#cfn-pcs-cluster-accounting-mode
	//
	// Default: - "NONE".
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The default value for all purge settings for `slurmdbd.conf` . For more information, see the [slurmdbd.conf documentation at SchedMD](https://docs.aws.amazon.com/https://slurm.schedmd.com/slurmdbd.conf.html) .
	//
	// The default value `-1` means there is no purge time and records persist as long as the cluster exists.
	//
	// > `0` isn't a valid value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-accounting.html#cfn-pcs-cluster-accounting-defaultpurgetimeindays
	//
	// Default: - -1.
	//
	DefaultPurgeTimeInDays *float64 `field:"optional" json:"defaultPurgeTimeInDays" yaml:"defaultPurgeTimeInDays"`
}

