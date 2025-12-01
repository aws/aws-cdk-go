// The CDK Construct Library for EC2 Image Builder
package awsimagebuilderalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImage",
		reflect.TypeOf((*AmazonManagedImage)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AmazonManagedImage{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImageAttributes",
		reflect.TypeOf((*AmazonManagedImageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedImageOptions",
		reflect.TypeOf((*AmazonManagedImageOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AmiDistribution",
		reflect.TypeOf((*AmiDistribution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AmiLaunchPermission",
		reflect.TypeOf((*AmiLaunchPermission)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponent",
		reflect.TypeOf((*AwsManagedComponent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AwsManagedComponent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedComponentAttributes",
		reflect.TypeOf((*AwsManagedComponentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		reflect.TypeOf((*AwsManagedWorkflow)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AwsManagedWorkflow{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflowAttributes",
		reflect.TypeOf((*AwsManagedWorkflowAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.AwsMarketplaceComponent",
		reflect.TypeOf((*AwsMarketplaceComponent)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AwsMarketplaceComponent{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.AwsMarketplaceComponentAttributes",
		reflect.TypeOf((*AwsMarketplaceComponentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.BaseContainerImage",
		reflect.TypeOf((*BaseContainerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
		},
		func() interface{} {
			return &jsiiProxy_BaseContainerImage{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.BaseImage",
		reflect.TypeOf((*BaseImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
		},
		func() interface{} {
			return &jsiiProxy_BaseImage{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.Component",
		reflect.TypeOf((*Component)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "componentArn", GoGetter: "ComponentArn"},
			_jsii_.MemberProperty{JsiiProperty: "componentName", GoGetter: "ComponentName"},
			_jsii_.MemberProperty{JsiiProperty: "componentType", GoGetter: "ComponentType"},
			_jsii_.MemberProperty{JsiiProperty: "componentVersion", GoGetter: "ComponentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "encrypted", GoGetter: "Encrypted"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Component{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComponent)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentAction",
		reflect.TypeOf((*ComponentAction)(nil)).Elem(),
		map[string]interface{}{
			"APPEND_FILE": ComponentAction_APPEND_FILE,
			"ASSERT": ComponentAction_ASSERT,
			"COPY_FILE": ComponentAction_COPY_FILE,
			"COPY_FOLDER": ComponentAction_COPY_FOLDER,
			"CREATE_FILE": ComponentAction_CREATE_FILE,
			"CREATE_FOLDER": ComponentAction_CREATE_FOLDER,
			"CREATE_SYMLINK": ComponentAction_CREATE_SYMLINK,
			"DELETE_FILE": ComponentAction_DELETE_FILE,
			"DELETE_FOLDER": ComponentAction_DELETE_FOLDER,
			"EXECUTE_BASH": ComponentAction_EXECUTE_BASH,
			"EXECUTE_BINARY": ComponentAction_EXECUTE_BINARY,
			"EXECUTE_DOCUMENT": ComponentAction_EXECUTE_DOCUMENT,
			"EXECUTE_POWERSHELL": ComponentAction_EXECUTE_POWERSHELL,
			"INSTALL_MSI": ComponentAction_INSTALL_MSI,
			"LIST_FILES": ComponentAction_LIST_FILES,
			"MOVE_FILE": ComponentAction_MOVE_FILE,
			"MOVE_FOLDER": ComponentAction_MOVE_FOLDER,
			"READ_FILE": ComponentAction_READ_FILE,
			"REBOOT": ComponentAction_REBOOT,
			"SET_FILE_ENCODING": ComponentAction_SET_FILE_ENCODING,
			"SET_FILE_OWNER": ComponentAction_SET_FILE_OWNER,
			"SET_FOLDER_OWNER": ComponentAction_SET_FOLDER_OWNER,
			"SET_FILE_PERMISSIONS": ComponentAction_SET_FILE_PERMISSIONS,
			"SET_FOLDER_PERMISSIONS": ComponentAction_SET_FOLDER_PERMISSIONS,
			"SET_REGISTRY": ComponentAction_SET_REGISTRY,
			"S3_DOWNLOAD": ComponentAction_S3_DOWNLOAD,
			"S3_UPLOAD": ComponentAction_S3_UPLOAD,
			"UNINSTALL_MSI": ComponentAction_UNINSTALL_MSI,
			"UPDATE_OS": ComponentAction_UPDATE_OS,
			"WEB_DOWNLOAD": ComponentAction_WEB_DOWNLOAD,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentAttributes",
		reflect.TypeOf((*ComponentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentConfiguration",
		reflect.TypeOf((*ComponentConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentConstantValue",
		reflect.TypeOf((*ComponentConstantValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ComponentConstantValue{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentData",
		reflect.TypeOf((*ComponentData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isS3Reference", GoGetter: "IsS3Reference"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ComponentData{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentDocument",
		reflect.TypeOf((*ComponentDocument)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentDocumentForLoop",
		reflect.TypeOf((*ComponentDocumentForLoop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentDocumentLoop",
		reflect.TypeOf((*ComponentDocumentLoop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentDocumentParameterDefinition",
		reflect.TypeOf((*ComponentDocumentParameterDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentDocumentPhase",
		reflect.TypeOf((*ComponentDocumentPhase)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentDocumentStep",
		reflect.TypeOf((*ComponentDocumentStep)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentOnFailure",
		reflect.TypeOf((*ComponentOnFailure)(nil)).Elem(),
		map[string]interface{}{
			"ABORT": ComponentOnFailure_ABORT,
			"CONTINUE": ComponentOnFailure_CONTINUE,
			"IGNORE": ComponentOnFailure_IGNORE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentParameterType",
		reflect.TypeOf((*ComponentParameterType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": ComponentParameterType_STRING,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentParameterValue",
		reflect.TypeOf((*ComponentParameterValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ComponentParameterValue{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentPhaseName",
		reflect.TypeOf((*ComponentPhaseName)(nil)).Elem(),
		map[string]interface{}{
			"BUILD": ComponentPhaseName_BUILD,
			"CONTAINER_HOST_TEST": ComponentPhaseName_CONTAINER_HOST_TEST,
			"TEST": ComponentPhaseName_TEST,
			"VALIDATE": ComponentPhaseName_VALIDATE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentProps",
		reflect.TypeOf((*ComponentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentSchemaVersion",
		reflect.TypeOf((*ComponentSchemaVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1_0": ComponentSchemaVersion_V1_0,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepIfCondition",
		reflect.TypeOf((*ComponentStepIfCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ifCondition", GoGetter: "IfCondition"},
		},
		func() interface{} {
			return &jsiiProxy_ComponentStepIfCondition{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ComponentStepInputs",
		reflect.TypeOf((*ComponentStepInputs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inputs", GoGetter: "Inputs"},
		},
		func() interface{} {
			return &jsiiProxy_ComponentStepInputs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerDistribution",
		reflect.TypeOf((*ContainerDistribution)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerInstanceImage",
		reflect.TypeOf((*ContainerInstanceImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
		},
		func() interface{} {
			return &jsiiProxy_ContainerInstanceImage{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerRecipe",
		reflect.TypeOf((*ContainerRecipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInstanceBlockDevice", GoMethod: "AddInstanceBlockDevice"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeArn", GoGetter: "ContainerRecipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeName", GoGetter: "ContainerRecipeName"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeVersion", GoGetter: "ContainerRecipeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ContainerRecipe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ContainerRecipeBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerRecipeAttributes",
		reflect.TypeOf((*ContainerRecipeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerRecipeBase",
		reflect.TypeOf((*ContainerRecipeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeArn", GoGetter: "ContainerRecipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeName", GoGetter: "ContainerRecipeName"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeVersion", GoGetter: "ContainerRecipeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ContainerRecipeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IContainerRecipe)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerRecipeProps",
		reflect.TypeOf((*ContainerRecipeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ContainerType",
		reflect.TypeOf((*ContainerType)(nil)).Elem(),
		map[string]interface{}{
			"DOCKER": ContainerType_DOCKER,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfiguration",
		reflect.TypeOf((*DistributionConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAmiDistributions", GoMethod: "AddAmiDistributions"},
			_jsii_.MemberMethod{JsiiMethod: "addContainerDistributions", GoMethod: "AddContainerDistributions"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfigurationArn", GoGetter: "DistributionConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfigurationName", GoGetter: "DistributionConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DistributionConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDistributionConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.DistributionConfigurationProps",
		reflect.TypeOf((*DistributionConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.DockerfileData",
		reflect.TypeOf((*DockerfileData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_DockerfileData{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.DockerfileTemplateConfig",
		reflect.TypeOf((*DockerfileTemplateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.FastLaunchConfiguration",
		reflect.TypeOf((*FastLaunchConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.HttpTokens",
		reflect.TypeOf((*HttpTokens)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": HttpTokens_OPTIONAL,
			"REQUIRED": HttpTokens_REQUIRED,
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IComponent",
		reflect.TypeOf((*IComponent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "componentArn", GoGetter: "ComponentArn"},
			_jsii_.MemberProperty{JsiiProperty: "componentName", GoGetter: "ComponentName"},
			_jsii_.MemberProperty{JsiiProperty: "componentVersion", GoGetter: "ComponentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IComponent{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IContainerRecipe",
		reflect.TypeOf((*IContainerRecipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeArn", GoGetter: "ContainerRecipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeName", GoGetter: "ContainerRecipeName"},
			_jsii_.MemberProperty{JsiiProperty: "containerRecipeVersion", GoGetter: "ContainerRecipeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IContainerRecipe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRecipeBase)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IDistributionConfiguration",
		reflect.TypeOf((*IDistributionConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfigurationArn", GoGetter: "DistributionConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionConfigurationName", GoGetter: "DistributionConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDistributionConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IImage",
		reflect.TypeOf((*IImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDefaultExecutionRolePermissions", GoMethod: "GrantDefaultExecutionRolePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "imageArn", GoGetter: "ImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersion", GoGetter: "ImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toBaseImage", GoMethod: "ToBaseImage"},
			_jsii_.MemberMethod{JsiiMethod: "toContainerBaseImage", GoMethod: "ToContainerBaseImage"},
		},
		func() interface{} {
			j := jsiiProxy_IImage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IImagePipeline",
		reflect.TypeOf((*IImagePipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDefaultExecutionRolePermissions", GoMethod: "GrantDefaultExecutionRolePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantStartExecution", GoMethod: "GrantStartExecution"},
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineArn", GoGetter: "ImagePipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineName", GoGetter: "ImagePipelineName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCVEDetected", GoMethod: "OnCVEDetected"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildCompleted", GoMethod: "OnImageBuildCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildFailed", GoMethod: "OnImageBuildFailed"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildStateChange", GoMethod: "OnImageBuildStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildSucceeded", GoMethod: "OnImageBuildSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "onImagePipelineAutoDisabled", GoMethod: "OnImagePipelineAutoDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "onWaitForAction", GoMethod: "OnWaitForAction"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IImagePipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IImageRecipe",
		reflect.TypeOf((*IImageRecipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "imageRecipeArn", GoGetter: "ImageRecipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "imageRecipeName", GoGetter: "ImageRecipeName"},
			_jsii_.MemberProperty{JsiiProperty: "imageRecipeVersion", GoGetter: "ImageRecipeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IImageRecipe{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRecipeBase)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IInfrastructureConfiguration",
		reflect.TypeOf((*IInfrastructureConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationArn", GoGetter: "InfrastructureConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationName", GoGetter: "InfrastructureConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IInfrastructureConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.ILifecyclePolicy",
		reflect.TypeOf((*ILifecyclePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "lifecyclePolicyArn", GoGetter: "LifecyclePolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecyclePolicyName", GoGetter: "LifecyclePolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_ILifecyclePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IRecipeBase",
		reflect.TypeOf((*IRecipeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IRecipeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-imagebuilder-alpha.IWorkflow",
		reflect.TypeOf((*IWorkflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "workflowArn", GoGetter: "WorkflowArn"},
			_jsii_.MemberProperty{JsiiProperty: "workflowName", GoGetter: "WorkflowName"},
			_jsii_.MemberProperty{JsiiProperty: "workflowType", GoGetter: "WorkflowType"},
			_jsii_.MemberProperty{JsiiProperty: "workflowVersion", GoGetter: "WorkflowVersion"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkflow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.Image",
		reflect.TypeOf((*Image)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDefaultExecutionRolePermissions", GoMethod: "GrantDefaultExecutionRolePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "imageArn", GoGetter: "ImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersion", GoGetter: "ImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfiguration", GoGetter: "InfrastructureConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toBaseImage", GoMethod: "ToBaseImage"},
			_jsii_.MemberMethod{JsiiMethod: "toContainerBaseImage", GoMethod: "ToContainerBaseImage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Image{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImage)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ImageArchitecture",
		reflect.TypeOf((*ImageArchitecture)(nil)).Elem(),
		map[string]interface{}{
			"ARM64": ImageArchitecture_ARM64,
			"X86_64": ImageArchitecture_X86_64,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ImageAttributes",
		reflect.TypeOf((*ImageAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ImagePipeline",
		reflect.TypeOf((*ImagePipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantDefaultExecutionRolePermissions", GoMethod: "GrantDefaultExecutionRolePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantStartExecution", GoMethod: "GrantStartExecution"},
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineArn", GoGetter: "ImagePipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "imagePipelineName", GoGetter: "ImagePipelineName"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfiguration", GoGetter: "InfrastructureConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCVEDetected", GoMethod: "OnCVEDetected"},
			_jsii_.MemberMethod{JsiiMethod: "onEvent", GoMethod: "OnEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildCompleted", GoMethod: "OnImageBuildCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildFailed", GoMethod: "OnImageBuildFailed"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildStateChange", GoMethod: "OnImageBuildStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "onImageBuildSucceeded", GoMethod: "OnImageBuildSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "onImagePipelineAutoDisabled", GoMethod: "OnImagePipelineAutoDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "onWaitForAction", GoMethod: "OnWaitForAction"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ImagePipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImagePipeline)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ImagePipelineProps",
		reflect.TypeOf((*ImagePipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ImagePipelineSchedule",
		reflect.TypeOf((*ImagePipelineSchedule)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ImagePipelineStatus",
		reflect.TypeOf((*ImagePipelineStatus)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": ImagePipelineStatus_ENABLED,
			"DISABLED": ImagePipelineStatus_DISABLED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ImageProps",
		reflect.TypeOf((*ImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.ImageRecipe",
		reflect.TypeOf((*ImageRecipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBlockDevice", GoMethod: "AddBlockDevice"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "imageRecipeArn", GoGetter: "ImageRecipeArn"},
			_jsii_.MemberProperty{JsiiProperty: "imageRecipeName", GoGetter: "ImageRecipeName"},
			_jsii_.MemberProperty{JsiiProperty: "imageRecipeVersion", GoGetter: "ImageRecipeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ImageRecipe{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImageRecipe)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ImageRecipeAttributes",
		reflect.TypeOf((*ImageRecipeAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.ImageRecipeProps",
		reflect.TypeOf((*ImageRecipeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ImageType",
		reflect.TypeOf((*ImageType)(nil)).Elem(),
		map[string]interface{}{
			"AMI": ImageType_AMI,
			"DOCKER": ImageType_DOCKER,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.InfrastructureConfiguration",
		reflect.TypeOf((*InfrastructureConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationArn", GoGetter: "InfrastructureConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureConfigurationName", GoGetter: "InfrastructureConfigurationName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfile", GoGetter: "InstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "logBucket", GoGetter: "LogBucket"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InfrastructureConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInfrastructureConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.InfrastructureConfigurationLogging",
		reflect.TypeOf((*InfrastructureConfigurationLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.InfrastructureConfigurationProps",
		reflect.TypeOf((*InfrastructureConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LaunchTemplateConfiguration",
		reflect.TypeOf((*LaunchTemplateConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicy",
		reflect.TypeOf((*LifecyclePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "lifecyclePolicyArn", GoGetter: "LifecyclePolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecyclePolicyName", GoGetter: "LifecyclePolicyName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LifecyclePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILifecyclePolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyAction",
		reflect.TypeOf((*LifecyclePolicyAction)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyActionType",
		reflect.TypeOf((*LifecyclePolicyActionType)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": LifecyclePolicyActionType_DELETE,
			"DEPRECATE": LifecyclePolicyActionType_DEPRECATE,
			"DISABLE": LifecyclePolicyActionType_DISABLE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyAgeFilter",
		reflect.TypeOf((*LifecyclePolicyAgeFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyAmiExclusionRules",
		reflect.TypeOf((*LifecyclePolicyAmiExclusionRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyCountFilter",
		reflect.TypeOf((*LifecyclePolicyCountFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyDetail",
		reflect.TypeOf((*LifecyclePolicyDetail)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyExclusionRules",
		reflect.TypeOf((*LifecyclePolicyExclusionRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyFilter",
		reflect.TypeOf((*LifecyclePolicyFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyImageExclusionRules",
		reflect.TypeOf((*LifecyclePolicyImageExclusionRules)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyProps",
		reflect.TypeOf((*LifecyclePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyResourceSelection",
		reflect.TypeOf((*LifecyclePolicyResourceSelection)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyResourceType",
		reflect.TypeOf((*LifecyclePolicyResourceType)(nil)).Elem(),
		map[string]interface{}{
			"AMI_IMAGE": LifecyclePolicyResourceType_AMI_IMAGE,
			"CONTAINER_IMAGE": LifecyclePolicyResourceType_CONTAINER_IMAGE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.LifecyclePolicyStatus",
		reflect.TypeOf((*LifecyclePolicyStatus)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": LifecyclePolicyStatus_ENABLED,
			"DISABLED": LifecyclePolicyStatus_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.OSVersion",
		reflect.TypeOf((*OSVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "osVersion", GoGetter: "OsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
		},
		func() interface{} {
			return &jsiiProxy_OSVersion{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.Platform",
		reflect.TypeOf((*Platform)(nil)).Elem(),
		map[string]interface{}{
			"LINUX": Platform_LINUX,
			"WINDOWS": Platform_WINDOWS,
			"MAC_OS": Platform_MAC_OS,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.Repository",
		reflect.TypeOf((*Repository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "repositoryName", GoGetter: "RepositoryName"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
		},
		func() interface{} {
			return &jsiiProxy_Repository{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.RepositoryService",
		reflect.TypeOf((*RepositoryService)(nil)).Elem(),
		map[string]interface{}{
			"ECR": RepositoryService_ECR,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.S3ComponentData",
		reflect.TypeOf((*S3ComponentData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "isS3Reference", GoGetter: "IsS3Reference"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_S3ComponentData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ComponentData)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.S3DockerfileData",
		reflect.TypeOf((*S3DockerfileData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			j := jsiiProxy_S3DockerfileData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DockerfileData)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.S3WorkflowData",
		reflect.TypeOf((*S3WorkflowData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			j := jsiiProxy_S3WorkflowData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_WorkflowData)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.SSMParameterConfigurations",
		reflect.TypeOf((*SSMParameterConfigurations)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.ScheduleStartCondition",
		reflect.TypeOf((*ScheduleStartCondition)(nil)).Elem(),
		map[string]interface{}{
			"EXPRESSION_MATCH_ONLY": ScheduleStartCondition_EXPRESSION_MATCH_ONLY,
			"EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE": ScheduleStartCondition_EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.Tenancy",
		reflect.TypeOf((*Tenancy)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": Tenancy_DEFAULT,
			"DEDICATED": Tenancy_DEDICATED,
			"HOST": Tenancy_HOST,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.Workflow",
		reflect.TypeOf((*Workflow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workflowArn", GoGetter: "WorkflowArn"},
			_jsii_.MemberProperty{JsiiProperty: "workflowName", GoGetter: "WorkflowName"},
			_jsii_.MemberProperty{JsiiProperty: "workflowType", GoGetter: "WorkflowType"},
			_jsii_.MemberProperty{JsiiProperty: "workflowVersion", GoGetter: "WorkflowVersion"},
		},
		func() interface{} {
			j := jsiiProxy_Workflow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkflow)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowAction",
		reflect.TypeOf((*WorkflowAction)(nil)).Elem(),
		map[string]interface{}{
			"APPLY_IMAGE_CONFIGURATIONS": WorkflowAction_APPLY_IMAGE_CONFIGURATIONS,
			"BOOTSTRAP_INSTANCE_FOR_CONTAINER": WorkflowAction_BOOTSTRAP_INSTANCE_FOR_CONTAINER,
			"COLLECT_IMAGE_METADATA": WorkflowAction_COLLECT_IMAGE_METADATA,
			"COLLECT_IMAGE_SCAN_FINDINGS": WorkflowAction_COLLECT_IMAGE_SCAN_FINDINGS,
			"CREATE_IMAGE": WorkflowAction_CREATE_IMAGE,
			"DISTRIBUTE_IMAGE": WorkflowAction_DISTRIBUTE_IMAGE,
			"EXECUTE_COMPONENTS": WorkflowAction_EXECUTE_COMPONENTS,
			"EXECUTE_STATE_MACHINE": WorkflowAction_EXECUTE_STATE_MACHINE,
			"LAUNCH_INSTANCE": WorkflowAction_LAUNCH_INSTANCE,
			"MODIFY_IMAGE_ATTRIBUTES": WorkflowAction_MODIFY_IMAGE_ATTRIBUTES,
			"RUN_COMMAND": WorkflowAction_RUN_COMMAND,
			"REGISTER_IMAGE": WorkflowAction_REGISTER_IMAGE,
			"RUN_SYS_PREP": WorkflowAction_RUN_SYS_PREP,
			"SANITIZE_INSTANCE": WorkflowAction_SANITIZE_INSTANCE,
			"TERMINATE_INSTANCE": WorkflowAction_TERMINATE_INSTANCE,
			"WAIT_FOR_ACTION": WorkflowAction_WAIT_FOR_ACTION,
			"WAIT_FOR_SSM_AGENT": WorkflowAction_WAIT_FOR_SSM_AGENT,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowAttributes",
		reflect.TypeOf((*WorkflowAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowConfiguration",
		reflect.TypeOf((*WorkflowConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowData",
		reflect.TypeOf((*WorkflowData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "render", GoMethod: "Render"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowData{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowDataConfig",
		reflect.TypeOf((*WorkflowDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowOnFailure",
		reflect.TypeOf((*WorkflowOnFailure)(nil)).Elem(),
		map[string]interface{}{
			"ABORT": WorkflowOnFailure_ABORT,
			"CONTINUE": WorkflowOnFailure_CONTINUE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterType",
		reflect.TypeOf((*WorkflowParameterType)(nil)).Elem(),
		map[string]interface{}{
			"BOOLEAN": WorkflowParameterType_BOOLEAN,
			"INTEGER": WorkflowParameterType_INTEGER,
			"STRING": WorkflowParameterType_STRING,
			"STRING_LIST": WorkflowParameterType_STRING_LIST,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		reflect.TypeOf((*WorkflowParameterValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_WorkflowParameterValue{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowProps",
		reflect.TypeOf((*WorkflowProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowSchemaVersion",
		reflect.TypeOf((*WorkflowSchemaVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1_0": WorkflowSchemaVersion_V1_0,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowType",
		reflect.TypeOf((*WorkflowType)(nil)).Elem(),
		map[string]interface{}{
			"BUILD": WorkflowType_BUILD,
			"TEST": WorkflowType_TEST,
			"DISTRIBUTION": WorkflowType_DISTRIBUTION,
		},
	)
}
