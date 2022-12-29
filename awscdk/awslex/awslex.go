package awslex

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lex.CfnBot",
		reflect.TypeOf((*CfnBot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "autoBuildBotLocales", GoGetter: "AutoBuildBotLocales"},
			_jsii_.MemberProperty{JsiiProperty: "botFileS3Location", GoGetter: "BotFileS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "botLocales", GoGetter: "BotLocales"},
			_jsii_.MemberProperty{JsiiProperty: "botTags", GoGetter: "BotTags"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataPrivacy", GoGetter: "DataPrivacy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "idleSessionTtlInSeconds", GoGetter: "IdleSessionTtlInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "testBotAliasSettings", GoGetter: "TestBotAliasSettings"},
			_jsii_.MemberProperty{JsiiProperty: "testBotAliasTags", GoGetter: "TestBotAliasTags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBot{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.AdvancedRecognitionSettingProperty",
		reflect.TypeOf((*CfnBot_AdvancedRecognitionSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.AllowedInputTypesProperty",
		reflect.TypeOf((*CfnBot_AllowedInputTypesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.AudioAndDTMFInputSpecificationProperty",
		reflect.TypeOf((*CfnBot_AudioAndDTMFInputSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.AudioLogDestinationProperty",
		reflect.TypeOf((*CfnBot_AudioLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.AudioLogSettingProperty",
		reflect.TypeOf((*CfnBot_AudioLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.AudioSpecificationProperty",
		reflect.TypeOf((*CfnBot_AudioSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.BotAliasLocaleSettingsItemProperty",
		reflect.TypeOf((*CfnBot_BotAliasLocaleSettingsItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.BotAliasLocaleSettingsProperty",
		reflect.TypeOf((*CfnBot_BotAliasLocaleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.BotLocaleProperty",
		reflect.TypeOf((*CfnBot_BotLocaleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.ButtonProperty",
		reflect.TypeOf((*CfnBot_ButtonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.CloudWatchLogGroupLogDestinationProperty",
		reflect.TypeOf((*CfnBot_CloudWatchLogGroupLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.CodeHookSpecificationProperty",
		reflect.TypeOf((*CfnBot_CodeHookSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.ConversationLogSettingsProperty",
		reflect.TypeOf((*CfnBot_ConversationLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.CustomPayloadProperty",
		reflect.TypeOf((*CfnBot_CustomPayloadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.CustomVocabularyItemProperty",
		reflect.TypeOf((*CfnBot_CustomVocabularyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.CustomVocabularyProperty",
		reflect.TypeOf((*CfnBot_CustomVocabularyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.DTMFSpecificationProperty",
		reflect.TypeOf((*CfnBot_DTMFSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.DataPrivacyProperty",
		reflect.TypeOf((*CfnBot_DataPrivacyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.DialogCodeHookSettingProperty",
		reflect.TypeOf((*CfnBot_DialogCodeHookSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.ExternalSourceSettingProperty",
		reflect.TypeOf((*CfnBot_ExternalSourceSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.FulfillmentCodeHookSettingProperty",
		reflect.TypeOf((*CfnBot_FulfillmentCodeHookSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.FulfillmentStartResponseSpecificationProperty",
		reflect.TypeOf((*CfnBot_FulfillmentStartResponseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.FulfillmentUpdateResponseSpecificationProperty",
		reflect.TypeOf((*CfnBot_FulfillmentUpdateResponseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.FulfillmentUpdatesSpecificationProperty",
		reflect.TypeOf((*CfnBot_FulfillmentUpdatesSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.GrammarSlotTypeSettingProperty",
		reflect.TypeOf((*CfnBot_GrammarSlotTypeSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.GrammarSlotTypeSourceProperty",
		reflect.TypeOf((*CfnBot_GrammarSlotTypeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.ImageResponseCardProperty",
		reflect.TypeOf((*CfnBot_ImageResponseCardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.InputContextProperty",
		reflect.TypeOf((*CfnBot_InputContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.IntentClosingSettingProperty",
		reflect.TypeOf((*CfnBot_IntentClosingSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.IntentConfirmationSettingProperty",
		reflect.TypeOf((*CfnBot_IntentConfirmationSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.IntentProperty",
		reflect.TypeOf((*CfnBot_IntentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.KendraConfigurationProperty",
		reflect.TypeOf((*CfnBot_KendraConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.LambdaCodeHookProperty",
		reflect.TypeOf((*CfnBot_LambdaCodeHookProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.MessageGroupProperty",
		reflect.TypeOf((*CfnBot_MessageGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.MessageProperty",
		reflect.TypeOf((*CfnBot_MessageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.MultipleValuesSettingProperty",
		reflect.TypeOf((*CfnBot_MultipleValuesSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.ObfuscationSettingProperty",
		reflect.TypeOf((*CfnBot_ObfuscationSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.OutputContextProperty",
		reflect.TypeOf((*CfnBot_OutputContextProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.PlainTextMessageProperty",
		reflect.TypeOf((*CfnBot_PlainTextMessageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.PostFulfillmentStatusSpecificationProperty",
		reflect.TypeOf((*CfnBot_PostFulfillmentStatusSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.PromptAttemptSpecificationProperty",
		reflect.TypeOf((*CfnBot_PromptAttemptSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.PromptSpecificationProperty",
		reflect.TypeOf((*CfnBot_PromptSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.ResponseSpecificationProperty",
		reflect.TypeOf((*CfnBot_ResponseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.S3BucketLogDestinationProperty",
		reflect.TypeOf((*CfnBot_S3BucketLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.S3LocationProperty",
		reflect.TypeOf((*CfnBot_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SSMLMessageProperty",
		reflect.TypeOf((*CfnBot_SSMLMessageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SampleUtteranceProperty",
		reflect.TypeOf((*CfnBot_SampleUtteranceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SampleValueProperty",
		reflect.TypeOf((*CfnBot_SampleValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SentimentAnalysisSettingsProperty",
		reflect.TypeOf((*CfnBot_SentimentAnalysisSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotDefaultValueProperty",
		reflect.TypeOf((*CfnBot_SlotDefaultValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotDefaultValueSpecificationProperty",
		reflect.TypeOf((*CfnBot_SlotDefaultValueSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotPriorityProperty",
		reflect.TypeOf((*CfnBot_SlotPriorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotProperty",
		reflect.TypeOf((*CfnBot_SlotProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotTypeProperty",
		reflect.TypeOf((*CfnBot_SlotTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotTypeValueProperty",
		reflect.TypeOf((*CfnBot_SlotTypeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotValueElicitationSettingProperty",
		reflect.TypeOf((*CfnBot_SlotValueElicitationSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotValueRegexFilterProperty",
		reflect.TypeOf((*CfnBot_SlotValueRegexFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.SlotValueSelectionSettingProperty",
		reflect.TypeOf((*CfnBot_SlotValueSelectionSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.StillWaitingResponseSpecificationProperty",
		reflect.TypeOf((*CfnBot_StillWaitingResponseSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.TestBotAliasSettingsProperty",
		reflect.TypeOf((*CfnBot_TestBotAliasSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.TextInputSpecificationProperty",
		reflect.TypeOf((*CfnBot_TextInputSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.TextLogDestinationProperty",
		reflect.TypeOf((*CfnBot_TextLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.TextLogSettingProperty",
		reflect.TypeOf((*CfnBot_TextLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.VoiceSettingsProperty",
		reflect.TypeOf((*CfnBot_VoiceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBot.WaitAndContinueSpecificationProperty",
		reflect.TypeOf((*CfnBot_WaitAndContinueSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lex.CfnBotAlias",
		reflect.TypeOf((*CfnBotAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrBotAliasId", GoGetter: "AttrBotAliasId"},
			_jsii_.MemberProperty{JsiiProperty: "attrBotAliasStatus", GoGetter: "AttrBotAliasStatus"},
			_jsii_.MemberProperty{JsiiProperty: "botAliasLocaleSettings", GoGetter: "BotAliasLocaleSettings"},
			_jsii_.MemberProperty{JsiiProperty: "botAliasName", GoGetter: "BotAliasName"},
			_jsii_.MemberProperty{JsiiProperty: "botAliasTags", GoGetter: "BotAliasTags"},
			_jsii_.MemberProperty{JsiiProperty: "botId", GoGetter: "BotId"},
			_jsii_.MemberProperty{JsiiProperty: "botVersion", GoGetter: "BotVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "conversationLogSettings", GoGetter: "ConversationLogSettings"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "sentimentAnalysisSettings", GoGetter: "SentimentAnalysisSettings"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBotAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.AudioLogDestinationProperty",
		reflect.TypeOf((*CfnBotAlias_AudioLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.AudioLogSettingProperty",
		reflect.TypeOf((*CfnBotAlias_AudioLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.BotAliasLocaleSettingsItemProperty",
		reflect.TypeOf((*CfnBotAlias_BotAliasLocaleSettingsItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.BotAliasLocaleSettingsProperty",
		reflect.TypeOf((*CfnBotAlias_BotAliasLocaleSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.CloudWatchLogGroupLogDestinationProperty",
		reflect.TypeOf((*CfnBotAlias_CloudWatchLogGroupLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.CodeHookSpecificationProperty",
		reflect.TypeOf((*CfnBotAlias_CodeHookSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.ConversationLogSettingsProperty",
		reflect.TypeOf((*CfnBotAlias_ConversationLogSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.LambdaCodeHookProperty",
		reflect.TypeOf((*CfnBotAlias_LambdaCodeHookProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.S3BucketLogDestinationProperty",
		reflect.TypeOf((*CfnBotAlias_S3BucketLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.SentimentAnalysisSettingsProperty",
		reflect.TypeOf((*CfnBotAlias_SentimentAnalysisSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.TextLogDestinationProperty",
		reflect.TypeOf((*CfnBotAlias_TextLogDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAlias.TextLogSettingProperty",
		reflect.TypeOf((*CfnBotAlias_TextLogSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotAliasProps",
		reflect.TypeOf((*CfnBotAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotProps",
		reflect.TypeOf((*CfnBotProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lex.CfnBotVersion",
		reflect.TypeOf((*CfnBotVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrBotVersion", GoGetter: "AttrBotVersion"},
			_jsii_.MemberProperty{JsiiProperty: "botId", GoGetter: "BotId"},
			_jsii_.MemberProperty{JsiiProperty: "botVersionLocaleSpecification", GoGetter: "BotVersionLocaleSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBotVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotVersion.BotVersionLocaleDetailsProperty",
		reflect.TypeOf((*CfnBotVersion_BotVersionLocaleDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotVersion.BotVersionLocaleSpecificationProperty",
		reflect.TypeOf((*CfnBotVersion_BotVersionLocaleSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnBotVersionProps",
		reflect.TypeOf((*CfnBotVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_lex.CfnResourcePolicy",
		reflect.TypeOf((*CfnResourcePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrRevisionId", GoGetter: "AttrRevisionId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArn", GoGetter: "ResourceArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourcePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_lex.CfnResourcePolicyProps",
		reflect.TypeOf((*CfnResourcePolicyProps)(nil)).Elem(),
	)
}
