package previewawsroute53mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53/previewawsroute53mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A complex type that contains an optional comment, the name and ID of the hosted zone that you want to make changes in, and values for the records that you want to create.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRecordSetGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnRecordSetGroupPropsMixin(&CfnRecordSetGroupMixinProps{
//   	Comment: jsii.String("comment"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	HostedZoneName: jsii.String("hostedZoneName"),
//   	RecordSets: []interface{}{
//   		&RecordSetProperty{
//   			AliasTarget: &AliasTargetProperty{
//   				DnsName: jsii.String("dnsName"),
//   				EvaluateTargetHealth: jsii.Boolean(false),
//   				HostedZoneId: jsii.String("hostedZoneId"),
//   			},
//   			CidrRoutingConfig: &CidrRoutingConfigProperty{
//   				CollectionId: jsii.String("collectionId"),
//   				LocationName: jsii.String("locationName"),
//   			},
//   			Failover: jsii.String("failover"),
//   			GeoLocation: &GeoLocationProperty{
//   				ContinentCode: jsii.String("continentCode"),
//   				CountryCode: jsii.String("countryCode"),
//   				SubdivisionCode: jsii.String("subdivisionCode"),
//   			},
//   			GeoProximityLocation: &GeoProximityLocationProperty{
//   				AwsRegion: jsii.String("awsRegion"),
//   				Bias: jsii.Number(123),
//   				Coordinates: &CoordinatesProperty{
//   					Latitude: jsii.String("latitude"),
//   					Longitude: jsii.String("longitude"),
//   				},
//   				LocalZoneGroup: jsii.String("localZoneGroup"),
//   			},
//   			HealthCheckId: jsii.String("healthCheckId"),
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   			HostedZoneName: jsii.String("hostedZoneName"),
//   			MultiValueAnswer: jsii.Boolean(false),
//   			Name: jsii.String("name"),
//   			Region: jsii.String("region"),
//   			ResourceRecords: []*string{
//   				jsii.String("resourceRecords"),
//   			},
//   			SetIdentifier: jsii.String("setIdentifier"),
//   			Ttl: jsii.String("ttl"),
//   			Type: jsii.String("type"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
//
type CfnRecordSetGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnRecordSetGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRecordSetGroupPropsMixin
type jsiiProxy_CfnRecordSetGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnRecordSetGroupPropsMixin) Props() *CfnRecordSetGroupMixinProps {
	var returns *CfnRecordSetGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRecordSetGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::RecordSetGroup`.
func NewCfnRecordSetGroupPropsMixin(props *CfnRecordSetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnRecordSetGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRecordSetGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRecordSetGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::RecordSetGroup`.
func NewCfnRecordSetGroupPropsMixin_Override(c CfnRecordSetGroupPropsMixin, props *CfnRecordSetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnRecordSetGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRecordSetGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRecordSetGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRecordSetGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRecordSetGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

