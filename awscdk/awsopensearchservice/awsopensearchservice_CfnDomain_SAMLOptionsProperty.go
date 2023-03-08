package awsopensearchservice


// Container for information about the SAML configuration for OpenSearch Dashboards.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sAMLOptionsProperty := &sAMLOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   	idp: &idpProperty{
//   		entityId: jsii.String("entityId"),
//   		metadataContent: jsii.String("metadataContent"),
//   	},
//   	masterBackendRole: jsii.String("masterBackendRole"),
//   	masterUserName: jsii.String("masterUserName"),
//   	rolesKey: jsii.String("rolesKey"),
//   	sessionTimeoutMinutes: jsii.Number(123),
//   	subjectKey: jsii.String("subjectKey"),
//   }
//
type CfnDomain_SAMLOptionsProperty struct {
	// True to enable SAML authentication for a domain.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The SAML Identity Provider's information.
	Idp interface{} `field:"optional" json:"idp" yaml:"idp"`
	// The backend role that the SAML master user is mapped to.
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// The SAML master user name, which is stored in the domain's internal user database.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Element of the SAML assertion to use for backend roles.
	//
	// Default is `roles` .
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// The duration, in minutes, after which a user session becomes inactive.
	//
	// Acceptable values are between 1 and 1440, and the default value is 60.
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Element of the SAML assertion to use for the user name.
	//
	// Default is `NameID` .
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

