package awsroute53


// A CloudWatch alarm that you want Amazon Route 53 health checker to use to determine whether this health check is healthy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmIdentifier := &AlarmIdentifier{
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   }
//
type AlarmIdentifier struct {
	// The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The region of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.
	Region *string `field:"required" json:"region" yaml:"region"`
}

