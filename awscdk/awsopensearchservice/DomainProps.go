package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for an Amazon OpenSearch Service domain.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	Ebs: &EbsOptions{
//   		VolumeSize: jsii.Number(100),
//   		VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD,
//   	},
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   })
//
type DomainProps struct {
	// The Elasticsearch/OpenSearch version that your domain will leverage.
	Version EngineVersion `field:"required" json:"version" yaml:"version"`
	// Domain access policies.
	// Default: - No access policies.
	//
	AccessPolicies *[]awsiam.PolicyStatement `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Additional options to specify for the Amazon OpenSearch Service domain.
	// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options
	//
	// Default: - no advanced options are specified.
	//
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// The hour in UTC during which the service takes an automated daily snapshot of the indices in the Amazon OpenSearch Service domain.
	//
	// Only applies for Elasticsearch versions
	// below 5.3.
	// Default: - Hourly automated snapshots not used.
	//
	AutomatedSnapshotStartHour *float64 `field:"optional" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
	// The cluster capacity configuration for the Amazon OpenSearch Service domain.
	// Default: - 1 r5.large.search data node; no dedicated master nodes.
	//
	Capacity *CapacityConfig `field:"optional" json:"capacity" yaml:"capacity"`
	// Configures Amazon OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	// Default: - Cognito not used for authentication to OpenSearch Dashboards.
	//
	CognitoDashboardsAuth *CognitoOptions `field:"optional" json:"cognitoDashboardsAuth" yaml:"cognitoDashboardsAuth"`
	// To configure a custom domain configure these options.
	//
	// If you specify a Route53 hosted zone it will create a CNAME record and use DNS validation for the certificate.
	// Default: - no custom domain endpoint will be configured.
	//
	CustomEndpoint *CustomEndpointOptions `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// Enforces a particular physical domain name.
	// Default: - A name will be auto-generated.
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon OpenSearch Service domain.
	// Default: - 10 GiB General Purpose (SSD) volumes per node.
	//
	Ebs *EbsOptions `field:"optional" json:"ebs" yaml:"ebs"`
	// Specifies whether automatic service software updates are enabled for the domain.
	// See: https://docs.aws.amazon.com/it_it/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-softwareupdateoptions.html
	//
	// Default: - false.
	//
	EnableAutoSoftwareUpdate *bool `field:"optional" json:"enableAutoSoftwareUpdate" yaml:"enableAutoSoftwareUpdate"`
	// To upgrade an Amazon OpenSearch Service domain to a new version, rather than replacing the entire domain resource, use the EnableVersionUpgrade update policy.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain
	//
	// Default: - false.
	//
	EnableVersionUpgrade *bool `field:"optional" json:"enableVersionUpgrade" yaml:"enableVersionUpgrade"`
	// Encryption at rest options for the cluster.
	// Default: - No encryption at rest.
	//
	EncryptionAtRest *EncryptionAtRestOptions `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// True to require that all traffic to the domain arrive over HTTPS.
	// Default: - false.
	//
	EnforceHttps *bool `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// Specifies options for fine-grained access control.
	//
	// Requires Elasticsearch version 6.7 or later or OpenSearch version 1.0 or later. Enabling fine-grained access control
	// also requires encryption of data at rest and node-to-node encryption, along with
	// enforced HTTPS.
	// Default: - fine-grained access control is disabled.
	//
	FineGrainedAccessControl *AdvancedSecurityOptions `field:"optional" json:"fineGrainedAccessControl" yaml:"fineGrainedAccessControl"`
	// Configuration log publishing configuration options.
	// Default: - No logs are published.
	//
	Logging *LoggingOptions `field:"optional" json:"logging" yaml:"logging"`
	// Specify true to enable node to node encryption.
	//
	// Requires Elasticsearch version 6.0 or later or OpenSearch version 1.0 or later.
	// Default: - Node to node encryption is not enabled.
	//
	NodeToNodeEncryption *bool `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// Options for enabling a domain's off-peak window, during which OpenSearch Service can perform mandatory configuration changes on the domain.
	//
	// Off-peak windows were introduced on February 16, 2023.
	// All domains created before this date have the off-peak window disabled by default.
	// You must manually enable and configure the off-peak window for these domains.
	// All domains created after this date will have the off-peak window enabled by default.
	// You can't disable the off-peak window for a domain after it's enabled.
	// See: https://docs.aws.amazon.com/it_it/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-offpeakwindow.html
	//
	// Default: - Disabled for domains created before February 16, 2023. Enabled for domains created after. Enabled if `offPeakWindowStart` is set.
	//
	OffPeakWindowEnabled *bool `field:"optional" json:"offPeakWindowEnabled" yaml:"offPeakWindowEnabled"`
	// Start time for the off-peak window, in Coordinated Universal Time (UTC).
	//
	// The window length will always be 10 hours, so you can't specify an end time.
	// For example, if you specify 11:00 P.M. UTC as a start time, the end time will automatically be set to 9:00 A.M.
	// Default: - 10:00 P.M. local time
	//
	OffPeakWindowStart *WindowStartTime `field:"optional" json:"offPeakWindowStart" yaml:"offPeakWindowStart"`
	// Policy to apply when the domain is removed from the stack.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The list of security groups that are associated with the VPC endpoints for the domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html
	//
	// Default: - One new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The minimum TLS version required for traffic to the domain.
	// Default: - TLSSecurityPolicy.TLS_1_0
	//
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
	// Default: - false.
	//
	UseUnsignedBasicAuth *bool `field:"optional" json:"useUnsignedBasicAuth" yaml:"useUnsignedBasicAuth"`
	// Place the domain inside this VPC.
	// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html
	//
	// Default: - Domain is not placed in a VPC.
	//
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
	// Default: - All private subnets.
	//
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The cluster zone awareness configuration for the Amazon OpenSearch Service domain.
	// Default: - no zone awareness (1 AZ).
	//
	ZoneAwareness *ZoneAwarenessConfig `field:"optional" json:"zoneAwareness" yaml:"zoneAwareness"`
}

