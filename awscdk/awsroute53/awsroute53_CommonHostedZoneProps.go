package awsroute53


// Common properties to create a Route 53 hosted zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonHostedZoneProps := &commonHostedZoneProps{
//   	zoneName: jsii.String("zoneName"),
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   	queryLogsLogGroupArn: jsii.String("queryLogsLogGroupArn"),
//   }
//
type CommonHostedZoneProps struct {
	// The name of the domain.
	//
	// For resource record types that include a domain
	// name, specify a fully qualified domain name.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
	// Any comments that you want to include about the hosted zone.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The Amazon Resource Name (ARN) for the log group that you want Amazon Route 53 to send query logs to.
	QueryLogsLogGroupArn *string `field:"optional" json:"queryLogsLogGroupArn" yaml:"queryLogsLogGroupArn"`
}

