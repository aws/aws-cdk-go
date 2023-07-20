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
//   maximumAllowedResourcesProperty := &MaximumAllowedResourcesProperty{
//   	Cpu: jsii.String("cpu"),
//   	Memory: jsii.String("memory"),
//
//   	// the properties below are optional
//   	Disk: jsii.String("disk"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-maximumallowedresources.html
//
type CfnApplication_MaximumAllowedResourcesProperty struct {
	// The maximum allowed CPU for an application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-maximumallowedresources.html#cfn-emrserverless-application-maximumallowedresources-cpu
	//
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// The maximum allowed resources for an application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-maximumallowedresources.html#cfn-emrserverless-application-maximumallowedresources-memory
	//
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// The maximum allowed disk for an application.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 15
	//
	// *Pattern* : `^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-maximumallowedresources.html#cfn-emrserverless-application-maximumallowedresources-disk
	//
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

