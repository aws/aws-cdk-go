package awsemrserverless


// The resource configuration of the initial capacity configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workerConfigurationProperty := &workerConfigurationProperty{
//   	cpu: jsii.String("cpu"),
//   	memory: jsii.String("memory"),
//
//   	// the properties below are optional
//   	disk: jsii.String("disk"),
//   }
//
type CfnApplication_WorkerConfigurationProperty struct {
	// *Minimum* : 1.
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$`
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// *Minimum* : 1.
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$`
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// *Minimum* : 1.
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"`
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

