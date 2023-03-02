package awsecs


// The `HostEntry` property specifies a hostname and an IP address that are added to the `/etc/hosts` file of a container through the `extraHosts` parameter of its `ContainerDefinition` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostEntryProperty := &hostEntryProperty{
//   	hostname: jsii.String("hostname"),
//   	ipAddress: jsii.String("ipAddress"),
//   }
//
type CfnTaskDefinition_HostEntryProperty struct {
	// The hostname to use in the `/etc/hosts` entry.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The IP address to use in the `/etc/hosts` entry.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

