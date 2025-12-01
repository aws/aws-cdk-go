package previewawssagemakerevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents",
		reflect.TypeOf((*EndpointConfigEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "sageMakerEndpointConfigStateChangePattern", GoMethod: "SageMakerEndpointConfigStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_EndpointConfigEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents.SageMakerEndpointConfigStateChange",
		reflect.TypeOf((*EndpointConfigEvents_SageMakerEndpointConfigStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EndpointConfigEvents_SageMakerEndpointConfigStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents.SageMakerEndpointConfigStateChange.SageMakerEndpointConfigStateChangeItem",
		reflect.TypeOf((*EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents.SageMakerEndpointConfigStateChange.SageMakerEndpointConfigStateChangeProps",
		reflect.TypeOf((*EndpointConfigEvents_SageMakerEndpointConfigStateChange_SageMakerEndpointConfigStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.EndpointConfigEvents.SageMakerEndpointConfigStateChange.Tags",
		reflect.TypeOf((*EndpointConfigEvents_SageMakerEndpointConfigStateChange_Tags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents",
		reflect.TypeOf((*ModelEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "sageMakerTransformJobStateChangePattern", GoMethod: "SageMakerTransformJobStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ModelEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ModelEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.AlgorithmSpecification",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_AlgorithmSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.DataSource",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.DataSource1",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_DataSource1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.HyperParameters",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_HyperParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.OutputDataConfig",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_OutputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.PrimaryContainer",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_PrimaryContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.RequestParametersItem2",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_RequestParametersItem2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.ResourceConfig",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_ResourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.S3DataSource",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_S3DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.S3DataSource1",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_S3DataSource1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.StoppingCondition",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_StoppingCondition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.TransformInput",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_TransformInput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.TransformOutput",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_TransformOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.TransformResources",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_TransformResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*ModelEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ModelEvents_SageMakerTransformJobStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.DataSource",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.S3DataSource",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_S3DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.SageMakerTransformJobStateChangeProps",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_SageMakerTransformJobStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.Tags",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_Tags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.TransformInput",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_TransformInput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.TransformOutput",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_TransformOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.events.ModelEvents.SageMakerTransformJobStateChange.TransformResources",
		reflect.TypeOf((*ModelEvents_SageMakerTransformJobStateChange_TransformResources)(nil)).Elem(),
	)
}
