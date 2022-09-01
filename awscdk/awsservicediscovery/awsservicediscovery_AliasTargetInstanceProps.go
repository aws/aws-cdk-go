package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var service service
//
//   aliasTargetInstanceProps := &aliasTargetInstanceProps{
//   	dnsName: jsii.String("dnsName"),
//   	service: service,
//
//   	// the properties below are optional
//   	customAttributes: map[string]*string{
//   		"customAttributesKey": jsii.String("customAttributes"),
//   	},
//   	instanceId: jsii.String("instanceId"),
//   }
//
type AliasTargetInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// DNS name of the target.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The Cloudmap service this resource is registered to.
	Service IService `field:"required" json:"service" yaml:"service"`
}

