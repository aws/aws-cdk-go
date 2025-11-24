package mixinsawsemrserverless


// The configuration of a worker.
//
// For more information, see [Supported worker configurations](https://docs.aws.amazon.com/emr/latest/EMR-Serverless-UserGuide/app-behavior.html#worker-configs) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workerConfigurationProperty := &WorkerConfigurationProperty{
//   	Cpu: jsii.String("cpu"),
//   	Disk: jsii.String("disk"),
//   	DiskType: jsii.String("diskType"),
//   	Memory: jsii.String("memory"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html
//
type CfnApplicationPropsMixin_WorkerConfigurationProperty struct {
	// The CPU requirements of the worker configuration.
	//
	// Each worker can have 1, 2, 4, 8, or 16 vCPUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-cpu
	//
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The disk requirements of the worker configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-disk
	//
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
	// The disk type for every worker instance of the work type.
	//
	// Shuffle optimized disks have higher performance characteristics and are better for shuffle heavy workloads. Default is `STANDARD` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-disktype
	//
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// The memory requirements of the worker configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-memory
	//
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

