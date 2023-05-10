package cloudassemblyschema


// Query to endpoint service context provider.
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
//
//   	// the properties below are optional
//   	LookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	// Experimental.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Query service name.
	// Experimental.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}

