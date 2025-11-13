package interfacesawsconnect


// A reference to a IntegrationAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationAssociationReference := &IntegrationAssociationReference{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationArn: jsii.String("integrationArn"),
//   	IntegrationType: jsii.String("integrationType"),
//   }
//
type IntegrationAssociationReference struct {
	// The InstanceId of the IntegrationAssociation resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The IntegrationArn of the IntegrationAssociation resource.
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// The IntegrationType of the IntegrationAssociation resource.
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
}

