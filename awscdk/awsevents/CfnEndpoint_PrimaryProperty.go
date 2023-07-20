package awsevents


// The primary Region of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   primaryProperty := &PrimaryProperty{
//   	HealthCheck: jsii.String("healthCheck"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-primary.html
//
type CfnEndpoint_PrimaryProperty struct {
	// The ARN of the health check used by the endpoint to determine whether failover is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-endpoint-primary.html#cfn-events-endpoint-primary-healthcheck
	//
	HealthCheck *string `field:"required" json:"healthCheck" yaml:"healthCheck"`
}

