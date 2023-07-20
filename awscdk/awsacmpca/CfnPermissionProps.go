package awsacmpca


// Properties for defining a `CfnPermission`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionProps := &CfnPermissionProps{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	Principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	SourceAccount: jsii.String("sourceAccount"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-permission.html
//
type CfnPermissionProps struct {
	// The private CA actions that can be performed by the designated AWS service.
	//
	// Supported actions are `IssueCertificate` , `GetCertificate` , and `ListPermissions` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-permission.html#cfn-acmpca-permission-actions
	//
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The Amazon Resource Number (ARN) of the private CA from which the permission was issued.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-permission.html#cfn-acmpca-permission-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The AWS service or entity that holds the permission.
	//
	// At this time, the only valid principal is `acm.amazonaws.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-permission.html#cfn-acmpca-permission-principal
	//
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The ID of the account that assigned the permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-permission.html#cfn-acmpca-permission-sourceaccount
	//
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
}

