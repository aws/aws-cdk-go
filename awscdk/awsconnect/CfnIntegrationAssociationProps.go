package awsconnect


// Properties for defining a `CfnIntegrationAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationAssociationProps := &CfnIntegrationAssociationProps{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationArn: jsii.String("integrationArn"),
//   	IntegrationType: jsii.String("integrationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-integrationassociation.html
//
type CfnIntegrationAssociationProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-integrationassociation.html#cfn-connect-integrationassociation-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// ARN of the integration being associated with the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `140`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-integrationassociation.html#cfn-connect-integrationassociation-integrationarn
	//
	IntegrationArn *string `field:"required" json:"integrationArn" yaml:"integrationArn"`
	// Specifies the integration type to be associated with the instance.
	//
	// *Allowed Values* : `LEX_BOT` | `LAMBDA_FUNCTION`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-integrationassociation.html#cfn-connect-integrationassociation-integrationtype
	//
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
}

