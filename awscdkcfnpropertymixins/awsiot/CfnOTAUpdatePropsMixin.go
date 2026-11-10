package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an IoT OTA update on a target group of things or groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnOTAUpdatePropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnOTAUpdatePropsMixin(&CfnOTAUpdateMixinProps{
//   	AdditionalParameters: map[string]*string{
//   		"additionalParametersKey": jsii.String("additionalParameters"),
//   	},
//   	AwsJobAbortConfig: &AwsJobAbortConfigProperty{
//   		AbortCriteriaList: []interface{}{
//   			&AwsJobAbortCriteriaProperty{
//   				Action: jsii.String("action"),
//   				FailureType: jsii.String("failureType"),
//   				MinNumberOfExecutedThings: jsii.Number(123),
//   				ThresholdPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	AwsJobExecutionsRolloutConfig: &AwsJobExecutionsRolloutConfigProperty{
//   		ExponentialRate: &AwsJobExponentialRolloutRateProperty{
//   			BaseRatePerMinute: jsii.Number(123),
//   			IncrementFactor: jsii.Number(123),
//   			RateIncreaseCriteria: &AwsJobRateIncreaseCriteriaProperty{
//   				NumberOfNotifiedThings: jsii.Number(123),
//   				NumberOfSucceededThings: jsii.Number(123),
//   			},
//   		},
//   		MaximumPerMinute: jsii.Number(123),
//   	},
//   	AwsJobPresignedUrlConfig: &AwsJobPresignedUrlConfigProperty{
//   		ExpiresInSec: jsii.Number(123),
//   	},
//   	AwsJobTimeoutConfig: &AwsJobTimeoutConfigProperty{
//   		InProgressTimeoutInMinutes: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Files: []interface{}{
//   		&OTAUpdateFileProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			CodeSigning: &CodeSigningProperty{
//   				AwsSignerJobId: jsii.String("awsSignerJobId"),
//   				CustomCodeSigning: &CustomCodeSigningProperty{
//   					CertificateChain: &CodeSigningCertificateChainProperty{
//   						CertificateName: jsii.String("certificateName"),
//   						InlineDocument: jsii.String("inlineDocument"),
//   					},
//   					HashAlgorithm: jsii.String("hashAlgorithm"),
//   					Signature: &CodeSigningSignatureProperty{
//   						InlineDocument: jsii.String("inlineDocument"),
//   					},
//   					SignatureAlgorithm: jsii.String("signatureAlgorithm"),
//   				},
//   				StartSigningJobParameter: &StartSigningJobParameterProperty{
//   					Destination: &DestinationProperty{
//   						S3Destination: &S3DestinationProperty{
//   							Bucket: jsii.String("bucket"),
//   							Prefix: jsii.String("prefix"),
//   						},
//   					},
//   					SigningProfileName: jsii.String("signingProfileName"),
//   					SigningProfileParameter: &SigningProfileParameterProperty{
//   						CertificateArn: jsii.String("certificateArn"),
//   						CertificatePathOnDevice: jsii.String("certificatePathOnDevice"),
//   						Platform: jsii.String("platform"),
//   					},
//   				},
//   			},
//   			FileLocation: &FileLocationProperty{
//   				S3Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					Key: jsii.String("key"),
//   					Version: jsii.String("version"),
//   				},
//   				Stream: &StreamProperty{
//   					FileId: jsii.Number(123),
//   					StreamId: jsii.String("streamId"),
//   				},
//   			},
//   			FileName: jsii.String("fileName"),
//   			FileType: jsii.Number(123),
//   			FileVersion: jsii.String("fileVersion"),
//   		},
//   	},
//   	OtaUpdateId: jsii.String("otaUpdateId"),
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//   	TargetSelection: jsii.String("targetSelection"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-otaupdate.html
//
type CfnOTAUpdatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnOTAUpdateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOTAUpdatePropsMixin
type jsiiProxy_CfnOTAUpdatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnOTAUpdatePropsMixin) Props() *CfnOTAUpdateMixinProps {
	var returns *CfnOTAUpdateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOTAUpdatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::OTAUpdate`.
func NewCfnOTAUpdatePropsMixin(props *CfnOTAUpdateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnOTAUpdatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOTAUpdatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOTAUpdatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnOTAUpdatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::OTAUpdate`.
func NewCfnOTAUpdatePropsMixin_Override(c CfnOTAUpdatePropsMixin, props *CfnOTAUpdateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnOTAUpdatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnOTAUpdatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOTAUpdatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnOTAUpdatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOTAUpdatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnOTAUpdatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOTAUpdatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOTAUpdatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

