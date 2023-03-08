package awsevents


// The primary Region of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryProperty := &primaryProperty{
//   	healthCheck: jsii.String("healthCheck"),
//   }
//
type CfnEndpoint_PrimaryProperty struct {
	// The ARN of the health check used by the endpoint to determine whether failover is triggered.
	HealthCheck *string `field:"required" json:"healthCheck" yaml:"healthCheck"`
}

