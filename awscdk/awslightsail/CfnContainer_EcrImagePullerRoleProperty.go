package awslightsail


// Describes the IAM role that you can use to grant a Lightsail container service access to Amazon ECR private repositories.
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
	// A boolean value that indicates whether the `ECRImagePullerRole` is active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html#cfn-lightsail-container-ecrimagepullerrole-isactive
	//
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The principle Amazon Resource Name (ARN) of the role.
	//
	// This property is read-only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html#cfn-lightsail-container-ecrimagepullerrole-principalarn
	//
	PrincipalArn *string `field:"optional" json:"principalArn" yaml:"principalArn"`
}

