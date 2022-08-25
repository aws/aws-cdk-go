package awsroute53


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmIdentifierProperty := &alarmIdentifierProperty{
//   	name: jsii.String("name"),
//   	region: jsii.String("region"),
//   }
//
type CfnHealthCheck_AlarmIdentifierProperty struct {
	// `CfnHealthCheck.AlarmIdentifierProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnHealthCheck.AlarmIdentifierProperty.Region`.
	Region *string `field:"required" json:"region" yaml:"region"`
}

