package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Source provider definition for a CodeBuild Project.
//
// Example:
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromSourceFilename(jsii.String("my-buildspec.yml")),
//   	Source: codebuild.Source_GitHub(&GitHubSourceProps{
//   		Owner: jsii.String("awslabs"),
//   		Repo: jsii.String("aws-cdk"),
//   	}),
//   })
//
type Source interface {
	ISource
	BadgeSupported() *bool
	Identifier() *string
	Type() *string
	// Called by the project when the source is added so that the source can perform binding operations on the source.
	//
	// For example, it can grant permissions to the
	// code build project to read from the S3 bucket.
	Bind(_scope constructs.Construct, _project IProject) *SourceConfig
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_Source) BadgeSupported() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"badgeSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Source) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Source) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewSource_Override(s Source, props *SourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Source",
		[]interface{}{props},
		s,
	)
}

func Source_BitBucket(props *BitBucketSourceProps) ISource {
	_init_.Initialize()

	if err := validateSource_BitBucketParameters(props); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"bitBucket",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_CodeCommit(props *CodeCommitSourceProps) ISource {
	_init_.Initialize()

	if err := validateSource_CodeCommitParameters(props); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"codeCommit",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_GitHub(props *GitHubSourceProps) ISource {
	_init_.Initialize()

	if err := validateSource_GitHubParameters(props); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"gitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_GitHubEnterprise(props *GitHubEnterpriseSourceProps) ISource {
	_init_.Initialize()

	if err := validateSource_GitHubEnterpriseParameters(props); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"gitHubEnterprise",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_S3(props *S3SourceProps) ISource {
	_init_.Initialize()

	if err := validateSource_S3Parameters(props); err != nil {
		panic(err)
	}
	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"s3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Source) Bind(_scope constructs.Construct, _project IProject) *SourceConfig {
	if err := s.validateBindParameters(_scope, _project); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _project},
		&returns,
	)

	return returns
}

