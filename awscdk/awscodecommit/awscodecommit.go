package awscodecommit

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codecommit.CfnRepository",
		reflect.TypeOf((*CfnRepository)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "attrCloneUrlHttp", GoGetter: "AttrCloneUrlHttp"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloneUrlSsh", GoGetter: "AttrCloneUrlSsh"},
			_jsii_.MemberProperty{JsiiProperty: "attrName", GoGetter: "AttrName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "repositoryDescription", GoGetter: "RepositoryDescription"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryName", GoGetter: "RepositoryName"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRepository{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.CfnRepository.CodeProperty",
		reflect.TypeOf((*CfnRepository_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.CfnRepository.RepositoryTriggerProperty",
		reflect.TypeOf((*CfnRepository_RepositoryTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.CfnRepository.S3Property",
		reflect.TypeOf((*CfnRepository_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.CfnRepositoryProps",
		reflect.TypeOf((*CfnRepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codecommit.Code",
		reflect.TypeOf((*Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Code{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_codecommit.IRepository",
		reflect.TypeOf((*IRepository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleSource", GoMethod: "BindAsNotificationRuleSource"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantPull", GoMethod: "GrantPull"},
			_jsii_.MemberMethod{JsiiMethod: "grantPullPush", GoMethod: "GrantPullPush"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOn", GoMethod: "NotifyOn"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnApprovalRuleOverridden", GoMethod: "NotifyOnApprovalRuleOverridden"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnApprovalStatusChanged", GoMethod: "NotifyOnApprovalStatusChanged"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnBranchOrTagCreated", GoMethod: "NotifyOnBranchOrTagCreated"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnBranchOrTagDeleted", GoMethod: "NotifyOnBranchOrTagDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnPullRequestComment", GoMethod: "NotifyOnPullRequestComment"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnPullRequestCreated", GoMethod: "NotifyOnPullRequestCreated"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnPullRequestMerged", GoMethod: "NotifyOnPullRequestMerged"},
			_jsii_.MemberMethod{JsiiMethod: "onCommentOnCommit", GoMethod: "OnCommentOnCommit"},
			_jsii_.MemberMethod{JsiiMethod: "onCommentOnPullRequest", GoMethod: "OnCommentOnPullRequest"},
			_jsii_.MemberMethod{JsiiMethod: "onCommit", GoMethod: "OnCommit"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onPullRequestStateChange", GoMethod: "OnPullRequestStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onReferenceCreated", GoMethod: "OnReferenceCreated"},
			_jsii_.MemberMethod{JsiiMethod: "onReferenceDeleted", GoMethod: "OnReferenceDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "onReferenceUpdated", GoMethod: "OnReferenceUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryArn", GoGetter: "RepositoryArn"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCloneUrlGrc", GoGetter: "RepositoryCloneUrlGrc"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCloneUrlHttp", GoGetter: "RepositoryCloneUrlHttp"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCloneUrlSsh", GoGetter: "RepositoryCloneUrlSsh"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryName", GoGetter: "RepositoryName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IRepository{}
			_jsii_.InitJsiiProxy(&j.Type__awscodestarnotificationsINotificationRuleSource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.OnCommitOptions",
		reflect.TypeOf((*OnCommitOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codecommit.ReferenceEvent",
		reflect.TypeOf((*ReferenceEvent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ReferenceEvent{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_codecommit.Repository",
		reflect.TypeOf((*Repository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bindAsNotificationRuleSource", GoMethod: "BindAsNotificationRuleSource"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantPull", GoMethod: "GrantPull"},
			_jsii_.MemberMethod{JsiiMethod: "grantPullPush", GoMethod: "GrantPullPush"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "notifiyOnPullRequestMerged", GoMethod: "NotifiyOnPullRequestMerged"},
			_jsii_.MemberMethod{JsiiMethod: "notify", GoMethod: "Notify"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOn", GoMethod: "NotifyOn"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnApprovalRuleOverridden", GoMethod: "NotifyOnApprovalRuleOverridden"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnApprovalStatusChanged", GoMethod: "NotifyOnApprovalStatusChanged"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnBranchOrTagCreated", GoMethod: "NotifyOnBranchOrTagCreated"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnBranchOrTagDeleted", GoMethod: "NotifyOnBranchOrTagDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnPullRequestComment", GoMethod: "NotifyOnPullRequestComment"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnPullRequestCreated", GoMethod: "NotifyOnPullRequestCreated"},
			_jsii_.MemberMethod{JsiiMethod: "notifyOnPullRequestMerged", GoMethod: "NotifyOnPullRequestMerged"},
			_jsii_.MemberMethod{JsiiMethod: "onCommentOnCommit", GoMethod: "OnCommentOnCommit"},
			_jsii_.MemberMethod{JsiiMethod: "onCommentOnPullRequest", GoMethod: "OnCommentOnPullRequest"},
			_jsii_.MemberMethod{JsiiMethod: "onCommit", GoMethod: "OnCommit"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onPullRequestStateChange", GoMethod: "OnPullRequestStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onReferenceCreated", GoMethod: "OnReferenceCreated"},
			_jsii_.MemberMethod{JsiiMethod: "onReferenceDeleted", GoMethod: "OnReferenceDeleted"},
			_jsii_.MemberMethod{JsiiMethod: "onReferenceUpdated", GoMethod: "OnReferenceUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryArn", GoGetter: "RepositoryArn"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCloneUrlGrc", GoGetter: "RepositoryCloneUrlGrc"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCloneUrlHttp", GoGetter: "RepositoryCloneUrlHttp"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCloneUrlSsh", GoGetter: "RepositoryCloneUrlSsh"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryName", GoGetter: "RepositoryName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Repository{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRepository)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codecommit.RepositoryEventTrigger",
		reflect.TypeOf((*RepositoryEventTrigger)(nil)).Elem(),
		map[string]interface{}{
			"ALL": RepositoryEventTrigger_ALL,
			"UPDATE_REF": RepositoryEventTrigger_UPDATE_REF,
			"CREATE_REF": RepositoryEventTrigger_CREATE_REF,
			"DELETE_REF": RepositoryEventTrigger_DELETE_REF,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_codecommit.RepositoryNotificationEvents",
		reflect.TypeOf((*RepositoryNotificationEvents)(nil)).Elem(),
		map[string]interface{}{
			"COMMIT_COMMENT": RepositoryNotificationEvents_COMMIT_COMMENT,
			"PULL_REQUEST_COMMENT": RepositoryNotificationEvents_PULL_REQUEST_COMMENT,
			"APPROVAL_STATUS_CHANGED": RepositoryNotificationEvents_APPROVAL_STATUS_CHANGED,
			"APPROVAL_RULE_OVERRIDDEN": RepositoryNotificationEvents_APPROVAL_RULE_OVERRIDDEN,
			"PULL_REQUEST_CREATED": RepositoryNotificationEvents_PULL_REQUEST_CREATED,
			"PULL_REQUEST_SOURCE_UPDATED": RepositoryNotificationEvents_PULL_REQUEST_SOURCE_UPDATED,
			"PULL_REQUEST_STATUS_CHANGED": RepositoryNotificationEvents_PULL_REQUEST_STATUS_CHANGED,
			"PULL_REQUEST_MERGED": RepositoryNotificationEvents_PULL_REQUEST_MERGED,
			"BRANCH_OR_TAG_CREATED": RepositoryNotificationEvents_BRANCH_OR_TAG_CREATED,
			"BRANCH_OR_TAG_DELETED": RepositoryNotificationEvents_BRANCH_OR_TAG_DELETED,
			"BRANCH_OR_TAG_UPDATED": RepositoryNotificationEvents_BRANCH_OR_TAG_UPDATED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.RepositoryNotifyOnOptions",
		reflect.TypeOf((*RepositoryNotifyOnOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.RepositoryProps",
		reflect.TypeOf((*RepositoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_codecommit.RepositoryTriggerOptions",
		reflect.TypeOf((*RepositoryTriggerOptions)(nil)).Elem(),
	)
}
