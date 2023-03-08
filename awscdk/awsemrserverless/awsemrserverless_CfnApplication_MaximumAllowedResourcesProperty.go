package awsemrserverless


// The maximum allowed cumulative resources for an application.
//
// No new resources will be created once the limit is hit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maximumAllowedResourcesProperty := &maximumAllowedResourcesProperty{
//   	cpu: jsii.String("cpu"),
//   	memory: jsii.String("memory"),
//
//   	// the properties below are optional
//   	disk: jsii.String("disk"),
//   }
//
type CfnApplication_MaximumAllowedResourcesProperty struct {
	// The maximum allowed CPU for an application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$`
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// The maximum allowed resources for an application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$`
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// The maximum allowed disk for an application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"`
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

