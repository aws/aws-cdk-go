package cxapi


// Query to hosted zone context provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointServiceAvailabilityZonesContextQuery := &EndpointServiceAvailabilityZonesContextQuery{
//   	Account: jsii.String("account"),
//   	Region: jsii.String("region"),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Query region.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Query service name.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

