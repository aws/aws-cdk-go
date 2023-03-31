package awsopensearchservice


// Specifies options for fine-grained access control.
//
// If you specify advanced security options, you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html) ) and encryption at rest ( [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html) ). You must also enable `EnforceHTTPS` within [DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html) , which requires HTTPS for all traffic to the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedSecurityOptionsInputProperty := &advancedSecurityOptionsInputProperty{
//   	anonymousAuthDisableDate: jsii.String("anonymousAuthDisableDate"),
//   	anonymousAuthEnabled: jsii.Boolean(false),
//   	enabled: jsii.Boolean(false),
//   	internalUserDatabaseEnabled: jsii.Boolean(false),
//   	masterUserOptions: &masterUserOptionsProperty{
//   		masterUserArn: jsii.String("masterUserArn"),
//   		masterUserName: jsii.String("masterUserName"),
//   		masterUserPassword: jsii.String("masterUserPassword"),
//   	},
//   	samlOptions: &sAMLOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		idp: &idpProperty{
//   			entityId: jsii.String("entityId"),
//   			metadataContent: jsii.String("metadataContent"),
//   		},
//   		masterBackendRole: jsii.String("masterBackendRole"),
//   		masterUserName: jsii.String("masterUserName"),
//   		rolesKey: jsii.String("rolesKey"),
//   		sessionTimeoutMinutes: jsii.Number(123),
//   		subjectKey: jsii.String("subjectKey"),
//   	},
//   }
//
type CfnDomain_AdvancedSecurityOptionsInputProperty struct {
	// Date and time when the migration period will be disabled.
	//
	// Only necessary when [enabling fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing) .
	AnonymousAuthDisableDate *string `field:"optional" json:"anonymousAuthDisableDate" yaml:"anonymousAuthDisableDate"`
	// True to enable a 30-day migration period during which administrators can create role mappings.
	//
	// Only necessary when [enabling fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing) .
	AnonymousAuthEnabled interface{} `field:"optional" json:"anonymousAuthEnabled" yaml:"anonymousAuthEnabled"`
	// True to enable fine-grained access control.
	//
	// You must also enable encryption of data at rest and node-to-node encryption. See [Fine-grained access control in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html) .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// True to enable the internal user database.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// Specifies information about the master user.
	MasterUserOptions interface{} `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
	// Container for information about the SAML configuration for OpenSearch Dashboards.
	SamlOptions interface{} `field:"optional" json:"samlOptions" yaml:"samlOptions"`
}

