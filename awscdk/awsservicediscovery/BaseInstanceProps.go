package awsservicediscovery


// Used when the resource that's associated with the service instance is accessible using values other than an IP address or a domain name (CNAME), i.e. for non-ip-instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseInstanceProps := &BaseInstanceProps{
//   	CustomAttributes: map[string]*string{
//   		"customAttributesKey": jsii.String("customAttributes"),
//   	},
//   	InstanceId: jsii.String("instanceId"),
//   }
//
type BaseInstanceProps struct {
	// Custom attributes of the instance.
	// Default: none.
	//
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	// Default: Automatically generated name.
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

