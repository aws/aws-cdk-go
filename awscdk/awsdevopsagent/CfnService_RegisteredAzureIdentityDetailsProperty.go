package awsdevopsagent


// Azure Identity service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredAzureIdentityDetailsProperty := &RegisteredAzureIdentityDetailsProperty{
//   	ClientId: jsii.String("clientId"),
//   	TenantId: jsii.String("tenantId"),
//   	WebIdentityRoleArn: jsii.String("webIdentityRoleArn"),
//   	WebIdentityTokenAudiences: []*string{
//   		jsii.String("webIdentityTokenAudiences"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredazureidentitydetails.html
//
type CfnService_RegisteredAzureIdentityDetailsProperty struct {
	// Azure AD application client ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredazureidentitydetails.html#cfn-devopsagent-service-registeredazureidentitydetails-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Azure AD tenant ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredazureidentitydetails.html#cfn-devopsagent-service-registeredazureidentitydetails-tenantid
	//
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// ARN of the IAM role for web identity token exchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredazureidentitydetails.html#cfn-devopsagent-service-registeredazureidentitydetails-webidentityrolearn
	//
	WebIdentityRoleArn *string `field:"required" json:"webIdentityRoleArn" yaml:"webIdentityRoleArn"`
	// List of audiences for the web identity token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredazureidentitydetails.html#cfn-devopsagent-service-registeredazureidentitydetails-webidentitytokenaudiences
	//
	WebIdentityTokenAudiences *[]*string `field:"required" json:"webIdentityTokenAudiences" yaml:"webIdentityTokenAudiences"`
}

