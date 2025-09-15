package awsappintegrations


// A reference to a EventIntegration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventIntegrationReference := &EventIntegrationReference{
//   	EventIntegrationArn: jsii.String("eventIntegrationArn"),
//   	EventIntegrationName: jsii.String("eventIntegrationName"),
//   }
//
type EventIntegrationReference struct {
	// The ARN of the EventIntegration resource.
	EventIntegrationArn *string `field:"required" json:"eventIntegrationArn" yaml:"eventIntegrationArn"`
	// The Name of the EventIntegration resource.
	EventIntegrationName *string `field:"required" json:"eventIntegrationName" yaml:"eventIntegrationName"`
}

