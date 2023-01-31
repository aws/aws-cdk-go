package awsconnect


// Properties for defining a `CfnIntegrationAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationAssociationProps := &cfnIntegrationAssociationProps{
//   	instanceId: jsii.String("instanceId"),
//   	integrationArn: jsii.String("integrationArn"),
//   	integrationType: jsii.String("integrationType"),
//   }
//
type CfnIntegrationAssociationProps struct {
	// The identifier of the Amazon Connect instance.
	//
	// You can find the instanceId in the ARN of the instance.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The Amazon Resource Name (ARN) for the AppIntegration.
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// The integration type.
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
}

