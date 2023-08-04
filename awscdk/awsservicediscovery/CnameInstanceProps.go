package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var service service
//
//   cnameInstanceProps := &CnameInstanceProps{
//   	InstanceCname: jsii.String("instanceCname"),
//   	Service: service,
//
//   	// the properties below are optional
//   	CustomAttributes: map[string]*string{
//   		"customAttributesKey": jsii.String("customAttributes"),
//   	},
//   	InstanceId: jsii.String("instanceId"),
//   }
//
type CnameInstanceProps struct {
	// Custom attributes of the instance.
	// Default: none.
	//
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	// Default: Automatically generated name.
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// If the service configuration includes a CNAME record, the domain name that you want Route 53 to return in response to DNS queries, for example, example.com. This value is required if the service specified by ServiceId includes settings for an CNAME record.
	InstanceCname *string `field:"required" json:"instanceCname" yaml:"instanceCname"`
	// The Cloudmap service this resource is registered to.
	Service IService `field:"required" json:"service" yaml:"service"`
}

