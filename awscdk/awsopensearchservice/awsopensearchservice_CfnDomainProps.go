package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnDomainProps := &cfnDomainProps{
//   	accessPolicies: accessPolicies,
//   	advancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	advancedSecurityOptions: &advancedSecurityOptionsInputProperty{
//   		enabled: jsii.Boolean(false),
//   		internalUserDatabaseEnabled: jsii.Boolean(false),
//   		masterUserOptions: &masterUserOptionsProperty{
//   			masterUserArn: jsii.String("masterUserArn"),
//   			masterUserName: jsii.String("masterUserName"),
//   			masterUserPassword: jsii.String("masterUserPassword"),
//   		},
//   	},
//   	clusterConfig: &clusterConfigProperty{
//   		dedicatedMasterCount: jsii.Number(123),
//   		dedicatedMasterEnabled: jsii.Boolean(false),
//   		dedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		instanceCount: jsii.Number(123),
//   		instanceType: jsii.String("instanceType"),
//   		warmCount: jsii.Number(123),
//   		warmEnabled: jsii.Boolean(false),
//   		warmType: jsii.String("warmType"),
//   		zoneAwarenessConfig: &zoneAwarenessConfigProperty{
//   			availabilityZoneCount: jsii.Number(123),
//   		},
//   		zoneAwarenessEnabled: jsii.Boolean(false),
//   	},
//   	cognitoOptions: &cognitoOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		identityPoolId: jsii.String("identityPoolId"),
//   		roleArn: jsii.String("roleArn"),
//   		userPoolId: jsii.String("userPoolId"),
//   	},
//   	domainEndpointOptions: &domainEndpointOptionsProperty{
//   		customEndpoint: jsii.String("customEndpoint"),
//   		customEndpointCertificateArn: jsii.String("customEndpointCertificateArn"),
//   		customEndpointEnabled: jsii.Boolean(false),
//   		enforceHttps: jsii.Boolean(false),
//   		tlsSecurityPolicy: jsii.String("tlsSecurityPolicy"),
//   	},
//   	domainName: jsii.String("domainName"),
//   	ebsOptions: &eBSOptionsProperty{
//   		ebsEnabled: jsii.Boolean(false),
//   		iops: jsii.Number(123),
//   		throughput: jsii.Number(123),
//   		volumeSize: jsii.Number(123),
//   		volumeType: jsii.String("volumeType"),
//   	},
//   	encryptionAtRestOptions: &encryptionAtRestOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	engineVersion: jsii.String("engineVersion"),
//   	logPublishingOptions: map[string]interface{}{
//   		"logPublishingOptionsKey": &LogPublishingOptionProperty{
//   			"cloudWatchLogsLogGroupArn": jsii.String("cloudWatchLogsLogGroupArn"),
//   			"enabled": jsii.Boolean(false),
//   		},
//   	},
//   	nodeToNodeEncryptionOptions: &nodeToNodeEncryptionOptionsProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	snapshotOptions: &snapshotOptionsProperty{
//   		automatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcOptions: &vPCOptionsProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDomainProps struct {
	// An AWS Identity and Access Management ( IAM ) policy document that specifies who can access the OpenSearch Service domain and their permissions.
	//
	// For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the *Amazon OpenSearch Service Developer Guide* .
	AccessPolicies interface{} `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Additional options to specify for the OpenSearch Service domain.
	//
	// For more information, see [AdvancedOptions](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-advancedoptions) in the OpenSearch Service configuration API reference.
	AdvancedOptions interface{} `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// Specifies options for fine-grained access control.
	AdvancedSecurityOptions interface{} `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// `ClusterConfig` is a property of the AWS::OpenSearchService::Domain resource that configures an Amazon OpenSearch Service cluster.
	ClusterConfig interface{} `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	CognitoOptions interface{} `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.
	DomainEndpointOptions interface{} `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// A name for the OpenSearch Service domain.
	//
	// For valid values, see the [DomainName](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-domainname) data type in the *Amazon OpenSearch Service Developer Guide* . If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the domain name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// Required when creating a new domain.
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the OpenSearch Service domain.
	//
	// For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the *Amazon OpenSearch Service Developer Guide* .
	EbsOptions interface{} `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// Whether the domain should encrypt data at rest, and if so, the AWS KMS key to use.
	//
	// See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html) .
	EncryptionAtRestOptions interface{} `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// The version of OpenSearch to use.
	//
	// The value must be in the format `OpenSearch_X.Y` or `Elasticsearch_X.Y` . If not specified, the latest version of OpenSearch is used. For information about the versions that OpenSearch Service supports, see [Supported versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the *Amazon OpenSearch Service Developer Guide* .
	//
	// If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true` , you can update `EngineVersion` without interruption. When `EnableVersionUpgrade` is set to `false` , or is not specified, updating `EngineVersion` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// An object with one or more of the following keys: `SEARCH_SLOW_LOGS` , `ES_APPLICATION_LOGS` , `INDEX_SLOW_LOGS` , `AUDIT_LOGS` , depending on the types of logs you want to publish.
	//
	// Each key needs a valid `LogPublishingOption` value. For the full syntax, see the [examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--examples) .
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// Specifies whether node-to-node encryption is enabled.
	//
	// See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html) .
	NodeToNodeEncryptionOptions interface{} `field:"optional" json:"nodeToNodeEncryptionOptions" yaml:"nodeToNodeEncryptionOptions"`
	// *DEPRECATED* .
	//
	// The automated snapshot configuration for the OpenSearch Service domain indices.
	SnapshotOptions interface{} `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The virtual private cloud (VPC) configuration for the OpenSearch Service domain.
	//
	// For more information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the *Amazon OpenSearch Service Developer Guide* .
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

