package awslogs


// A reference to a Integration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationReference := &IntegrationReference{
//   	IntegrationName: jsii.String("integrationName"),
//   }
//
type IntegrationReference struct {
	// The IntegrationName of the Integration resource.
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
}

