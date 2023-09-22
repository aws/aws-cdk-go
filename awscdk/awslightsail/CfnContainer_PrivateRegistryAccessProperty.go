package awslightsail


// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.
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
	// An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-privateregistryaccess.html#cfn-lightsail-container-privateregistryaccess-ecrimagepullerrole
	//
	EcrImagePullerRole interface{} `field:"optional" json:"ecrImagePullerRole" yaml:"ecrImagePullerRole"`
}

