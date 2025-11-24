package mixinsawsopensearchservice


// Specifies options for fine-grained access control.
//
// If you specify advanced security options, you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html) ) and encryption at rest ( [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html) ). You must also enable `EnforceHTTPS` within [DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html) , which requires HTTPS for all traffic to the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advancedSecurityOptionsInputProperty := &AdvancedSecurityOptionsInputProperty{
//   	AnonymousAuthDisableDate: jsii.String("anonymousAuthDisableDate"),
//   	AnonymousAuthEnabled: jsii.Boolean(false),
//   	Enabled: jsii.Boolean(false),
//   	IamFederationOptions: map[string]interface{}{
//   		"enabled": jsii.Boolean(false),
//   		"rolesKey": jsii.String("rolesKey"),
//   		"subjectKey": jsii.String("subjectKey"),
//   	},
//   	InternalUserDatabaseEnabled: jsii.Boolean(false),
//   	JwtOptions: &JWTOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		PublicKey: jsii.String("publicKey"),
//   		RolesKey: jsii.String("rolesKey"),
//   		SubjectKey: jsii.String("subjectKey"),
//   	},
//   	MasterUserOptions: &MasterUserOptionsProperty{
//   		MasterUserArn: jsii.String("masterUserArn"),
//   		MasterUserName: jsii.String("masterUserName"),
//   		MasterUserPassword: jsii.String("masterUserPassword"),
//   	},
//   	SamlOptions: &SAMLOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		Idp: &IdpProperty{
//   			EntityId: jsii.String("entityId"),
//   			MetadataContent: jsii.String("metadataContent"),
//   		},
//   		MasterBackendRole: jsii.String("masterBackendRole"),
//   		MasterUserName: jsii.String("masterUserName"),
//   		RolesKey: jsii.String("rolesKey"),
//   		SessionTimeoutMinutes: jsii.Number(123),
//   		SubjectKey: jsii.String("subjectKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html
//
type CfnDomainPropsMixin_AdvancedSecurityOptionsInputProperty struct {
	// Date and time when the migration period will be disabled.
	//
	// Only necessary when [enabling fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-anonymousauthdisabledate
	//
	AnonymousAuthDisableDate *string `field:"optional" json:"anonymousAuthDisableDate" yaml:"anonymousAuthDisableDate"`
	// True to enable a 30-day migration period during which administrators can create role mappings.
	//
	// Only necessary when [enabling fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-anonymousauthenabled
	//
	AnonymousAuthEnabled interface{} `field:"optional" json:"anonymousAuthEnabled" yaml:"anonymousAuthEnabled"`
	// True to enable fine-grained access control.
	//
	// You must also enable encryption of data at rest and node-to-node encryption. See [Fine-grained access control in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Input configuration for IAM identity federation within advanced security options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-iamfederationoptions
	//
	IamFederationOptions interface{} `field:"optional" json:"iamFederationOptions" yaml:"iamFederationOptions"`
	// True to enable the internal user database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-internaluserdatabaseenabled
	//
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// Container for information about the JWT configuration of the Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-jwtoptions
	//
	JwtOptions interface{} `field:"optional" json:"jwtOptions" yaml:"jwtOptions"`
	// Specifies information about the master user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-masteruseroptions
	//
	MasterUserOptions interface{} `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
	// Container for information about the SAML configuration for OpenSearch Dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.html#cfn-opensearchservice-domain-advancedsecurityoptionsinput-samloptions
	//
	SamlOptions interface{} `field:"optional" json:"samlOptions" yaml:"samlOptions"`
}

