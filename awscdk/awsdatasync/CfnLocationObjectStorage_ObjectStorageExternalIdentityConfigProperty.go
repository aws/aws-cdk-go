package awsdatasync


// Specifies the external (non-AWS) identity provider that DataSync federates with to access your object storage location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectStorageExternalIdentityConfigProperty := &ObjectStorageExternalIdentityConfigProperty{
//   	GoogleOidc: &GoogleOidcConfigProperty{
//   		IdentityPoolName: jsii.String("identityPoolName"),
//   		IdentityProviderName: jsii.String("identityProviderName"),
//   		ProjectName: jsii.String("projectName"),
//   		ProjectNumber: jsii.String("projectNumber"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-objectstorageexternalidentityconfig.html
//
type CfnLocationObjectStorage_ObjectStorageExternalIdentityConfigProperty struct {
	// Specifies the Google Cloud workload identity federation configuration that DataSync uses to obtain an access token for your Google Cloud Storage bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-objectstorageexternalidentityconfig.html#cfn-datasync-locationobjectstorage-objectstorageexternalidentityconfig-googleoidc
	//
	GoogleOidc interface{} `field:"optional" json:"googleOidc" yaml:"googleOidc"`
}

