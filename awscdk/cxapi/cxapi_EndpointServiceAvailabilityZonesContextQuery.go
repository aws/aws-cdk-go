package cxapi


// Query to hosted zone context provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointServiceAvailabilityZonesContextQuery := &endpointServiceAvailabilityZonesContextQuery{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Query region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Query service name.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

