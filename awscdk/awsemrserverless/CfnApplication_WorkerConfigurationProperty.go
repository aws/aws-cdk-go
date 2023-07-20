package awsemrserverless


// The resource configuration of the initial capacity configuration.
//
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html
//
type CfnApplication_WorkerConfigurationProperty struct {
	// *Minimum* : 1.
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-cpu
	//
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// *Minimum* : 1.
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-memory
	//
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// *Minimum* : 1.
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-workerconfiguration.html#cfn-emrserverless-application-workerconfiguration-disk
	//
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

