package awsdatasync


// Specifies the identity federation configuration that DataSync uses to access your object storage location using an OpenID Connect (OIDC) token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   objectStorageFederatedIdentityConfigProperty := &ObjectStorageFederatedIdentityConfigProperty{
//   	AwsIamRole: jsii.String("awsIamRole"),
//   	ExternalIdentity: &ObjectStorageExternalIdentityConfigProperty{
//   		GoogleOidc: &GoogleOidcConfigProperty{
//   			IdentityPoolName: jsii.String("identityPoolName"),
//   			IdentityProviderName: jsii.String("identityProviderName"),
//   			ProjectName: jsii.String("projectName"),
//   			ProjectNumber: jsii.String("projectNumber"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-objectstoragefederatedidentityconfig.html
//
type CfnLocationObjectStoragePropsMixin_ObjectStorageFederatedIdentityConfigProperty struct {
	// Specifies the ARN of the AWS Identity and Access Management (IAM) role that DataSync assumes to mint the OIDC token used to authenticate with the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-objectstoragefederatedidentityconfig.html#cfn-datasync-locationobjectstorage-objectstoragefederatedidentityconfig-awsiamrole
	//
	AwsIamRole *string `field:"optional" json:"awsIamRole" yaml:"awsIamRole"`
	// Specifies the external (non-AWS) identity provider that DataSync federates with to access your object storage location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-objectstoragefederatedidentityconfig.html#cfn-datasync-locationobjectstorage-objectstoragefederatedidentityconfig-externalidentity
	//
	ExternalIdentity interface{} `field:"optional" json:"externalIdentity" yaml:"externalIdentity"`
}

