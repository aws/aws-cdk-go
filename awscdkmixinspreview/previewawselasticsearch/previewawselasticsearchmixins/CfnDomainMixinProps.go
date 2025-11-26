package previewawselasticsearchmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDomainPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var accessPolicies interface{}
//
//   cfnDomainMixinProps := &CfnDomainMixinProps{
//   	AccessPolicies: accessPolicies,
//   	AdvancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	AdvancedSecurityOptions: &AdvancedSecurityOptionsInputProperty{
//   		AnonymousAuthEnabled: jsii.Boolean(false),
//   		Enabled: jsii.Boolean(false),
//   		InternalUserDatabaseEnabled: jsii.Boolean(false),
//   		MasterUserOptions: &MasterUserOptionsProperty{
//   			MasterUserArn: jsii.String("masterUserArn"),
//   			MasterUserName: jsii.String("masterUserName"),
//   			MasterUserPassword: jsii.String("masterUserPassword"),
//   		},
//   	},
//   	CognitoOptions: &CognitoOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		IdentityPoolId: jsii.String("identityPoolId"),
//   		RoleArn: jsii.String("roleArn"),
//   		UserPoolId: jsii.String("userPoolId"),
//   	},
//   	DomainArn: jsii.String("domainArn"),
//   	DomainEndpointOptions: &DomainEndpointOptionsProperty{
//   		CustomEndpoint: jsii.String("customEndpoint"),
//   		CustomEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   		CustomEndpointEnabled: jsii.Boolean(false),
//   		EnforceHttps: jsii.Boolean(false),
//   		TlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   	},
//   	DomainName: jsii.String("domainName"),
//   	EbsOptions: &EBSOptionsProperty{
//   		EbsEnabled: jsii.Boolean(false),
//   		Iops: jsii.Number(123),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	ElasticsearchClusterConfig: &ElasticsearchClusterConfigProperty{
//   		ColdStorageOptions: &ColdStorageOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DedicatedMasterCount: jsii.Number(123),
//   		DedicatedMasterEnabled: jsii.Boolean(false),
//   		DedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		WarmCount: jsii.Number(123),
//   		WarmEnabled: jsii.Boolean(false),
//   		WarmType: jsii.String("warmType"),
//   		ZoneAwarenessConfig: &ZoneAwarenessConfigProperty{
//   			AvailabilityZoneCount: jsii.Number(123),
//   		},
//   		ZoneAwarenessEnabled: jsii.Boolean(false),
//   	},
//   	ElasticsearchVersion: jsii.String("elasticsearchVersion"),
//   	EncryptionAtRestOptions: &EncryptionAtRestOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	LogPublishingOptions: map[string]interface{}{
//   		"logPublishingOptionsKey": &LogPublishingOptionProperty{
//   			"cloudWatchLogsLogGroupArn": jsii.String("cloudWatchLogsLogGroupArn"),
//   			"enabled": jsii.Boolean(false),
//   		},
//   	},
//   	NodeToNodeEncryptionOptions: &NodeToNodeEncryptionOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	SnapshotOptions: &SnapshotOptionsProperty{
//   		AutomatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcOptions: &VPCOptionsProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html
//
type CfnDomainMixinProps struct {
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	//
	// For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guid* e.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-accesspolicies
	//
	AccessPolicies interface{} `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Additional options to specify for the OpenSearch Service domain.
	//
	// For more information, see [Advanced cluster parameters](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) in the *Amazon OpenSearch Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-advancedoptions
	//
	AdvancedOptions interface{} `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// Specifies options for fine-grained access control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-advancedsecurityoptions
	//
	AdvancedSecurityOptions interface{} `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-cognitooptions
	//
	CognitoOptions interface{} `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-domainarn
	//
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-domainendpointoptions
	//
	DomainEndpointOptions interface{} `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// A name for the OpenSearch Service domain.
	//
	// For valid values, see the [DomainName](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-domainname) data type in the *Amazon OpenSearch Service Developer Guide* . If you don't specify a name, CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	//
	// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-ebsoptions
	//
	EbsOptions interface{} `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// ElasticsearchClusterConfig is a property of the AWS::Elasticsearch::Domain resource that configures the cluster of an Amazon OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchclusterconfig
	//
	ElasticsearchClusterConfig interface{} `field:"optional" json:"elasticsearchClusterConfig" yaml:"elasticsearchClusterConfig"`
	// The version of Elasticsearch to use, such as 2.3. If not specified, 1.5 is used as the default. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `ElasticsearchVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `ElasticsearchVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchversion
	//
	ElasticsearchVersion *string `field:"optional" json:"elasticsearchVersion" yaml:"elasticsearchVersion"`
	// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service key to use.
	//
	// See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-encryptionatrestoptions
	//
	EncryptionAtRestOptions interface{} `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish.
	//
	// Each key needs a valid `LogPublishingOption` value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-logpublishingoptions
	//
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// Specifies whether node-to-node encryption is enabled.
	//
	// See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-nodetonodeencryptionoptions
	//
	NodeToNodeEncryptionOptions interface{} `field:"optional" json:"nodeToNodeEncryptionOptions" yaml:"nodeToNodeEncryptionOptions"`
	// *DEPRECATED* .
	//
	// The automated snapshot configuration for the OpenSearch Service domain indices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-snapshotoptions
	//
	SnapshotOptions interface{} `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	//
	// For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-vpcoptions
	//
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

