package previewawss3mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.AutoDeleteObjects",
		reflect.TypeOf((*AutoDeleteObjects)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_AutoDeleteObjects{}
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantMixinProps",
		reflect.TypeOf((*CfnAccessGrantMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantPropsMixin",
		reflect.TypeOf((*CfnAccessGrantPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessGrantPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantPropsMixin.AccessGrantsLocationConfigurationProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_AccessGrantsLocationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantPropsMixin.GranteeProperty",
		reflect.TypeOf((*CfnAccessGrantPropsMixin_GranteeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsInstanceMixinProps",
		reflect.TypeOf((*CfnAccessGrantsInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsInstancePropsMixin",
		reflect.TypeOf((*CfnAccessGrantsInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessGrantsInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsLocationMixinProps",
		reflect.TypeOf((*CfnAccessGrantsLocationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsLocationPropsMixin",
		reflect.TypeOf((*CfnAccessGrantsLocationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessGrantsLocationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessPointMixinProps",
		reflect.TypeOf((*CfnAccessPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessPointPropsMixin",
		reflect.TypeOf((*CfnAccessPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessPointPropsMixin.PublicAccessBlockConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_PublicAccessBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessPointPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketMixinProps",
		reflect.TypeOf((*CfnBucketMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPolicyMixinProps",
		reflect.TypeOf((*CfnBucketPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPolicyPropsMixin",
		reflect.TypeOf((*CfnBucketPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBucketPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin",
		reflect.TypeOf((*CfnBucketPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBucketPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.AbortIncompleteMultipartUploadProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_AbortIncompleteMultipartUploadProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.AccelerateConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_AccelerateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.AccessControlTranslationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_AccessControlTranslationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.AnalyticsConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_AnalyticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.BlockedEncryptionTypesProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_BlockedEncryptionTypesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.BucketEncryptionProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_BucketEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.CorsConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_CorsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.CorsRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_CorsRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.DataExportProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_DataExportProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.DefaultRetentionProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_DefaultRetentionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.DeleteMarkerReplicationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_DeleteMarkerReplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.EventBridgeConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_EventBridgeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.FilterRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_FilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.IntelligentTieringConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_IntelligentTieringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.InventoryConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_InventoryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.InventoryTableConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_InventoryTableConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.JournalTableConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_JournalTableConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.LambdaConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_LambdaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.LoggingConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_LoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.MetadataConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_MetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.MetadataDestinationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_MetadataDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.MetadataTableConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_MetadataTableConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.MetadataTableEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_MetadataTableEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.MetricsConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_MetricsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.MetricsProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_MetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.NoncurrentVersionExpirationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_NoncurrentVersionExpirationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.NoncurrentVersionTransitionProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_NoncurrentVersionTransitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.NotificationConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_NotificationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.NotificationFilterProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_NotificationFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ObjectLockConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ObjectLockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ObjectLockRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ObjectLockRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.OwnershipControlsProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_OwnershipControlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.OwnershipControlsRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_OwnershipControlsRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.PartitionedPrefixProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_PartitionedPrefixProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.PublicAccessBlockConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_PublicAccessBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.QueueConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_QueueConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.RecordExpirationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_RecordExpirationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.RedirectAllRequestsToProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_RedirectAllRequestsToProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.RedirectRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_RedirectRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicaModificationsProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicaModificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationDestinationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationRuleAndOperatorProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationRuleAndOperatorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationRuleFilterProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationRuleFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationTimeProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationTimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ReplicationTimeValueProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ReplicationTimeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.RoutingRuleConditionProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_RoutingRuleConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.RoutingRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_RoutingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.S3KeyFilterProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_S3KeyFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.S3TablesDestinationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_S3TablesDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ServerSideEncryptionByDefaultProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ServerSideEncryptionByDefaultProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.ServerSideEncryptionRuleProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_ServerSideEncryptionRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.SourceSelectionCriteriaProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_SourceSelectionCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.SseKmsEncryptedObjectsProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_SseKmsEncryptedObjectsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.StorageClassAnalysisProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_StorageClassAnalysisProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.TagFilterProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.TargetObjectKeyFormatProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_TargetObjectKeyFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.TieringProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_TieringProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.TopicConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_TopicConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.TransitionProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_TransitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.VersioningConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_VersioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnBucketPropsMixin.WebsiteConfigurationProperty",
		reflect.TypeOf((*CfnBucketPropsMixin_WebsiteConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointMixinProps",
		reflect.TypeOf((*CfnMultiRegionAccessPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPolicyMixinProps",
		reflect.TypeOf((*CfnMultiRegionAccessPointPolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPolicyPropsMixin",
		reflect.TypeOf((*CfnMultiRegionAccessPointPolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMultiRegionAccessPointPolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPolicyPropsMixin.PolicyStatusProperty",
		reflect.TypeOf((*CfnMultiRegionAccessPointPolicyPropsMixin_PolicyStatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin",
		reflect.TypeOf((*CfnMultiRegionAccessPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMultiRegionAccessPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin.PublicAccessBlockConfigurationProperty",
		reflect.TypeOf((*CfnMultiRegionAccessPointPropsMixin_PublicAccessBlockConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnMultiRegionAccessPointPropsMixin.RegionProperty",
		reflect.TypeOf((*CfnMultiRegionAccessPointPropsMixin_RegionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupMixinProps",
		reflect.TypeOf((*CfnStorageLensGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin",
		reflect.TypeOf((*CfnStorageLensGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStorageLensGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin.AndProperty",
		reflect.TypeOf((*CfnStorageLensGroupPropsMixin_AndProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnStorageLensGroupPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin.MatchObjectAgeProperty",
		reflect.TypeOf((*CfnStorageLensGroupPropsMixin_MatchObjectAgeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin.MatchObjectSizeProperty",
		reflect.TypeOf((*CfnStorageLensGroupPropsMixin_MatchObjectSizeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensGroupPropsMixin.OrProperty",
		reflect.TypeOf((*CfnStorageLensGroupPropsMixin_OrProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensMixinProps",
		reflect.TypeOf((*CfnStorageLensMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin",
		reflect.TypeOf((*CfnStorageLensPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStorageLensPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.AccountLevelProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_AccountLevelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.ActivityMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_ActivityMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.AdvancedCostOptimizationMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_AdvancedCostOptimizationMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.AdvancedDataProtectionMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_AdvancedDataProtectionMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.AdvancedPerformanceMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_AdvancedPerformanceMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.AwsOrgProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_AwsOrgProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.BucketLevelProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_BucketLevelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.BucketsAndRegionsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_BucketsAndRegionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.CloudWatchMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_CloudWatchMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.DataExportProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_DataExportProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.DetailedStatusCodesMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_DetailedStatusCodesMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.EncryptionProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_EncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.PrefixLevelProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_PrefixLevelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.PrefixLevelStorageMetricsProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_PrefixLevelStorageMetricsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.S3BucketDestinationProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_S3BucketDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.SSEKMSProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_SSEKMSProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.SelectionCriteriaProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_SelectionCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.StorageLensConfigurationProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_StorageLensConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.StorageLensExpandedPrefixesDataExportProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_StorageLensExpandedPrefixesDataExportProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.StorageLensGroupLevelProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_StorageLensGroupLevelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.StorageLensGroupSelectionCriteriaProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_StorageLensGroupSelectionCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnStorageLensPropsMixin.StorageLensTableDestinationProperty",
		reflect.TypeOf((*CfnStorageLensPropsMixin_StorageLensTableDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.mixins.EnableVersioning",
		reflect.TypeOf((*EnableVersioning)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_EnableVersioning{}
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
