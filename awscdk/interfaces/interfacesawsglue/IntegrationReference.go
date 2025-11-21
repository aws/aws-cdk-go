package interfacesawsglue


// A reference to a Integration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationReference := &IntegrationReference{
//   	IntegrationArn: jsii.String("integrationArn"),
//   	IntegrationName: jsii.String("integrationName"),
//   }
//
type IntegrationReference struct {
	// The IntegrationArn of the Integration resource.
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// The IntegrationName of the Integration resource.
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
}

