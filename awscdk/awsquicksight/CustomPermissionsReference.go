package awsquicksight


// A reference to a CustomPermissions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customPermissionsReference := &CustomPermissionsReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	CustomPermissionsArn: jsii.String("customPermissionsArn"),
//   	CustomPermissionsName: jsii.String("customPermissionsName"),
//   }
//
type CustomPermissionsReference struct {
	// The AwsAccountId of the CustomPermissions resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the CustomPermissions resource.
	CustomPermissionsArn *string `field:"required" json:"customPermissionsArn" yaml:"customPermissionsArn"`
	// The CustomPermissionsName of the CustomPermissions resource.
	CustomPermissionsName *string `field:"required" json:"customPermissionsName" yaml:"customPermissionsName"`
}

