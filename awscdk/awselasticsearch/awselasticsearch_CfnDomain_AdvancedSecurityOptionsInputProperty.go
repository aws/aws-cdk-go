package awselasticsearch


// Specifies options for fine-grained access control.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedSecurityOptionsInputProperty := &advancedSecurityOptionsInputProperty{
//   	anonymousAuthEnabled: jsii.Boolean(false),
//   	enabled: jsii.Boolean(false),
//   	internalUserDatabaseEnabled: jsii.Boolean(false),
//   	masterUserOptions: &masterUserOptionsProperty{
//   		masterUserArn: jsii.String("masterUserArn"),
//   		masterUserName: jsii.String("masterUserName"),
//   		masterUserPassword: jsii.String("masterUserPassword"),
//   	},
//   }
//
type CfnDomain_AdvancedSecurityOptionsInputProperty struct {
	// `CfnDomain.AdvancedSecurityOptionsInputProperty.AnonymousAuthEnabled`.
	AnonymousAuthEnabled interface{} `field:"optional" json:"anonymousAuthEnabled" yaml:"anonymousAuthEnabled"`
	// True to enable fine-grained access control.
	//
	// You must also enable encryption of data at rest and node-to-node encryption.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// True to enable the internal user database.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// Specifies information about the master user.
	MasterUserOptions interface{} `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

