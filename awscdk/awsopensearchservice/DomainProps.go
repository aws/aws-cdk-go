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
	AccessPolicies *[]awsiam.PolicyStatement `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Additional options to specify for the Amazon OpenSearch Service domain.
	// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options
	//
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// The hour in UTC during which the service takes an automated daily snapshot of the indices in the Amazon OpenSearch Service domain.
	//
	// Only applies for Elasticsearch versions
	// below 5.3.
	AutomatedSnapshotStartHour *float64 `field:"optional" json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
	// The cluster capacity configuration for the Amazon OpenSearch Service domain.
	Capacity *CapacityConfig `field:"optional" json:"capacity" yaml:"capacity"`
	// Configures Amazon OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
	CognitoDashboardsAuth *CognitoOptions `field:"optional" json:"cognitoDashboardsAuth" yaml:"cognitoDashboardsAuth"`
	// To configure a custom domain configure these options.
	//
	// If you specify a Route53 hosted zone it will create a CNAME record and use DNS validation for the certificate.
	CustomEndpoint *CustomEndpointOptions `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// Enforces a particular physical domain name.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon OpenSearch Service domain.
	Ebs *EbsOptions `field:"optional" json:"ebs" yaml:"ebs"`
	// To upgrade an Amazon OpenSearch Service domain to a new version, rather than replacing the entire domain resource, use the EnableVersionUpgrade update policy.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain
	//
	EnableVersionUpgrade *bool `field:"optional" json:"enableVersionUpgrade" yaml:"enableVersionUpgrade"`
	// Encryption at rest options for the cluster.
	EncryptionAtRest *EncryptionAtRestOptions `field:"optional" json:"encryptionAtRest" yaml:"encryptionAtRest"`
	// True to require that all traffic to the domain arrive over HTTPS.
	EnforceHttps *bool `field:"optional" json:"enforceHttps" yaml:"enforceHttps"`
	// Specifies options for fine-grained access control.
	//
	// Requires Elasticsearch version 6.7 or later or OpenSearch version 1.0 or later. Enabling fine-grained access control
	// also requires encryption of data at rest and node-to-node encryption, along with
	// enforced HTTPS.
	FineGrainedAccessControl *AdvancedSecurityOptions `field:"optional" json:"fineGrainedAccessControl" yaml:"fineGrainedAccessControl"`
	// Configuration log publishing configuration options.
	Logging *LoggingOptions `field:"optional" json:"logging" yaml:"logging"`
	// Specify true to enable node to node encryption.
	//
	// Requires Elasticsearch version 6.0 or later or OpenSearch version 1.0 or later.
	NodeToNodeEncryption *bool `field:"optional" json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// Policy to apply when the domain is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The list of security groups that are associated with the VPC endpoints for the domain.
	//
	// Only used if `vpc` is specified.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The minimum TLS version required for traffic to the domain.
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
	UseUnsignedBasicAuth *bool `field:"optional" json:"useUnsignedBasicAuth" yaml:"useUnsignedBasicAuth"`
	// Place the domain inside this VPC.
	// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html
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
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The cluster zone awareness configuration for the Amazon OpenSearch Service domain.
	ZoneAwareness *ZoneAwarenessConfig `field:"optional" json:"zoneAwareness" yaml:"zoneAwareness"`
}

