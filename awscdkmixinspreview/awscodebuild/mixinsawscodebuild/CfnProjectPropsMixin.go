package mixinsawscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscodebuild/mixinsawscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeBuild::Project` resource configures how AWS CodeBuild builds your source code.
//
// For example, it tells CodeBuild where to get the source code and which build environment to use.
//
// > To unset or remove a project value via CFN, explicitly provide the attribute with value as empty input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProjectPropsMixin := awscdkmixinspreview.Mixins.NewCfnProjectPropsMixin(&CfnProjectMixinProps{
//   	Artifacts: &ArtifactsProperty{
//   		ArtifactIdentifier: jsii.String("artifactIdentifier"),
//   		EncryptionDisabled: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   		Name: jsii.String("name"),
//   		NamespaceType: jsii.String("namespaceType"),
//   		OverrideArtifactName: jsii.Boolean(false),
//   		Packaging: jsii.String("packaging"),
//   		Path: jsii.String("path"),
//   		Type: jsii.String("type"),
//   	},
//   	AutoRetryLimit: jsii.Number(123),
//   	BadgeEnabled: jsii.Boolean(false),
//   	BuildBatchConfig: &ProjectBuildBatchConfigProperty{
//   		BatchReportMode: jsii.String("batchReportMode"),
//   		CombineArtifacts: jsii.Boolean(false),
//   		Restrictions: &BatchRestrictionsProperty{
//   			ComputeTypesAllowed: []*string{
//   				jsii.String("computeTypesAllowed"),
//   			},
//   			MaximumBuildsAllowed: jsii.Number(123),
//   		},
//   		ServiceRole: jsii.String("serviceRole"),
//   		TimeoutInMins: jsii.Number(123),
//   	},
//   	Cache: &ProjectCacheProperty{
//   		CacheNamespace: jsii.String("cacheNamespace"),
//   		Location: jsii.String("location"),
//   		Modes: []*string{
//   			jsii.String("modes"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	ConcurrentBuildLimit: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Environment: &EnvironmentProperty{
//   		Certificate: jsii.String("certificate"),
//   		ComputeType: jsii.String("computeType"),
//   		DockerServer: &DockerServerProperty{
//   			ComputeType: jsii.String("computeType"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   		},
//   		EnvironmentVariables: []interface{}{
//   			&EnvironmentVariableProperty{
//   				Name: jsii.String("name"),
//   				Type: jsii.String("type"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Fleet: &ProjectFleetProperty{
//   			FleetArn: jsii.String("fleetArn"),
//   		},
//   		Image: jsii.String("image"),
//   		ImagePullCredentialsType: jsii.String("imagePullCredentialsType"),
//   		PrivilegedMode: jsii.Boolean(false),
//   		RegistryCredential: &RegistryCredentialProperty{
//   			Credential: jsii.String("credential"),
//   			CredentialProvider: jsii.String("credentialProvider"),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	FileSystemLocations: []interface{}{
//   		&ProjectFileSystemLocationProperty{
//   			Identifier: jsii.String("identifier"),
//   			Location: jsii.String("location"),
//   			MountOptions: jsii.String("mountOptions"),
//   			MountPoint: jsii.String("mountPoint"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	LogsConfig: &LogsConfigProperty{
//   		CloudWatchLogs: &CloudWatchLogsConfigProperty{
//   			GroupName: jsii.String("groupName"),
//   			Status: jsii.String("status"),
//   			StreamName: jsii.String("streamName"),
//   		},
//   		S3Logs: &S3LogsConfigProperty{
//   			EncryptionDisabled: jsii.Boolean(false),
//   			Location: jsii.String("location"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	QueuedTimeoutInMinutes: jsii.Number(123),
//   	ResourceAccessRole: jsii.String("resourceAccessRole"),
//   	SecondaryArtifacts: []interface{}{
//   		&ArtifactsProperty{
//   			ArtifactIdentifier: jsii.String("artifactIdentifier"),
//   			EncryptionDisabled: jsii.Boolean(false),
//   			Location: jsii.String("location"),
//   			Name: jsii.String("name"),
//   			NamespaceType: jsii.String("namespaceType"),
//   			OverrideArtifactName: jsii.Boolean(false),
//   			Packaging: jsii.String("packaging"),
//   			Path: jsii.String("path"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SecondarySources: []interface{}{
//   		&SourceProperty{
//   			Auth: &SourceAuthProperty{
//   				Resource: jsii.String("resource"),
//   				Type: jsii.String("type"),
//   			},
//   			BuildSpec: jsii.String("buildSpec"),
//   			BuildStatusConfig: &BuildStatusConfigProperty{
//   				Context: jsii.String("context"),
//   				TargetUrl: jsii.String("targetUrl"),
//   			},
//   			GitCloneDepth: jsii.Number(123),
//   			GitSubmodulesConfig: &GitSubmodulesConfigProperty{
//   				FetchSubmodules: jsii.Boolean(false),
//   			},
//   			InsecureSsl: jsii.Boolean(false),
//   			Location: jsii.String("location"),
//   			ReportBuildStatus: jsii.Boolean(false),
//   			SourceIdentifier: jsii.String("sourceIdentifier"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SecondarySourceVersions: []interface{}{
//   		&ProjectSourceVersionProperty{
//   			SourceIdentifier: jsii.String("sourceIdentifier"),
//   			SourceVersion: jsii.String("sourceVersion"),
//   		},
//   	},
//   	ServiceRole: jsii.String("serviceRole"),
//   	Source: &SourceProperty{
//   		Auth: &SourceAuthProperty{
//   			Resource: jsii.String("resource"),
//   			Type: jsii.String("type"),
//   		},
//   		BuildSpec: jsii.String("buildSpec"),
//   		BuildStatusConfig: &BuildStatusConfigProperty{
//   			Context: jsii.String("context"),
//   			TargetUrl: jsii.String("targetUrl"),
//   		},
//   		GitCloneDepth: jsii.Number(123),
//   		GitSubmodulesConfig: &GitSubmodulesConfigProperty{
//   			FetchSubmodules: jsii.Boolean(false),
//   		},
//   		InsecureSsl: jsii.Boolean(false),
//   		Location: jsii.String("location"),
//   		ReportBuildStatus: jsii.Boolean(false),
//   		SourceIdentifier: jsii.String("sourceIdentifier"),
//   		Type: jsii.String("type"),
//   	},
//   	SourceVersion: jsii.String("sourceVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutInMinutes: jsii.Number(123),
//   	Triggers: &ProjectTriggersProperty{
//   		BuildType: jsii.String("buildType"),
//   		FilterGroups: []interface{}{
//   			[]interface{}{
//   				&WebhookFilterProperty{
//   					ExcludeMatchedPattern: jsii.Boolean(false),
//   					Pattern: jsii.String("pattern"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		PullRequestBuildPolicy: &PullRequestBuildPolicyProperty{
//   			ApproverRoles: []*string{
//   				jsii.String("approverRoles"),
//   			},
//   			RequiresCommentApproval: jsii.String("requiresCommentApproval"),
//   		},
//   		ScopeConfiguration: &ScopeConfigurationProperty{
//   			Domain: jsii.String("domain"),
//   			Name: jsii.String("name"),
//   			Scope: jsii.String("scope"),
//   		},
//   		Webhook: jsii.Boolean(false),
//   	},
//   	Visibility: jsii.String("visibility"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html
//
type CfnProjectPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnProjectMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProjectPropsMixin
type jsiiProxy_CfnProjectPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnProjectPropsMixin) Props() *CfnProjectMixinProps {
	var returns *CfnProjectMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProjectPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeBuild::Project`.
func NewCfnProjectPropsMixin(props *CfnProjectMixinProps, options *mixins.CfnPropertyMixinOptions) CfnProjectPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProjectPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProjectPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeBuild::Project`.
func NewCfnProjectPropsMixin_Override(c CfnProjectPropsMixin, props *CfnProjectMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnProjectPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProjectPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProjectPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codebuild.mixins.CfnProjectPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProjectPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProjectPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

