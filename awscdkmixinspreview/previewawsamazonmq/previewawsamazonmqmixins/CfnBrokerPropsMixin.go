package previewawsamazonmqmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsamazonmq/previewawsamazonmqmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a broker. Note: This API is asynchronous.
//
// To create a broker, you must either use the `AmazonMQFullAccess` IAM policy or include the following EC2 permissions in your IAM policy.
//
// - `ec2:CreateNetworkInterface`
//
// This permission is required to allow Amazon MQ to create an elastic network interface (ENI) on behalf of your account.
// - `ec2:CreateNetworkInterfacePermission`
//
// This permission is required to attach the ENI to the broker instance.
// - `ec2:DeleteNetworkInterface`
// - `ec2:DeleteNetworkInterfacePermission`
// - `ec2:DetachNetworkInterface`
// - `ec2:DescribeInternetGateways`
// - `ec2:DescribeNetworkInterfaces`
// - `ec2:DescribeNetworkInterfacePermissions`
// - `ec2:DescribeRouteTables`
// - `ec2:DescribeSecurityGroups`
// - `ec2:DescribeSubnets`
// - `ec2:DescribeVpcs`
//
// For more information, see [Create an IAM User and Get Your AWS Credentials](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/amazon-mq-setting-up.html#create-iam-user) and [Never Modify or Delete the Amazon MQ Elastic Network Interface](https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/connecting-to-amazon-mq.html#never-modify-delete-elastic-network-interface) in the *Amazon MQ Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrokerPropsMixin := awscdkmixinspreview.Mixins.NewCfnBrokerPropsMixin(&CfnBrokerMixinProps{
//   	AuthenticationStrategy: jsii.String("authenticationStrategy"),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	BrokerName: jsii.String("brokerName"),
//   	Configuration: &ConfigurationIdProperty{
//   		Id: jsii.String("id"),
//   		Revision: jsii.Number(123),
//   	},
//   	DataReplicationMode: jsii.String("dataReplicationMode"),
//   	DataReplicationPrimaryBrokerArn: jsii.String("dataReplicationPrimaryBrokerArn"),
//   	DeploymentMode: jsii.String("deploymentMode"),
//   	EncryptionOptions: &EncryptionOptionsProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		UseAwsOwnedKey: jsii.Boolean(false),
//   	},
//   	EngineType: jsii.String("engineType"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	HostInstanceType: jsii.String("hostInstanceType"),
//   	LdapServerMetadata: &LdapServerMetadataProperty{
//   		Hosts: []*string{
//   			jsii.String("hosts"),
//   		},
//   		RoleBase: jsii.String("roleBase"),
//   		RoleName: jsii.String("roleName"),
//   		RoleSearchMatching: jsii.String("roleSearchMatching"),
//   		RoleSearchSubtree: jsii.Boolean(false),
//   		ServiceAccountPassword: jsii.String("serviceAccountPassword"),
//   		ServiceAccountUsername: jsii.String("serviceAccountUsername"),
//   		UserBase: jsii.String("userBase"),
//   		UserRoleName: jsii.String("userRoleName"),
//   		UserSearchMatching: jsii.String("userSearchMatching"),
//   		UserSearchSubtree: jsii.Boolean(false),
//   	},
//   	Logs: &LogListProperty{
//   		Audit: jsii.Boolean(false),
//   		General: jsii.Boolean(false),
//   	},
//   	MaintenanceWindowStartTime: &MaintenanceWindowProperty{
//   		DayOfWeek: jsii.String("dayOfWeek"),
//   		TimeOfDay: jsii.String("timeOfDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   	PubliclyAccessible: jsii.Boolean(false),
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	StorageType: jsii.String("storageType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Users: []interface{}{
//   		&UserProperty{
//   			ConsoleAccess: jsii.Boolean(false),
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			Password: jsii.String("password"),
//   			ReplicationUser: jsii.Boolean(false),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html
//
type CfnBrokerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBrokerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBrokerPropsMixin
type jsiiProxy_CfnBrokerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBrokerPropsMixin) Props() *CfnBrokerMixinProps {
	var returns *CfnBrokerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBrokerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AmazonMQ::Broker`.
func NewCfnBrokerPropsMixin(props *CfnBrokerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBrokerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBrokerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBrokerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AmazonMQ::Broker`.
func NewCfnBrokerPropsMixin_Override(c CfnBrokerPropsMixin, props *CfnBrokerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBrokerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBrokerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBrokerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_amazonmq.mixins.CfnBrokerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBrokerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnBrokerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

