package awspinpointemail

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet",
		reflect.TypeOf((*CfnConfigurationSet)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationSetRef", GoGetter: "ConfigurationSetRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryOptions", GoGetter: "DeliveryOptions"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
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
			_jsii_.MemberProperty{JsiiProperty: "reputationOptions", GoGetter: "ReputationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "sendingOptions", GoGetter: "SendingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trackingOptions", GoGetter: "TrackingOptions"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConfigurationSet{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConfigurationSetRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet.DeliveryOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSet_DeliveryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet.ReputationOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSet_ReputationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet.SendingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSet_SendingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSet.TrackingOptionsProperty",
		reflect.TypeOf((*CfnConfigurationSet_TrackingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination",
		reflect.TypeOf((*CfnConfigurationSetEventDestination)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "configurationSetEventDestinationRef", GoGetter: "ConfigurationSetEventDestinationRef"},
			_jsii_.MemberProperty{JsiiProperty: "configurationSetName", GoGetter: "ConfigurationSetName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventDestination", GoGetter: "EventDestination"},
			_jsii_.MemberProperty{JsiiProperty: "eventDestinationName", GoGetter: "EventDestinationName"},
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
			j := jsiiProxy_CfnConfigurationSetEventDestination{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IConfigurationSetEventDestinationRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination.CloudWatchDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestination_CloudWatchDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination.DimensionConfigurationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestination_DimensionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination.EventDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestination_EventDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination.KinesisFirehoseDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestination_KinesisFirehoseDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination.PinpointDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestination_PinpointDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestination.SnsDestinationProperty",
		reflect.TypeOf((*CfnConfigurationSetEventDestination_SnsDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetEventDestinationProps",
		reflect.TypeOf((*CfnConfigurationSetEventDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnConfigurationSetProps",
		reflect.TypeOf((*CfnConfigurationSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPool",
		reflect.TypeOf((*CfnDedicatedIpPool)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dedicatedIpPoolRef", GoGetter: "DedicatedIpPoolRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "poolName", GoGetter: "PoolName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDedicatedIpPool{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDedicatedIpPoolRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnDedicatedIpPoolProps",
		reflect.TypeOf((*CfnDedicatedIpPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity",
		reflect.TypeOf((*CfnIdentity)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityDnsRecordName1", GoGetter: "AttrIdentityDnsRecordName1"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityDnsRecordName2", GoGetter: "AttrIdentityDnsRecordName2"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityDnsRecordName3", GoGetter: "AttrIdentityDnsRecordName3"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityDnsRecordValue1", GoGetter: "AttrIdentityDnsRecordValue1"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityDnsRecordValue2", GoGetter: "AttrIdentityDnsRecordValue2"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentityDnsRecordValue3", GoGetter: "AttrIdentityDnsRecordValue3"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dkimSigningEnabled", GoGetter: "DkimSigningEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "feedbackForwardingEnabled", GoGetter: "FeedbackForwardingEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "identityRef", GoGetter: "IdentityRef"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mailFromAttributes", GoGetter: "MailFromAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
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
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IIdentityRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentity.MailFromAttributesProperty",
		reflect.TypeOf((*CfnIdentity_MailFromAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.CfnIdentityProps",
		reflect.TypeOf((*CfnIdentityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.ConfigurationSetEventDestinationReference",
		reflect.TypeOf((*ConfigurationSetEventDestinationReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.ConfigurationSetReference",
		reflect.TypeOf((*ConfigurationSetReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.DedicatedIpPoolReference",
		reflect.TypeOf((*DedicatedIpPoolReference)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_pinpointemail.IConfigurationSetEventDestinationRef",
		reflect.TypeOf((*IConfigurationSetEventDestinationRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigurationSetEventDestinationRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_pinpointemail.IConfigurationSetRef",
		reflect.TypeOf((*IConfigurationSetRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IConfigurationSetRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_pinpointemail.IDedicatedIpPoolRef",
		reflect.TypeOf((*IDedicatedIpPoolRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDedicatedIpPoolRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_pinpointemail.IIdentityRef",
		reflect.TypeOf((*IIdentityRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IIdentityRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pinpointemail.IdentityReference",
		reflect.TypeOf((*IdentityReference)(nil)).Elem(),
	)
}
