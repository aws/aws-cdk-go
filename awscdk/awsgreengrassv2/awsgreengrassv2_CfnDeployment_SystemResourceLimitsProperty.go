package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemResourceLimitsProperty := &systemResourceLimitsProperty{
//   	cpus: jsii.Number(123),
//   	memory: jsii.Number(123),
//   }
//
type CfnDeployment_SystemResourceLimitsProperty struct {
	// `CfnDeployment.SystemResourceLimitsProperty.Cpus`.
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// `CfnDeployment.SystemResourceLimitsProperty.Memory`.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
}

