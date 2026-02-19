package previewawsopensearchservicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsopensearchservice/previewawsopensearchservicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::OpenSearchService::Domain resource creates an Amazon OpenSearch Service domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var accessPolicies interface{}
//
//   cfnDomainPropsMixin := awscdkmixinspreview.Mixins.NewCfnDomainPropsMixin(&CfnDomainMixinProps{
//   	AccessPolicies: accessPolicies,
//   	AdvancedOptions: map[string]*string{
//   		"advancedOptionsKey": jsii.String("advancedOptions"),
//   	},
//   	AdvancedSecurityOptions: &AdvancedSecurityOptionsInputProperty{
//   		AnonymousAuthDisableDate: jsii.String("anonymousAuthDisableDate"),
//   		AnonymousAuthEnabled: jsii.Boolean(false),
//   		Enabled: jsii.Boolean(false),
//   		IamFederationOptions: map[string]interface{}{
//   			"enabled": jsii.Boolean(false),
//   			"rolesKey": jsii.String("rolesKey"),
//   			"subjectKey": jsii.String("subjectKey"),
//   		},
//   		InternalUserDatabaseEnabled: jsii.Boolean(false),
//   		JwtOptions: &JWTOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			PublicKey: jsii.String("publicKey"),
//   			RolesKey: jsii.String("rolesKey"),
//   			SubjectKey: jsii.String("subjectKey"),
//   		},
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
//   	AimlOptions: &AIMLOptionsProperty{
//   		S3VectorsEngine: &S3VectorsEngineProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		ServerlessVectorAcceleration: &ServerlessVectorAccelerationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	ClusterConfig: &ClusterConfigProperty{
//   		ColdStorageOptions: &ColdStorageOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		DedicatedMasterCount: jsii.Number(123),
//   		DedicatedMasterEnabled: jsii.Boolean(false),
//   		DedicatedMasterType: jsii.String("dedicatedMasterType"),
//   		InstanceCount: jsii.Number(123),
//   		InstanceType: jsii.String("instanceType"),
//   		MultiAzWithStandbyEnabled: jsii.Boolean(false),
//   		NodeOptions: []interface{}{
//   			&NodeOptionProperty{
//   				NodeConfig: &NodeConfigProperty{
//   					Count: jsii.Number(123),
//   					Enabled: jsii.Boolean(false),
//   					Type: jsii.String("type"),
//   				},
//   				NodeType: jsii.String("nodeType"),
//   			},
//   		},
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
//   	IdentityCenterOptions: &IdentityCenterOptionsProperty{
//   		EnabledApiAccess: jsii.Boolean(false),
//   		IdentityCenterApplicationArn: jsii.String("identityCenterApplicationArn"),
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   		IdentityStoreId: jsii.String("identityStoreId"),
//   		RolesKey: jsii.String("rolesKey"),
//   		SubjectKey: jsii.String("subjectKey"),
//   	},
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
//   	SkipShardMigrationWait: jsii.Boolean(false),
//   	SnapshotOptions: &SnapshotOptionsProperty{
//   		AutomatedSnapshotStartHour: jsii.Number(123),
//   	},
//   	SoftwareUpdateOptions: &SoftwareUpdateOptionsProperty{
//   		AutoSoftwareUpdateEnabled: jsii.Boolean(false),
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html
//
type CfnDomainPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDomainMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDomainPropsMixin
type jsiiProxy_CfnDomainPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDomainPropsMixin) Props() *CfnDomainMixinProps {
	var returns *CfnDomainMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::OpenSearchService::Domain`.
func NewCfnDomainPropsMixin(props *CfnDomainMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDomainPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDomainPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDomainPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::OpenSearchService::Domain`.
func NewCfnDomainPropsMixin_Override(c CfnDomainPropsMixin, props *CfnDomainMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDomainPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDomainPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDomainPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDomainPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

