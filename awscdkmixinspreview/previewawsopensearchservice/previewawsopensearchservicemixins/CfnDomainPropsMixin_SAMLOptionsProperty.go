package previewawsopensearchservicemixins


// Container for information about the SAML configuration for OpenSearch Dashboards.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sAMLOptionsProperty := &SAMLOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	Idp: &IdpProperty{
//   		EntityId: jsii.String("entityId"),
//   		MetadataContent: jsii.String("metadataContent"),
//   	},
//   	MasterBackendRole: jsii.String("masterBackendRole"),
//   	MasterUserName: jsii.String("masterUserName"),
//   	RolesKey: jsii.String("rolesKey"),
//   	SessionTimeoutMinutes: jsii.Number(123),
//   	SubjectKey: jsii.String("subjectKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html
//
type CfnDomainPropsMixin_SAMLOptionsProperty struct {
	// True to enable SAML authentication for a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The SAML Identity Provider's information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-idp
	//
	Idp interface{} `field:"optional" json:"idp" yaml:"idp"`
	// The backend role that the SAML master user is mapped to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-masterbackendrole
	//
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// The SAML master user name, which is stored in the domain's internal user database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-masterusername
	//
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Element of the SAML assertion to use for backend roles.
	//
	// Default is `roles` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-roleskey
	//
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// The duration, in minutes, after which a user session becomes inactive.
	//
	// Acceptable values are between 1 and 1440, and the default value is 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-sessiontimeoutminutes
	//
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Element of the SAML assertion to use for the user name.
	//
	// Default is `NameID` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-samloptions.html#cfn-opensearchservice-domain-samloptions-subjectkey
	//
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

