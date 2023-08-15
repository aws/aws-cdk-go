package awscdk


// Options for specifying a role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   roleOptions := &RoleOptions{
//   	AssumeRoleArn: jsii.String("assumeRoleArn"),
//
//   	// the properties below are optional
//   	AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   }
//
type RoleOptions struct {
	// ARN of the role to assume.
	AssumeRoleArn *string `field:"required" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// External ID to use when assuming the role.
	// Default: - No external ID.
	//
	AssumeRoleExternalId *string `field:"optional" json:"assumeRoleExternalId" yaml:"assumeRoleExternalId"`
}

