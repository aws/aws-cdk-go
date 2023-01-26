package awskendra

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		reflect.TypeOf((*CfnDataSource)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customDocumentEnrichmentConfiguration", GoGetter: "CustomDocumentEnrichmentConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceConfiguration", GoGetter: "DataSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "indexId", GoGetter: "IndexId"},
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
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.AccessControlListConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_AccessControlListConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.AclConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_AclConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ColumnConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ColumnConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceAttachmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceAttachmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceAttachmentToIndexFieldMappingProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceAttachmentToIndexFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceBlogConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceBlogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceBlogToIndexFieldMappingProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceBlogToIndexFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluencePageConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluencePageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluencePageToIndexFieldMappingProperty",
		reflect.TypeOf((*CfnDataSource_ConfluencePageToIndexFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceSpaceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceSpaceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConfluenceSpaceToIndexFieldMappingProperty",
		reflect.TypeOf((*CfnDataSource_ConfluenceSpaceToIndexFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ConnectionConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ConnectionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.CustomDocumentEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_CustomDocumentEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DataSourceToIndexFieldMappingProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceToIndexFieldMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DataSourceVpcConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_DataSourceVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DatabaseConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_DatabaseConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DocumentAttributeConditionProperty",
		reflect.TypeOf((*CfnDataSource_DocumentAttributeConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DocumentAttributeTargetProperty",
		reflect.TypeOf((*CfnDataSource_DocumentAttributeTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DocumentAttributeValueProperty",
		reflect.TypeOf((*CfnDataSource_DocumentAttributeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.DocumentsMetadataConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_DocumentsMetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.GoogleDriveConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_GoogleDriveConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.HookConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_HookConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.InlineCustomDocumentEnrichmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_InlineCustomDocumentEnrichmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.OneDriveConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_OneDriveConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.OneDriveUsersProperty",
		reflect.TypeOf((*CfnDataSource_OneDriveUsersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ProxyConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ProxyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.S3DataSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_S3DataSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.S3PathProperty",
		reflect.TypeOf((*CfnDataSource_S3PathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceChatterFeedConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceChatterFeedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceCustomKnowledgeArticleTypeConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceCustomKnowledgeArticleTypeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceKnowledgeArticleConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceKnowledgeArticleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceStandardKnowledgeArticleTypeConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceStandardKnowledgeArticleTypeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceStandardObjectAttachmentConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceStandardObjectAttachmentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SalesforceStandardObjectConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SalesforceStandardObjectConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ServiceNowConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ServiceNowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ServiceNowKnowledgeArticleConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ServiceNowKnowledgeArticleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.ServiceNowServiceCatalogConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_ServiceNowServiceCatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SharePointConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SharePointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.SqlConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_SqlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WebCrawlerAuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerAuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WebCrawlerBasicAuthenticationProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerBasicAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WebCrawlerConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WebCrawlerSeedUrlConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerSeedUrlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WebCrawlerSiteMapsConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerSiteMapsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WebCrawlerUrlsProperty",
		reflect.TypeOf((*CfnDataSource_WebCrawlerUrlsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSource.WorkDocsConfigurationProperty",
		reflect.TypeOf((*CfnDataSource_WorkDocsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnDataSourceProps",
		reflect.TypeOf((*CfnDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		reflect.TypeOf((*CfnFaq)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fileFormat", GoGetter: "FileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "indexId", GoGetter: "IndexId"},
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
			_jsii_.MemberProperty{JsiiProperty: "s3Path", GoGetter: "S3Path"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFaq{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnFaq.S3PathProperty",
		reflect.TypeOf((*CfnFaq_S3PathProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnFaqProps",
		reflect.TypeOf((*CfnFaqProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		reflect.TypeOf((*CfnIndex)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "capacityUnits", GoGetter: "CapacityUnits"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "documentMetadataConfigurations", GoGetter: "DocumentMetadataConfigurations"},
			_jsii_.MemberProperty{JsiiProperty: "edition", GoGetter: "Edition"},
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
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryptionConfiguration", GoGetter: "ServerSideEncryptionConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberProperty{JsiiProperty: "userContextPolicy", GoGetter: "UserContextPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "userTokenConfigurations", GoGetter: "UserTokenConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.CapacityUnitsConfigurationProperty",
		reflect.TypeOf((*CfnIndex_CapacityUnitsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.DocumentMetadataConfigurationProperty",
		reflect.TypeOf((*CfnIndex_DocumentMetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.JsonTokenTypeConfigurationProperty",
		reflect.TypeOf((*CfnIndex_JsonTokenTypeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.JwtTokenTypeConfigurationProperty",
		reflect.TypeOf((*CfnIndex_JwtTokenTypeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.RelevanceProperty",
		reflect.TypeOf((*CfnIndex_RelevanceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.SearchProperty",
		reflect.TypeOf((*CfnIndex_SearchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.ServerSideEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnIndex_ServerSideEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.UserTokenConfigurationProperty",
		reflect.TypeOf((*CfnIndex_UserTokenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndex.ValueImportanceItemProperty",
		reflect.TypeOf((*CfnIndex_ValueImportanceItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kendra.CfnIndexProps",
		reflect.TypeOf((*CfnIndexProps)(nil)).Elem(),
	)
}
