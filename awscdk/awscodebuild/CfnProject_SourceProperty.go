package awscodebuild


// `Source` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies the source code settings for the project, such as the source code's repository type and location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Auth: &SourceAuthProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Resource: jsii.String("resource"),
//   	},
//   	BuildSpec: jsii.String("buildSpec"),
//   	BuildStatusConfig: &BuildStatusConfigProperty{
//   		Context: jsii.String("context"),
//   		TargetUrl: jsii.String("targetUrl"),
//   	},
//   	GitCloneDepth: jsii.Number(123),
//   	GitSubmodulesConfig: &GitSubmodulesConfigProperty{
//   		FetchSubmodules: jsii.Boolean(false),
//   	},
//   	InsecureSsl: jsii.Boolean(false),
//   	Location: jsii.String("location"),
//   	ReportBuildStatus: jsii.Boolean(false),
//   	SourceIdentifier: jsii.String("sourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html
//
type CfnProject_SourceProperty struct {
	// The type of repository that contains the source code to be built. Valid values include:.
	//
	// - `BITBUCKET` : The source code is in a Bitbucket repository.
	// - `CODECOMMIT` : The source code is in an CodeCommit repository.
	// - `CODEPIPELINE` : The source code settings are specified in the source action of a pipeline in CodePipeline.
	// - `GITHUB` : The source code is in a GitHub repository.
	// - `GITHUB_ENTERPRISE` : The source code is in a GitHub Enterprise Server repository.
	// - `GITLAB` : The source code is in a GitLab repository.
	// - `GITLAB_SELF_MANAGED` : The source code is in a self-managed GitLab repository.
	// - `NO_SOURCE` : The project does not have input source code.
	// - `S3` : The source code is in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Information about the authorization settings for AWS CodeBuild to access the source code to be built.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// The build specification for the project.
	//
	// If this value is not provided, then the source code must contain a buildspec file named `buildspec.yml` at the root level. If this value is provided, it can be either a single string containing the entire build specification, or the path to an alternate buildspec file relative to the value of the built-in environment variable `CODEBUILD_SRC_DIR` . The alternate buildspec file can have a name other than `buildspec.yml` , for example `myspec.yml` or `build_spec_qa.yml` or similar. For more information, see the [Build Spec Reference](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example) in the *AWS CodeBuild User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-buildspec
	//
	BuildSpec *string `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Contains information that defines how the build project reports the build status to the source provider.
	//
	// This option is only used when the source provider is `GITHUB` , `GITHUB_ENTERPRISE` , or `BITBUCKET` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-buildstatusconfig
	//
	BuildStatusConfig interface{} `field:"optional" json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// The depth of history to download.
	//
	// Minimum value is 0. If this value is 0, greater than 25, or not provided, then the full history is downloaded with each build project. If your source type is Amazon S3, this value is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-gitclonedepth
	//
	GitCloneDepth *float64 `field:"optional" json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// Information about the Git submodules configuration for the build project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-gitsubmodulesconfig
	//
	GitSubmodulesConfig interface{} `field:"optional" json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// This is used with GitHub Enterprise only.
	//
	// Set to true to ignore SSL warnings while connecting to your GitHub Enterprise project repository. The default value is `false` . `InsecureSsl` should be used for testing purposes only. It should not be used in a production environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-insecuressl
	//
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// Information about the location of the source code to be built. Valid values include:.
	//
	// - For source code settings that are specified in the source action of a pipeline in CodePipeline, `location` should not be specified. If it is specified, CodePipeline ignores it. This is because CodePipeline uses the settings in a pipeline's source action instead of this value.
	// - For source code in an CodeCommit repository, the HTTPS clone URL to the repository that contains the source code and the buildspec file (for example, `https://git-codecommit.<region-ID>.amazonaws.com/v1/repos/<repo-name>` ).
	// - For source code in an Amazon S3 input bucket, one of the following.
	//
	// - The path to the ZIP file that contains the source code (for example, `<bucket-name>/<path>/<object-name>.zip` ).
	// - The path to the folder that contains the source code (for example, `<bucket-name>/<path-to-source-code>/<folder>/` ).
	// - For source code in a GitHub repository, the HTTPS clone URL to the repository that contains the source and the buildspec file. You must connect your AWS account to your GitHub account. Use the AWS CodeBuild console to start creating a build project. When you use the console to connect (or reconnect) with GitHub, on the GitHub *Authorize application* page, for *Organization access* , choose *Request access* next to each repository you want to allow AWS CodeBuild to have access to, and then choose *Authorize application* . (After you have connected to your GitHub account, you do not need to finish creating the build project. You can leave the AWS CodeBuild console.) To instruct AWS CodeBuild to use this connection, in the `source` object, set the `auth` object's `type` value to `OAUTH` .
	// - For source code in an GitLab or self-managed GitLab repository, the HTTPS clone URL to the repository that contains the source and the buildspec file. You must connect your AWS account to your GitLab account. Use the AWS CodeBuild console to start creating a build project. When you use the console to connect (or reconnect) with GitLab, on the Connections *Authorize application* page, choose *Authorize* . Then on the AWS CodeConnections *Create GitLab connection* page, choose *Connect to GitLab* . (After you have connected to your GitLab account, you do not need to finish creating the build project. You can leave the AWS CodeBuild console.) To instruct AWS CodeBuild to override the default connection and use this connection instead, set the `auth` object's `type` value to `CODECONNECTIONS` in the `source` object.
	// - For source code in a Bitbucket repository, the HTTPS clone URL to the repository that contains the source and the buildspec file. You must connect your AWS account to your Bitbucket account. Use the AWS CodeBuild console to start creating a build project. When you use the console to connect (or reconnect) with Bitbucket, on the Bitbucket *Confirm access to your account* page, choose *Grant access* . (After you have connected to your Bitbucket account, you do not need to finish creating the build project. You can leave the AWS CodeBuild console.) To instruct AWS CodeBuild to use this connection, in the `source` object, set the `auth` object's `type` value to `OAUTH` .
	//
	// If you specify `CODEPIPELINE` for the `Type` property, don't specify this property. For all of the other types, you must specify `Location` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-location
	//
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Set to true to report the status of a build's start and finish to your source provider.
	//
	// This option is valid only when your source provider is GitHub, GitHub Enterprise, or Bitbucket. If this is set and you use a different source provider, an `invalidInputException` is thrown.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-reportbuildstatus
	//
	ReportBuildStatus interface{} `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// An identifier for this project source.
	//
	// The identifier can only contain alphanumeric characters and underscores, and must be less than 128 characters in length.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html#cfn-codebuild-project-source-sourceidentifier
	//
	SourceIdentifier *string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

