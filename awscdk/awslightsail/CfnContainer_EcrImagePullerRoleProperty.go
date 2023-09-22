package awslightsail


// An object to describe a request to activate or deactivate the role that you can use to grant an Amazon Lightsail container service access to Amazon Elastic Container Registry (Amazon ECR) private repositories.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecrImagePullerRoleProperty := &EcrImagePullerRoleProperty{
//   	IsActive: jsii.Boolean(false),
//   	PrincipalArn: jsii.String("principalArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html
//
type CfnContainer_EcrImagePullerRoleProperty struct {
	// A Boolean value that indicates whether to activate the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html#cfn-lightsail-container-ecrimagepullerrole-isactive
	//
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The Amazon Resource Name (ARN) of the role, if it is activated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html#cfn-lightsail-container-ecrimagepullerrole-principalarn
	//
	PrincipalArn *string `field:"optional" json:"principalArn" yaml:"principalArn"`
}

