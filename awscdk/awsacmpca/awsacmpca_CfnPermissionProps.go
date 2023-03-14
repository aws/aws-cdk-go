package awsacmpca


// Properties for defining a `CfnPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionProps := &cfnPermissionProps{
//   	actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	sourceAccount: jsii.String("sourceAccount"),
//   }
//
type CfnPermissionProps struct {
	// The private CA actions that can be performed by the designated AWS service.
	//
	// Supported actions are `IssueCertificate` , `GetCertificate` , and `ListPermissions` .
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Number (ARN) of the private CA from which the permission was issued.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The AWS service or entity that holds the permission.
	//
	// At this time, the only valid principal is `acm.amazonaws.com` .
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The ID of the account that assigned the permission.
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
}

