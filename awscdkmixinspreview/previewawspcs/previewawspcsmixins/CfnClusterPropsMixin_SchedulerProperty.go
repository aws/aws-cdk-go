package previewawspcsmixins


// The cluster management and job scheduling software associated with the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schedulerProperty := &SchedulerProperty{
//   	Type: jsii.String("type"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-scheduler.html
//
type CfnClusterPropsMixin_SchedulerProperty struct {
	// The software AWS PCS uses to manage cluster scaling and job scheduling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-scheduler.html#cfn-pcs-cluster-scheduler-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The version of the specified scheduling software that AWS PCS uses to manage cluster scaling and job scheduling.
	//
	// For more information, see [Slurm versions in AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/slurm-versions.html) in the *AWS PCS User Guide* .
	//
	// Valid Values: `23.11 | 24.05 | 24.11`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-scheduler.html#cfn-pcs-cluster-scheduler-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

