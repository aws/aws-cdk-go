package awsemrserverless


// The initial capacity configuration per worker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initialCapacityConfigProperty := &InitialCapacityConfigProperty{
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		Cpu: jsii.String("cpu"),
//   		Memory: jsii.String("memory"),
//
//   		// the properties below are optional
//   		Disk: jsii.String("disk"),
//   		DiskType: jsii.String("diskType"),
//   	},
//   	WorkerCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfig.html
//
type CfnApplication_InitialCapacityConfigProperty struct {
	// The resource configuration of the initial capacity configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfig.html#cfn-emrserverless-application-initialcapacityconfig-workerconfiguration
	//
	WorkerConfiguration interface{} `field:"required" json:"workerConfiguration" yaml:"workerConfiguration"`
	// The number of workers in the initial capacity configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-initialcapacityconfig.html#cfn-emrserverless-application-initialcapacityconfig-workercount
	//
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
}

