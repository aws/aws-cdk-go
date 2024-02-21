package awslightsail


// Describes the configuration for an Amazon Lightsail container service to access private container image repositories, such as Amazon Elastic Container Registry ( Amazon ECR ) private repositories.
//
// For more information, see [Configuring access to an Amazon ECR private repository for an Amazon Lightsail container service](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-container-service-ecr-private-repo-access) in the *Amazon Lightsail Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateRegistryAccessProperty := &PrivateRegistryAccessProperty{
//   	EcrImagePullerRole: &EcrImagePullerRoleProperty{
//   		IsActive: jsii.Boolean(false),
//   		PrincipalArn: jsii.String("principalArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-privateregistryaccess.html
//
type CfnContainer_PrivateRegistryAccessProperty struct {
	// An object that describes the activation status of the role that you can use to grant a Lightsail container service access to Amazon ECR private repositories.
	//
	// If the role is activated, the Amazon Resource Name (ARN) of the role is also listed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-privateregistryaccess.html#cfn-lightsail-container-privateregistryaccess-ecrimagepullerrole
	//
	EcrImagePullerRole interface{} `field:"optional" json:"ecrImagePullerRole" yaml:"ecrImagePullerRole"`
}

