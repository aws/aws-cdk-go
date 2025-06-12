package awstransfer


// A structure that describes the values to use for the IAM Identity Center settings when you create or update a web app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderDetailsProperty := &IdentityProviderDetailsProperty{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-identityproviderdetails.html
//
type CfnWebApp_IdentityProviderDetailsProperty struct {
	// The Amazon Resource Name (ARN) for the IAM Identity Center application: this value is set automatically when you create your web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-identityproviderdetails.html#cfn-transfer-webapp-identityproviderdetails-applicationarn
	//
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The Amazon Resource Name (ARN) for the IAM Identity Center used for the web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-identityproviderdetails.html#cfn-transfer-webapp-identityproviderdetails-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// The IAM role in IAM Identity Center used for the web app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-identityproviderdetails.html#cfn-transfer-webapp-identityproviderdetails-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
}

