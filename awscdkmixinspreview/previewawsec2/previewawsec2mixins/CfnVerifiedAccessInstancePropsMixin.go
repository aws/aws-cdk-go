package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Verified Access instance is a regional entity that evaluates application requests and grants access only when your security requirements are met.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVerifiedAccessInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnVerifiedAccessInstancePropsMixin(&CfnVerifiedAccessInstanceMixinProps{
//   	CidrEndpointsCustomSubDomain: jsii.String("cidrEndpointsCustomSubDomain"),
//   	Description: jsii.String("description"),
//   	FipsEnabled: jsii.Boolean(false),
//   	LoggingConfigurations: &VerifiedAccessLogsProperty{
//   		CloudWatchLogs: &CloudWatchLogsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   		},
//   		IncludeTrustContext: jsii.Boolean(false),
//   		KinesisDataFirehose: &KinesisDataFirehoseProperty{
//   			DeliveryStream: jsii.String("deliveryStream"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   		LogVersion: jsii.String("logVersion"),
//   		S3: &S3Property{
//   			BucketName: jsii.String("bucketName"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			Enabled: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VerifiedAccessTrustProviderIds: []*string{
//   		jsii.String("verifiedAccessTrustProviderIds"),
//   	},
//   	VerifiedAccessTrustProviders: []interface{}{
//   		&VerifiedAccessTrustProviderProperty{
//   			Description: jsii.String("description"),
//   			DeviceTrustProviderType: jsii.String("deviceTrustProviderType"),
//   			TrustProviderType: jsii.String("trustProviderType"),
//   			UserTrustProviderType: jsii.String("userTrustProviderType"),
//   			VerifiedAccessTrustProviderId: jsii.String("verifiedAccessTrustProviderId"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-verifiedaccessinstance.html
//
type CfnVerifiedAccessInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVerifiedAccessInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVerifiedAccessInstancePropsMixin
type jsiiProxy_CfnVerifiedAccessInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVerifiedAccessInstancePropsMixin) Props() *CfnVerifiedAccessInstanceMixinProps {
	var returns *CfnVerifiedAccessInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVerifiedAccessInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VerifiedAccessInstance`.
func NewCfnVerifiedAccessInstancePropsMixin(props *CfnVerifiedAccessInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVerifiedAccessInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVerifiedAccessInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVerifiedAccessInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VerifiedAccessInstance`.
func NewCfnVerifiedAccessInstancePropsMixin_Override(c CfnVerifiedAccessInstancePropsMixin, props *CfnVerifiedAccessInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVerifiedAccessInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVerifiedAccessInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVerifiedAccessInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVerifiedAccessInstancePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVerifiedAccessInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

