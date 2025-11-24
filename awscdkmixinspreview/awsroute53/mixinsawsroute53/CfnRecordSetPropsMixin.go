package mixinsawsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsroute53/mixinsawsroute53/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Information about the record that you want to create.
//
// The `AWS::Route53::RecordSet` type can be used as a standalone resource or as an embedded property in the `AWS::Route53::RecordSetGroup` type. Note that some `AWS::Route53::RecordSet` properties are valid only when used within `AWS::Route53::RecordSetGroup` .
//
// For more information, see [ChangeResourceRecordSets](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) in the *Amazon Route 53 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRecordSetPropsMixin := awscdkmixinspreview.Mixins.NewCfnRecordSetPropsMixin(&CfnRecordSetMixinProps{
//   	AliasTarget: &AliasTargetProperty{
//   		DnsName: jsii.String("dnsName"),
//   		EvaluateTargetHealth: jsii.Boolean(false),
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   	},
//   	CidrRoutingConfig: &CidrRoutingConfigProperty{
//   		CollectionId: jsii.String("collectionId"),
//   		LocationName: jsii.String("locationName"),
//   	},
//   	Comment: jsii.String("comment"),
//   	Failover: jsii.String("failover"),
//   	GeoLocation: &GeoLocationProperty{
//   		ContinentCode: jsii.String("continentCode"),
//   		CountryCode: jsii.String("countryCode"),
//   		SubdivisionCode: jsii.String("subdivisionCode"),
//   	},
//   	GeoProximityLocation: &GeoProximityLocationProperty{
//   		AwsRegion: jsii.String("awsRegion"),
//   		Bias: jsii.Number(123),
//   		Coordinates: &CoordinatesProperty{
//   			Latitude: jsii.String("latitude"),
//   			Longitude: jsii.String("longitude"),
//   		},
//   		LocalZoneGroup: jsii.String("localZoneGroup"),
//   	},
//   	HealthCheckId: jsii.String("healthCheckId"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	HostedZoneName: jsii.String("hostedZoneName"),
//   	MultiValueAnswer: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	Region: jsii.String("region"),
//   	ResourceRecords: []*string{
//   		jsii.String("resourceRecords"),
//   	},
//   	SetIdentifier: jsii.String("setIdentifier"),
//   	Ttl: jsii.String("ttl"),
//   	Type: jsii.String("type"),
//   	Weight: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordset.html
//
type CfnRecordSetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRecordSetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRecordSetPropsMixin
type jsiiProxy_CfnRecordSetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRecordSetPropsMixin) Props() *CfnRecordSetMixinProps {
	var returns *CfnRecordSetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::RecordSet`.
func NewCfnRecordSetPropsMixin(props *CfnRecordSetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRecordSetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRecordSetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecordSetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::RecordSet`.
func NewCfnRecordSetPropsMixin_Override(c CfnRecordSetPropsMixin, props *CfnRecordSetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRecordSetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordSetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecordSetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnRecordSetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

