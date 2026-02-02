package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::SageMaker::EndpointConfig` resource creates a configuration for an Amazon SageMaker endpoint.
//
// For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) in the *SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEndpointConfigPropsMixin := awscdkmixinspreview.Mixins.NewCfnEndpointConfigPropsMixin(&CfnEndpointConfigMixinProps{
//   	AsyncInferenceConfig: &AsyncInferenceConfigProperty{
//   		ClientConfig: &AsyncInferenceClientConfigProperty{
//   			MaxConcurrentInvocationsPerInstance: jsii.Number(123),
//   		},
//   		OutputConfig: &AsyncInferenceOutputConfigProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			NotificationConfig: &AsyncInferenceNotificationConfigProperty{
//   				ErrorTopic: jsii.String("errorTopic"),
//   				IncludeInferenceResponseIn: []*string{
//   					jsii.String("includeInferenceResponseIn"),
//   				},
//   				SuccessTopic: jsii.String("successTopic"),
//   			},
//   			S3FailurePath: jsii.String("s3FailurePath"),
//   			S3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   	DataCaptureConfig: &DataCaptureConfigProperty{
//   		CaptureContentTypeHeader: &CaptureContentTypeHeaderProperty{
//   			CsvContentTypes: []*string{
//   				jsii.String("csvContentTypes"),
//   			},
//   			JsonContentTypes: []*string{
//   				jsii.String("jsonContentTypes"),
//   			},
//   		},
//   		CaptureOptions: []interface{}{
//   			&CaptureOptionProperty{
//   				CaptureMode: jsii.String("captureMode"),
//   			},
//   		},
//   		DestinationS3Uri: jsii.String("destinationS3Uri"),
//   		EnableCapture: jsii.Boolean(false),
//   		InitialSamplingPercentage: jsii.Number(123),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EnableNetworkIsolation: jsii.Boolean(false),
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	ExplainerConfig: &ExplainerConfigProperty{
//   		ClarifyExplainerConfig: &ClarifyExplainerConfigProperty{
//   			EnableExplanations: jsii.String("enableExplanations"),
//   			InferenceConfig: &ClarifyInferenceConfigProperty{
//   				ContentTemplate: jsii.String("contentTemplate"),
//   				FeatureHeaders: []*string{
//   					jsii.String("featureHeaders"),
//   				},
//   				FeaturesAttribute: jsii.String("featuresAttribute"),
//   				FeatureTypes: []*string{
//   					jsii.String("featureTypes"),
//   				},
//   				LabelAttribute: jsii.String("labelAttribute"),
//   				LabelHeaders: []*string{
//   					jsii.String("labelHeaders"),
//   				},
//   				LabelIndex: jsii.Number(123),
//   				MaxPayloadInMb: jsii.Number(123),
//   				MaxRecordCount: jsii.Number(123),
//   				ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   				ProbabilityIndex: jsii.Number(123),
//   			},
//   			ShapConfig: &ClarifyShapConfigProperty{
//   				NumberOfSamples: jsii.Number(123),
//   				Seed: jsii.Number(123),
//   				ShapBaselineConfig: &ClarifyShapBaselineConfigProperty{
//   					MimeType: jsii.String("mimeType"),
//   					ShapBaseline: jsii.String("shapBaseline"),
//   					ShapBaselineUri: jsii.String("shapBaselineUri"),
//   				},
//   				TextConfig: &ClarifyTextConfigProperty{
//   					Granularity: jsii.String("granularity"),
//   					Language: jsii.String("language"),
//   				},
//   				UseLogit: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ProductionVariants: []interface{}{
//   		&ProductionVariantProperty{
//   			AcceleratorType: jsii.String("acceleratorType"),
//   			CapacityReservationConfig: &CapacityReservationConfigProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				MlReservationArn: jsii.String("mlReservationArn"),
//   			},
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			EnableSsmAccess: jsii.Boolean(false),
//   			InferenceAmiVersion: jsii.String("inferenceAmiVersion"),
//   			InitialInstanceCount: jsii.Number(123),
//   			InitialVariantWeight: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   				MaxInstanceCount: jsii.Number(123),
//   				MinInstanceCount: jsii.Number(123),
//   				Status: jsii.String("status"),
//   			},
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			ModelName: jsii.String("modelName"),
//   			RoutingConfig: &RoutingConfigProperty{
//   				RoutingStrategy: jsii.String("routingStrategy"),
//   			},
//   			ServerlessConfig: &ServerlessConfigProperty{
//   				MaxConcurrency: jsii.Number(123),
//   				MemorySizeInMb: jsii.Number(123),
//   				ProvisionedConcurrency: jsii.Number(123),
//   			},
//   			VariantName: jsii.String("variantName"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	ShadowProductionVariants: []interface{}{
//   		&ProductionVariantProperty{
//   			AcceleratorType: jsii.String("acceleratorType"),
//   			CapacityReservationConfig: &CapacityReservationConfigProperty{
//   				CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   				MlReservationArn: jsii.String("mlReservationArn"),
//   			},
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			EnableSsmAccess: jsii.Boolean(false),
//   			InferenceAmiVersion: jsii.String("inferenceAmiVersion"),
//   			InitialInstanceCount: jsii.Number(123),
//   			InitialVariantWeight: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			ManagedInstanceScaling: &ManagedInstanceScalingProperty{
//   				MaxInstanceCount: jsii.Number(123),
//   				MinInstanceCount: jsii.Number(123),
//   				Status: jsii.String("status"),
//   			},
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   			ModelName: jsii.String("modelName"),
//   			RoutingConfig: &RoutingConfigProperty{
//   				RoutingStrategy: jsii.String("routingStrategy"),
//   			},
//   			ServerlessConfig: &ServerlessConfigProperty{
//   				MaxConcurrency: jsii.Number(123),
//   				MemorySizeInMb: jsii.Number(123),
//   				ProvisionedConcurrency: jsii.Number(123),
//   			},
//   			VariantName: jsii.String("variantName"),
//   			VolumeSizeInGb: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html
//
type CfnEndpointConfigPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEndpointConfigMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEndpointConfigPropsMixin
type jsiiProxy_CfnEndpointConfigPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEndpointConfigPropsMixin) Props() *CfnEndpointConfigMixinProps {
	var returns *CfnEndpointConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointConfigPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::EndpointConfig`.
func NewCfnEndpointConfigPropsMixin(props *CfnEndpointConfigMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEndpointConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEndpointConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::EndpointConfig`.
func NewCfnEndpointConfigPropsMixin_Override(c CfnEndpointConfigPropsMixin, props *CfnEndpointConfigMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEndpointConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnEndpointConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEndpointConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

