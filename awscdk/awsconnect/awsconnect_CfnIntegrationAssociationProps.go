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
	// The Amazon Resource Name (ARN) of the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// ARN of the integration being associated with the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `140`.
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// Specifies the integration type to be associated with the instance.
	//
	// *Allowed Values* : `LEX_BOT` | `LAMBDA_FUNCTION`.
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
}

