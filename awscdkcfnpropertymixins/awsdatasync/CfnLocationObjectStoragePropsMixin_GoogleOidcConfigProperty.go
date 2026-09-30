package awsdatasync


// Specifies the Google Cloud workload identity federation configuration that DataSync uses to obtain an access token for your Google Cloud Storage bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   googleOidcConfigProperty := &GoogleOidcConfigProperty{
//   	IdentityPoolName: jsii.String("identityPoolName"),
//   	IdentityProviderName: jsii.String("identityProviderName"),
//   	ProjectName: jsii.String("projectName"),
//   	ProjectNumber: jsii.String("projectNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-googleoidcconfig.html
//
type CfnLocationObjectStoragePropsMixin_GoogleOidcConfigProperty struct {
	// The name of the Google Cloud workload identity pool that DataSync federates with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-googleoidcconfig.html#cfn-datasync-locationobjectstorage-googleoidcconfig-identitypoolname
	//
	IdentityPoolName *string `field:"optional" json:"identityPoolName" yaml:"identityPoolName"`
	// The name of the OIDC identity provider configured in the Google Cloud workload identity pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-googleoidcconfig.html#cfn-datasync-locationobjectstorage-googleoidcconfig-identityprovidername
	//
	IdentityProviderName *string `field:"optional" json:"identityProviderName" yaml:"identityProviderName"`
	// The human-readable Google Cloud project name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-googleoidcconfig.html#cfn-datasync-locationobjectstorage-googleoidcconfig-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// The numeric Google Cloud project ID, as a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationobjectstorage-googleoidcconfig.html#cfn-datasync-locationobjectstorage-googleoidcconfig-projectnumber
	//
	ProjectNumber *string `field:"optional" json:"projectNumber" yaml:"projectNumber"`
}

