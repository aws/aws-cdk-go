package previewawsdatazonemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdatazone/previewawsdatazonemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// In Amazon DataZone, a connection enables you to connect your resources (domains, projects, and environments) to external resources and services.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConnectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnConnectionPropsMixin(&CfnConnectionMixinProps{
//   	AwsLocation: &AwsLocationProperty{
//   		AccessRole: jsii.String("accessRole"),
//   		AwsAccountId: jsii.String("awsAccountId"),
//   		AwsRegion: jsii.String("awsRegion"),
//   		IamConnectionId: jsii.String("iamConnectionId"),
//   	},
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	EnableTrustedIdentityPropagation: jsii.Boolean(false),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	Props: &ConnectionPropertiesInputProperty{
//   		AmazonQProperties: &AmazonQPropertiesInputProperty{
//   			AuthMode: jsii.String("authMode"),
//   			IsEnabled: jsii.Boolean(false),
//   			ProfileArn: jsii.String("profileArn"),
//   		},
//   		AthenaProperties: &AthenaPropertiesInputProperty{
//   			WorkgroupName: jsii.String("workgroupName"),
//   		},
//   		GlueProperties: &GluePropertiesInputProperty{
//   			GlueConnectionInput: &GlueConnectionInputProperty{
//   				AthenaProperties: map[string]*string{
//   					"athenaPropertiesKey": jsii.String("athenaProperties"),
//   				},
//   				AuthenticationConfiguration: &AuthenticationConfigurationInputProperty{
//   					AuthenticationType: jsii.String("authenticationType"),
//   					BasicAuthenticationCredentials: &BasicAuthenticationCredentialsProperty{
//   						Password: jsii.String("password"),
//   						UserName: jsii.String("userName"),
//   					},
//   					CustomAuthenticationCredentials: map[string]*string{
//   						"customAuthenticationCredentialsKey": jsii.String("customAuthenticationCredentials"),
//   					},
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					OAuth2Properties: &OAuth2PropertiesProperty{
//   						AuthorizationCodeProperties: &AuthorizationCodePropertiesProperty{
//   							AuthorizationCode: jsii.String("authorizationCode"),
//   							RedirectUri: jsii.String("redirectUri"),
//   						},
//   						OAuth2ClientApplication: &OAuth2ClientApplicationProperty{
//   							AwsManagedClientApplicationReference: jsii.String("awsManagedClientApplicationReference"),
//   							UserManagedClientApplicationClientId: jsii.String("userManagedClientApplicationClientId"),
//   						},
//   						OAuth2Credentials: &GlueOAuth2CredentialsProperty{
//   							AccessToken: jsii.String("accessToken"),
//   							JwtToken: jsii.String("jwtToken"),
//   							RefreshToken: jsii.String("refreshToken"),
//   							UserManagedClientApplicationClientSecret: jsii.String("userManagedClientApplicationClientSecret"),
//   						},
//   						OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   						TokenUrl: jsii.String("tokenUrl"),
//   						TokenUrlParametersMap: map[string]*string{
//   							"tokenUrlParametersMapKey": jsii.String("tokenUrlParametersMap"),
//   						},
//   					},
//   					SecretArn: jsii.String("secretArn"),
//   				},
//   				ConnectionProperties: map[string]*string{
//   					"connectionPropertiesKey": jsii.String("connectionProperties"),
//   				},
//   				ConnectionType: jsii.String("connectionType"),
//   				Description: jsii.String("description"),
//   				MatchCriteria: jsii.String("matchCriteria"),
//   				Name: jsii.String("name"),
//   				PhysicalConnectionRequirements: &PhysicalConnectionRequirementsProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					SecurityGroupIdList: []*string{
//   						jsii.String("securityGroupIdList"),
//   					},
//   					SubnetId: jsii.String("subnetId"),
//   					SubnetIdList: []*string{
//   						jsii.String("subnetIdList"),
//   					},
//   				},
//   				PythonProperties: map[string]*string{
//   					"pythonPropertiesKey": jsii.String("pythonProperties"),
//   				},
//   				SparkProperties: map[string]*string{
//   					"sparkPropertiesKey": jsii.String("sparkProperties"),
//   				},
//   				ValidateCredentials: jsii.Boolean(false),
//   				ValidateForComputeEnvironments: []*string{
//   					jsii.String("validateForComputeEnvironments"),
//   				},
//   			},
//   		},
//   		HyperPodProperties: &HyperPodPropertiesInputProperty{
//   			ClusterName: jsii.String("clusterName"),
//   		},
//   		IamProperties: &IamPropertiesInputProperty{
//   			GlueLineageSyncEnabled: jsii.Boolean(false),
//   		},
//   		RedshiftProperties: &RedshiftPropertiesInputProperty{
//   			Credentials: &RedshiftCredentialsProperty{
//   				SecretArn: jsii.String("secretArn"),
//   				UsernamePassword: &UsernamePasswordProperty{
//   					Password: jsii.String("password"),
//   					Username: jsii.String("username"),
//   				},
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Host: jsii.String("host"),
//   			LineageSync: &RedshiftLineageSyncConfigurationInputProperty{
//   				Enabled: jsii.Boolean(false),
//   				Schedule: &LineageSyncScheduleProperty{
//   					Schedule: jsii.String("schedule"),
//   				},
//   			},
//   			Port: jsii.Number(123),
//   			Storage: &RedshiftStoragePropertiesProperty{
//   				ClusterName: jsii.String("clusterName"),
//   				WorkgroupName: jsii.String("workgroupName"),
//   			},
//   		},
//   		S3Properties: &S3PropertiesInputProperty{
//   			S3AccessGrantLocationId: jsii.String("s3AccessGrantLocationId"),
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		SparkEmrProperties: &SparkEmrPropertiesInputProperty{
//   			ComputeArn: jsii.String("computeArn"),
//   			InstanceProfileArn: jsii.String("instanceProfileArn"),
//   			JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   			LogUri: jsii.String("logUri"),
//   			ManagedEndpointArn: jsii.String("managedEndpointArn"),
//   			PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   			RuntimeRole: jsii.String("runtimeRole"),
//   			TrustedCertificatesS3Uri: jsii.String("trustedCertificatesS3Uri"),
//   		},
//   		SparkGlueProperties: &SparkGluePropertiesInputProperty{
//   			AdditionalArgs: &SparkGlueArgsProperty{
//   				Connection: jsii.String("connection"),
//   			},
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   			GlueVersion: jsii.String("glueVersion"),
//   			IdleTimeout: jsii.Number(123),
//   			JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   			NumberOfWorkers: jsii.Number(123),
//   			PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   			WorkerType: jsii.String("workerType"),
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-connection.html
//
type CfnConnectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnConnectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnConnectionPropsMixin
type jsiiProxy_CfnConnectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnConnectionPropsMixin) Props() *CfnConnectionMixinProps {
	var returns *CfnConnectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::Connection`.
func NewCfnConnectionPropsMixin(props *CfnConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnConnectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnConnectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnConnectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::Connection`.
func NewCfnConnectionPropsMixin_Override(c CfnConnectionPropsMixin, props *CfnConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnConnectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnConnectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_datazone.mixins.CfnConnectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnConnectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

