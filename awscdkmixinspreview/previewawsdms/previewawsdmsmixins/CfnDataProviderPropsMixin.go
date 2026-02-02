package previewawsdmsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdms/previewawsdmsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provides information that defines a data provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDataProviderPropsMixin := awscdkmixinspreview.Mixins.NewCfnDataProviderPropsMixin(&CfnDataProviderMixinProps{
//   	DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   	DataProviderName: jsii.String("dataProviderName"),
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	ExactSettings: jsii.Boolean(false),
//   	Settings: &SettingsProperty{
//   		DocDbSettings: &DocDbSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		IbmDb2LuwSettings: &IbmDb2LuwSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		IbmDb2ZOsSettings: &IbmDb2zOsSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MariaDbSettings: &MariaDbSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MongoDbSettings: &MongoDbSettingsProperty{
//   			AuthMechanism: jsii.String("authMechanism"),
//   			AuthSource: jsii.String("authSource"),
//   			AuthType: jsii.String("authType"),
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		MySqlSettings: &MySqlSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		OracleSettings: &OracleSettingsProperty{
//   			AsmServer: jsii.String("asmServer"),
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   			SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   			SecretsManagerSecurityDbEncryptionAccessRoleArn: jsii.String("secretsManagerSecurityDbEncryptionAccessRoleArn"),
//   			SecretsManagerSecurityDbEncryptionSecretId: jsii.String("secretsManagerSecurityDbEncryptionSecretId"),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   		RedshiftSettings: &RedshiftSettingsProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   		},
//   		SybaseAseSettings: &SybaseAseSettingsProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			DatabaseName: jsii.String("databaseName"),
//   			EncryptPassword: jsii.Boolean(false),
//   			Port: jsii.Number(123),
//   			ServerName: jsii.String("serverName"),
//   			SslMode: jsii.String("sslMode"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-dataprovider.html
//
type CfnDataProviderPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDataProviderMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDataProviderPropsMixin
type jsiiProxy_CfnDataProviderPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDataProviderPropsMixin) Props() *CfnDataProviderMixinProps {
	var returns *CfnDataProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataProviderPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DMS::DataProvider`.
func NewCfnDataProviderPropsMixin(props *CfnDataProviderMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDataProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDataProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DMS::DataProvider`.
func NewCfnDataProviderPropsMixin_Override(c CfnDataProviderPropsMixin, props *CfnDataProviderMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDataProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnDataProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataProviderPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDataProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

