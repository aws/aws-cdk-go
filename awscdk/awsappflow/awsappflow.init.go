package awsappflow

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_appflow.CfnConnectorProfile",
		reflect.TypeOf((*CfnConnectorProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrConnectorProfileArn", GoGetter: "AttrConnectorProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCredentialsArn", GoGetter: "AttrCredentialsArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionMode", GoGetter: "ConnectionMode"},
			_jsii_.MemberProperty{JsiiProperty: "connectorProfileConfig", GoGetter: "ConnectorProfileConfig"},
			_jsii_.MemberProperty{JsiiProperty: "connectorProfileName", GoGetter: "ConnectorProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "connectorType", GoGetter: "ConnectorType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsArn", GoGetter: "KmsArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectorProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.AmplitudeConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_AmplitudeConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ConnectorOAuthRequestProperty",
		reflect.TypeOf((*CfnConnectorProfile_ConnectorOAuthRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ConnectorProfileConfigProperty",
		reflect.TypeOf((*CfnConnectorProfile_ConnectorProfileConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_ConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_ConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.DatadogConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_DatadogConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.DatadogConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_DatadogConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.DynatraceConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_DynatraceConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.DynatraceConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_DynatraceConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.GoogleAnalyticsConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_GoogleAnalyticsConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.InforNexusConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_InforNexusConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.InforNexusConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_InforNexusConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.MarketoConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_MarketoConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.MarketoConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_MarketoConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.RedshiftConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_RedshiftConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.RedshiftConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_RedshiftConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SalesforceConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_SalesforceConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SalesforceConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_SalesforceConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ServiceNowConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_ServiceNowConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ServiceNowConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_ServiceNowConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SingularConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_SingularConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SlackConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_SlackConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SlackConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_SlackConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SnowflakeConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_SnowflakeConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.SnowflakeConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_SnowflakeConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.TrendmicroConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_TrendmicroConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.VeevaConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_VeevaConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.VeevaConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_VeevaConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ZendeskConnectorProfileCredentialsProperty",
		reflect.TypeOf((*CfnConnectorProfile_ZendeskConnectorProfileCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfile.ZendeskConnectorProfilePropertiesProperty",
		reflect.TypeOf((*CfnConnectorProfile_ZendeskConnectorProfilePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnConnectorProfileProps",
		reflect.TypeOf((*CfnConnectorProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_appflow.CfnFlow",
		reflect.TypeOf((*CfnFlow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrFlowArn", GoGetter: "AttrFlowArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "destinationFlowConfigList", GoGetter: "DestinationFlowConfigList"},
			_jsii_.MemberProperty{JsiiProperty: "flowName", GoGetter: "FlowName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kmsArn", GoGetter: "KmsArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "sourceFlowConfig", GoGetter: "SourceFlowConfig"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tasks", GoGetter: "Tasks"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfig", GoGetter: "TriggerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.AggregationConfigProperty",
		reflect.TypeOf((*CfnFlow_AggregationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.AmplitudeSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_AmplitudeSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.ConnectorOperatorProperty",
		reflect.TypeOf((*CfnFlow_ConnectorOperatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.DatadogSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_DatadogSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.DestinationConnectorPropertiesProperty",
		reflect.TypeOf((*CfnFlow_DestinationConnectorPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.DestinationFlowConfigProperty",
		reflect.TypeOf((*CfnFlow_DestinationFlowConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.DynatraceSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_DynatraceSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.ErrorHandlingConfigProperty",
		reflect.TypeOf((*CfnFlow_ErrorHandlingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.EventBridgeDestinationPropertiesProperty",
		reflect.TypeOf((*CfnFlow_EventBridgeDestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.GoogleAnalyticsSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_GoogleAnalyticsSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.IdFieldNamesListProperty",
		reflect.TypeOf((*CfnFlow_IdFieldNamesListProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.IncrementalPullConfigProperty",
		reflect.TypeOf((*CfnFlow_IncrementalPullConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.InforNexusSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_InforNexusSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.MarketoSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_MarketoSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.PrefixConfigProperty",
		reflect.TypeOf((*CfnFlow_PrefixConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.RedshiftDestinationPropertiesProperty",
		reflect.TypeOf((*CfnFlow_RedshiftDestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.S3DestinationPropertiesProperty",
		reflect.TypeOf((*CfnFlow_S3DestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.S3OutputFormatConfigProperty",
		reflect.TypeOf((*CfnFlow_S3OutputFormatConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.S3SourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_S3SourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SalesforceDestinationPropertiesProperty",
		reflect.TypeOf((*CfnFlow_SalesforceDestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SalesforceSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_SalesforceSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.ScheduledTriggerPropertiesProperty",
		reflect.TypeOf((*CfnFlow_ScheduledTriggerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.ServiceNowSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_ServiceNowSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SingularSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_SingularSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SlackSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_SlackSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SnowflakeDestinationPropertiesProperty",
		reflect.TypeOf((*CfnFlow_SnowflakeDestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SourceConnectorPropertiesProperty",
		reflect.TypeOf((*CfnFlow_SourceConnectorPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.SourceFlowConfigProperty",
		reflect.TypeOf((*CfnFlow_SourceFlowConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.TaskPropertiesObjectProperty",
		reflect.TypeOf((*CfnFlow_TaskPropertiesObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.TaskProperty",
		reflect.TypeOf((*CfnFlow_TaskProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.TrendmicroSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_TrendmicroSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.TriggerConfigProperty",
		reflect.TypeOf((*CfnFlow_TriggerConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.UpsolverDestinationPropertiesProperty",
		reflect.TypeOf((*CfnFlow_UpsolverDestinationPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.UpsolverS3OutputFormatConfigProperty",
		reflect.TypeOf((*CfnFlow_UpsolverS3OutputFormatConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.VeevaSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_VeevaSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlow.ZendeskSourcePropertiesProperty",
		reflect.TypeOf((*CfnFlow_ZendeskSourcePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_appflow.CfnFlowProps",
		reflect.TypeOf((*CfnFlowProps)(nil)).Elem(),
	)
}
