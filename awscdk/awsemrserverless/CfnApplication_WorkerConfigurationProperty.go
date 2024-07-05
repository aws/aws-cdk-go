package awsemrserverless


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerConfigurationProperty := &WorkerConfigurationProperty{
//   	Cpu: jsii.String("cpu"),
//   	Memory: jsii.String("memory"),
//
//   	// the properties below are optional
//   	Disk: jsii.String("disk"),
//   	DiskType: jsii.String("diskType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html
//
type CfnApplication_WorkerConfigurationProperty struct {
	// Per worker CPU resource.
	//
	// vCPU is the only supported unit and specifying vCPU is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-cpu
	//
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Per worker memory resource.
	//
	// GB is the only supported unit and specifying GB is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-memory
	//
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Per worker Disk resource.
	//
	// GB is the only supported unit and specifying GB is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-disk
	//
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
	// Per worker DiskType resource.
	//
	// Shuffle optimized and Standard are only supported types and specifying diskType is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-disktype
	//
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

