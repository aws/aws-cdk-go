package interfacesawsconnect


// A reference to a UseCase resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   useCaseReference := &UseCaseReference{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationAssociationId: jsii.String("integrationAssociationId"),
//   	UseCaseArn: jsii.String("useCaseArn"),
//   	UseCaseId: jsii.String("useCaseId"),
//   }
//
type UseCaseReference struct {
	// The InstanceId of the UseCase resource.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The IntegrationAssociationId of the UseCase resource.
	IntegrationAssociationId *string `field:"required" json:"integrationAssociationId" yaml:"integrationAssociationId"`
	// The ARN of the UseCase resource.
	UseCaseArn *string `field:"required" json:"useCaseArn" yaml:"useCaseArn"`
	// The UseCaseId of the UseCase resource.
	UseCaseId *string `field:"required" json:"useCaseId" yaml:"useCaseId"`
}

