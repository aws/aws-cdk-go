package awspinpointemail


// An object that defines a Amazon Pinpoint destination for email events.
//
// You can use Amazon Pinpoint events to create attributes in Amazon Pinpoint projects. You can use these attributes to create segments for your campaigns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pinpointDestinationProperty := &pinpointDestinationProperty{
//   	applicationArn: jsii.String("applicationArn"),
//   }
//
type CfnConfigurationSetEventDestination_PinpointDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon Pinpoint project that you want to send email events to.
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
}

