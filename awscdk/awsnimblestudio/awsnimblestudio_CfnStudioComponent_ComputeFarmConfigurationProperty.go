package awsnimblestudio


// The configuration for a render farm that is associated with a studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeFarmConfigurationProperty := &computeFarmConfigurationProperty{
//   	activeDirectoryUser: jsii.String("activeDirectoryUser"),
//   	endpoint: jsii.String("endpoint"),
//   }
//
type CfnStudioComponent_ComputeFarmConfigurationProperty struct {
	// The name of an Active Directory user that is used on ComputeFarm worker instances.
	ActiveDirectoryUser *string `field:"optional" json:"activeDirectoryUser" yaml:"activeDirectoryUser"`
	// The endpoint of the ComputeFarm that is accessed by the studio component resource.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

