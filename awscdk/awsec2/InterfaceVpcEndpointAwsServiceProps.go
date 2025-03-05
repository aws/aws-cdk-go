package awsec2


// Optional properties for the InterfaceVpcEndpointAwsService class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   interfaceVpcEndpointAwsServiceProps := &InterfaceVpcEndpointAwsServiceProps{
//   	Global: jsii.Boolean(false),
//   }
//
type InterfaceVpcEndpointAwsServiceProps struct {
	// If true, the service is a global endpoint and its name will not be prefixed with the stack's region.
	// Default: false.
	//
	Global *bool `field:"optional" json:"global" yaml:"global"`
}

