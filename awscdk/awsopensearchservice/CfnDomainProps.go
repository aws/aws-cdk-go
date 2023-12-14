package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDomain`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicies interface{}
//
//   cfnDomainProps := &CfnDomainProps{
//   	AccessPolicies: accessPolicies,
//   	AdvancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	AdvancedSecurityOptions: &AdvancedSecurityOptionsInputProperty{
//   		AnonymousAuthDisableDate: jsii.String("anonymousAuthDisableDate"),
//   		AnonymousAuthEnabled: jsii.Boolean(false),
//   		Enabled: jsii.Boolean(false),
//   		InternalUserDatabaseEnabled: jsii.Boolean(false),
//   		MasterUserOptions: &MasterUserOptionsProperty{
//   			MasterUserArn: jsii.String("masterUserArn"),
//   			MasterUserName: jsii.String("masterUserName"),
//   			MasterUserPassword: jsii.String("masterUserPassword"),
//   		},
//   		SamlOptions: &SAMLOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			Idp: &IdpProperty{
//   				EntityId: jsii.String("entityId"),
//   				MetadataContent: jsii.String("metadataContent"),
//   			},
//   			MasterBackendRole: jsii.String("masterBackendRole"),
//   			MasterUserName: jsii.String("masterUserName"),
//   			RolesKey: jsii.String("rolesKey"),
//   			SessionTimeoutMinutes: jsii.Number(123),
//   			SubjectKey: jsii.String("subjectKey"),
//   		},
//   	},
//   	ClusterConfig: &ClusterConfigProperty{
//   		DedicatedMasterCount: jsii.Number(123),
//   		DedicatedMasterEnabled: jsii.Boolean(false),
//   		DedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		MultiAzWithStandbyEnabled: jsii.Boolean(false),
//   		WarmCount: jsii.Number(123),
//   		WarmEnabled: jsii.Boolean(false),
//   		WarmType: jsii.String("warmType"),
//   		ZoneAwarenessConfig: &ZoneAwarenessConfigProperty{
//   			AvailabilityZoneCount: jsii.Number(123),
//   		},
//   		ZoneAwarenessEnabled: jsii.Boolean(false),
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
//   		Throughput: jsii.Number(123),
//   		VolumeSize: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	EncryptionAtRestOptions: &EncryptionAtRestOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	LogPublishingOptions: map[string]interface{}{
//   		"logPublishingOptionsKey": &LogPublishingOptionProperty{
//   			"cloudWatchLogsLogGroupArn": jsii.String("cloudWatchLogsLogGroupArn"),
//   			"enabled": jsii.Boolean(false),
//   		},
//   	},
//   	NodeToNodeEncryptionOptions: &NodeToNodeEncryptionOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	OffPeakWindowOptions: &OffPeakWindowOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		OffPeakWindow: &OffPeakWindowProperty{
//   			WindowStartTime: &WindowStartTimeProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SnapshotOptions: &SnapshotOptionsProperty{
//   		AutomatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	SoftwareUpdateOptions: &SoftwareUpdateOptionsProperty{
//   		AutoSoftwareUpdateEnabled: jsii.Boolean(false),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html
//
type CfnDomainProps struct {
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	//
	// For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-accesspolicies
	//
	AccessPolicies interface{} `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Additional options to specify for the OpenSearch Service domain.
	//
	// For more information, see [AdvancedOptions](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_CreateDomain.html#API_CreateDomain_RequestBody) in the OpenSearch Service API reference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-advancedoptions
	//
	AdvancedOptions interface{} `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// Specifies options for fine-grained access control and SAML authentication.
	//
	// If you specify advanced security options, you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html) ) and encryption at rest ( [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html) ). You must also enable `EnforceHTTPS` within [DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html) , which requires HTTPS for all traffic to the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-advancedsecurityoptions
	//
	AdvancedSecurityOptions interface{} `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// Container for the cluster configuration of a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-clusterconfig
	//
	ClusterConfig interface{} `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-cognitooptions
	//
	CognitoOptions interface{} `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-domainarn
	//
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-domainendpointoptions
	//
	DomainEndpointOptions interface{} `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// A name for the OpenSearch Service domain.
	//
	// The name must have a minimum length of 3 and a maximum length of 28. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// Required when creating a new domain.
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	//
	// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-ebsoptions
	//
	EbsOptions interface{} `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// Whether the domain should encrypt data at rest, and if so, the AWS KMS key to use.
	//
	// See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-encryptionatrestoptions
	//
	EncryptionAtRestOptions interface{} `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// The version of OpenSearch to use.
	//
	// The value must be in the format `OpenSearch_X.Y` or `Elasticsearch_X.Y` . If not specified, the latest version of OpenSearch is used. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `EngineVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `EngineVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish.
	//
	// Each key needs a valid `LogPublishingOption` value. For the full syntax, see the [examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--examples) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-logpublishingoptions
	//
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// Specifies whether node-to-node encryption is enabled.
	//
	// See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-nodetonodeencryptionoptions
	//
	NodeToNodeEncryptionOptions interface{} `field:"optional" json:"nodeToNodeEncryptionOptions" yaml:"nodeToNodeEncryptionOptions"`
	// Options for a domain's off-peak window, during which OpenSearch Service can perform mandatory configuration changes on the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-offpeakwindowoptions
	//
	OffPeakWindowOptions interface{} `field:"optional" json:"offPeakWindowOptions" yaml:"offPeakWindowOptions"`
	// *DEPRECATED* .
	//
	// The automated snapshot configuration for the OpenSearch Service domain indexes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-snapshotoptions
	//
	SnapshotOptions interface{} `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// Options for configuring service software updates for a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-softwareupdateoptions
	//
	SoftwareUpdateOptions interface{} `field:"optional" json:"softwareUpdateOptions" yaml:"softwareUpdateOptions"`
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	//
	// For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you remove this entity altogether, along with its associated properties, it causes a replacement. You might encounter this scenario if you're updating your security configuration from a VPC to a public endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#cfn-opensearchservice-domain-vpcoptions
	//
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

