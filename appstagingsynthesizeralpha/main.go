// Cdk synthesizer for with app-scoped staging stack
package appstagingsynthesizeralpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		reflect.TypeOf((*AppStagingSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBootstrapVersionRule", GoMethod: "AddBootstrapVersionRule"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapQualifier", GoGetter: "BootstrapQualifier"},
			_jsii_.MemberProperty{JsiiProperty: "boundStack", GoGetter: "BoundStack"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromDockerImageAsset", GoMethod: "CloudFormationLocationFromDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "cloudFormationLocationFromFileAsset", GoMethod: "CloudFormationLocationFromFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "emitArtifact", GoMethod: "EmitArtifact"},
			_jsii_.MemberMethod{JsiiMethod: "emitStackArtifact", GoMethod: "EmitStackArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "lookupRole", GoGetter: "LookupRole"},
			_jsii_.MemberMethod{JsiiMethod: "reusableBind", GoMethod: "ReusableBind"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeStackTemplate", GoMethod: "SynthesizeStackTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeTemplate", GoMethod: "SynthesizeTemplate"},
		},
		func() interface{} {
			j := jsiiProxy_AppStagingSynthesizer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStackSynthesizer)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIReusableStackSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizerOptions",
		reflect.TypeOf((*AppStagingSynthesizerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/app-staging-synthesizer-alpha.BootstrapRole",
		reflect.TypeOf((*BootstrapRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "isCliCredentials", GoMethod: "IsCliCredentials"},
		},
		func() interface{} {
			return &jsiiProxy_BootstrapRole{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.BootstrapRoles",
		reflect.TypeOf((*BootstrapRoles)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.CustomFactoryOptions",
		reflect.TypeOf((*CustomFactoryOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.CustomResourcesOptions",
		reflect.TypeOf((*CustomResourcesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.DefaultBootstrapRolesOptions",
		reflect.TypeOf((*DefaultBootstrapRolesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.DefaultResourcesOptions",
		reflect.TypeOf((*DefaultResourcesOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/app-staging-synthesizer-alpha.DefaultStagingStack",
		reflect.TypeOf((*DefaultStagingStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addFile", GoMethod: "AddFile"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addStackTag", GoMethod: "AddStackTag"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyStack", GoGetter: "DependencyStack"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportStringListValue", GoMethod: "ExportStringListValue"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "removeStackTag", GoMethod: "RemoveStackTag"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "stagingBucket", GoGetter: "StagingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "stagingRepos", GoGetter: "StagingRepos"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYamlString", GoMethod: "ToYamlString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_DefaultStagingStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStack)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStagingResources)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.DefaultStagingStackOptions",
		reflect.TypeOf((*DefaultStagingStackOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.DefaultStagingStackProps",
		reflect.TypeOf((*DefaultStagingStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/app-staging-synthesizer-alpha.DeploymentIdentities",
		reflect.TypeOf((*DeploymentIdentities)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFormationExecutionRole", GoGetter: "CloudFormationExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRole", GoGetter: "DeploymentRole"},
			_jsii_.MemberProperty{JsiiProperty: "lookupRole", GoGetter: "LookupRole"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentIdentities{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.FileStagingLocation",
		reflect.TypeOf((*FileStagingLocation)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/app-staging-synthesizer-alpha.IStagingResources",
		reflect.TypeOf((*IStagingResources)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDockerImage", GoMethod: "AddDockerImage"},
			_jsii_.MemberMethod{JsiiMethod: "addFile", GoMethod: "AddFile"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IStagingResources{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/app-staging-synthesizer-alpha.IStagingResourcesFactory",
		reflect.TypeOf((*IStagingResourcesFactory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "obtainStagingResources", GoMethod: "ObtainStagingResources"},
		},
		func() interface{} {
			return &jsiiProxy_IStagingResourcesFactory{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.ImageStagingLocation",
		reflect.TypeOf((*ImageStagingLocation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.ObtainStagingResourcesContext",
		reflect.TypeOf((*ObtainStagingResourcesContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/app-staging-synthesizer-alpha.StagingRoles",
		reflect.TypeOf((*StagingRoles)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/app-staging-synthesizer-alpha.UsingAppStagingSynthesizer",
		reflect.TypeOf((*UsingAppStagingSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_UsingAppStagingSynthesizer{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
}
