package awsopensearchservice


// Container for information about the SAML configuration for OpenSearch Dashboards.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	EnforceHttps: jsii.Boolean(true),
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	FineGrainedAccessControl: &AdvancedSecurityOptions{
//   		MasterUserName: jsii.String("master-user"),
//   		SamlAuthenticationEnabled: jsii.Boolean(true),
//   		SamlAuthenticationOptions: &SAMLOptionsProperty{
//   			IdpEntityId: jsii.String("entity-id"),
//   			IdpMetadataContent: jsii.String("metadata-content-with-quotes-escaped"),
//   		},
//   	},
//   })
//
type SAMLOptionsProperty struct {
	// The unique entity ID of the application in the SAML identity provider.
	IdpEntityId *string `field:"required" json:"idpEntityId" yaml:"idpEntityId"`
	// The metadata of the SAML application, in XML format.
	IdpMetadataContent *string `field:"required" json:"idpMetadataContent" yaml:"idpMetadataContent"`
	// The backend role that the SAML master user is mapped to.
	//
	// Any users with this backend role receives full permission in OpenSearch Dashboards/Kibana.
	// To use a SAML master backend role, configure the `rolesKey` property.
	// Default: - The master user is not mapped to a backend role.
	//
	MasterBackendRole *string `field:"optional" json:"masterBackendRole" yaml:"masterBackendRole"`
	// The SAML master username, which is stored in the domain's internal user database.
	//
	// This SAML user receives full permission in OpenSearch Dashboards/Kibana.
	// Creating a new master username does not delete any existing master usernames.
	// Default: - No master user name is configured.
	//
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Element of the SAML assertion to use for backend roles.
	// Default: - roles.
	//
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// The duration, in minutes, after which a user session becomes inactive.
	// Default: - 60.
	//
	SessionTimeoutMinutes *float64 `field:"optional" json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Element of the SAML assertion to use for the user name.
	// Default: - NameID element of the SAML assertion fot the user name.
	//
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

