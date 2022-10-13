package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for an AWS Elasticsearch Domain.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: es.elasticsearchVersion_V7_4(),
//   	ebs: &ebsOptions{
//   		volumeSize: jsii.Number(100),
//   		volumeType: ec2.ebsDeviceVolumeType_GENERAL_PURPOSE_SSD,
//   	},
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   })
//
// Deprecated: use opensearchservice module instead.
type DomainProps struct {
	// The Elasticsearch version that your domain will leverage.
	// Deprecated: use opensearchservice module instead.
	Version ElasticsearchVersion `field:"required" json:"version" yaml:"version"`
	// Domain Access policies.
	// Deprecated: use opensearchservice module instead.
	AccessPolicies *[]awsiam.PolicyStatement `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Additional options to specify for the Amazon ES domain.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options
	//
	// Deprecated: use opensearchservice module instead.
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// The hour in UTC during which the service takes an automated daily snapshot of the indices in the Amazon ES domain.
	//
	// Only applies for Elasticsearch
	// versions below 5.3.
	// Deprecated: use opensearchservice module instead.
	AutomatedSnapshotStartHour *float64 `field:"optional" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
	// The cluster capacity configuration for the Amazon ES domain.
	// Deprecated: use opensearchservice module instead.
	Capacity *CapacityConfig `field:"optional" json:"capacity" yaml:"capacity"`
	// Configures Amazon ES to use Amazon Cognito authentication for Kibana.
	// Deprecated: use opensearchservice module instead.
	CognitoKibanaAuth *CognitoOptions `field:"optional" json:"cognitoKibanaAuth" yaml:"cognitoKibanaAuth"`
	// To configure a custom domain configure these options.
	//
	// If you specify a Route53 hosted zone it will create a CNAME record and use DNS validation for the certificate.
	// Deprecated: use opensearchservice module instead.
	CustomEndpoint *CustomEndpointOptions `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// Enforces a particular physical domain name.
	// Deprecated: use opensearchservice module instead.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon ES domain.
	//
	// For more information, see
	// [Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Deprecated: use opensearchservice module instead.
	Ebs *EbsOptions `field:"optional" json:"ebs" yaml:"ebs"`
	// To upgrade an Amazon ES domain to a new version of Elasticsearch rather than replacing the entire domain resource, use the EnableVersionUpgrade update policy.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeelasticsearchdomain
	//
	// Deprecated: use opensearchservice module instead.
	EnableVersionUpgrade *bool `field:"optional" json:"enableVersionUpgrade" yaml:"enableVersionUpgrade"`
	// Encryption at rest options for the cluster.
	// Deprecated: use opensearchservice module instead.
	EncryptionAtRest *EncryptionAtRestOptions `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// True to require that all traffic to the domain arrive over HTTPS.
	// Deprecated: use opensearchservice module instead.
	EnforceHttps *bool `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// Specifies options for fine-grained access control.
	//
	// Requires Elasticsearch version 6.7 or later. Enabling fine-grained access control
	// also requires encryption of data at rest and node-to-node encryption, along with
	// enforced HTTPS.
	// Deprecated: use opensearchservice module instead.
	FineGrainedAccessControl *AdvancedSecurityOptions `field:"optional" json:"fineGrainedAccessControl" yaml:"fineGrainedAccessControl"`
	// Configuration log publishing configuration options.
	// Deprecated: use opensearchservice module instead.
	Logging *LoggingOptions `field:"optional" json:"logging" yaml:"logging"`
	// Specify true to enable node to node encryption.
	//
	// Requires Elasticsearch version 6.0 or later.
	// Deprecated: use opensearchservice module instead.
	NodeToNodeEncryption *bool `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// Policy to apply when the domain is removed from the stack.
	// Deprecated: use opensearchservice module instead.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The list of security groups that are associated with the VPC endpoints for the domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html
	//
	// Deprecated: use opensearchservice module instead.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The minimum TLS version required for traffic to the domain.
	// Deprecated: use opensearchservice module instead.
	TlsSecurityPolicy TLSSecurityPolicy `field:"optional" json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
	// Configures the domain so that unsigned basic auth is enabled.
	//
	// If no master user is provided a default master user
	// with username `admin` and a dynamically generated password stored in KMS is created. The password can be retrieved
	// by getting `masterUserPassword` from the domain instance.
	//
	// Setting this to true will also add an access policy that allows unsigned
	// access, enable node to node encryption, encryption at rest. If conflicting
	// settings are encountered (like disabling encryption at rest) enabling this
	// setting will cause a failure.
	// Deprecated: use opensearchservice module instead.
	UseUnsignedBasicAuth *bool `field:"optional" json:"useUnsignedBasicAuth" yaml:"useUnsignedBasicAuth"`
	// Place the domain inside this VPC.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html
	//
	// Deprecated: use opensearchservice module instead.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The specific vpc subnets the domain will be placed in.
	//
	// You must provide one subnet for each Availability Zone
	// that your domain uses. For example, you must specify three subnet IDs for a three Availability Zone
	// domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html
	//
	// Deprecated: use opensearchservice module instead.
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The cluster zone awareness configuration for the Amazon ES domain.
	// Deprecated: use opensearchservice module instead.
	ZoneAwareness *ZoneAwarenessConfig `field:"optional" json:"zoneAwareness" yaml:"zoneAwareness"`
}

