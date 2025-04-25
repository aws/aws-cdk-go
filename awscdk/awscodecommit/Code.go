package awscodecommit

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the contents to initialize the repository with.
//
// Example:
//   repo := codecommit.NewRepository(this, jsii.String("Repository"), &RepositoryProps{
//   	RepositoryName: jsii.String("MyRepositoryName"),
//   	Code: codecommit.Code_FromDirectory(path.join(__dirname, jsii.String("directory/")), jsii.String("develop")),
//   })
//
type Code interface {
	// This method is called after a repository is passed this instance of Code in its 'code' property.
	Bind(scope constructs.Construct) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codecommit.Code",
		nil, // no parameters
		c,
	)
}

// Code from user-supplied asset.
func Code_FromAsset(asset awss3assets.Asset, branch *string) Code {
	_init_.Initialize()

	if err := validateCode_FromAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codecommit.Code",
		"fromAsset",
		[]interface{}{asset, branch},
		&returns,
	)

	return returns
}

// Code from directory.
func Code_FromDirectory(directoryPath *string, branch *string) Code {
	_init_.Initialize()

	if err := validateCode_FromDirectoryParameters(directoryPath); err != nil {
		panic(err)
	}
	var returns Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codecommit.Code",
		"fromDirectory",
		[]interface{}{directoryPath, branch},
		&returns,
	)

	return returns
}

// Code from preexisting ZIP file.
func Code_FromZipFile(filePath *string, branch *string) Code {
	_init_.Initialize()

	if err := validateCode_FromZipFileParameters(filePath); err != nil {
		panic(err)
	}
	var returns Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codecommit.Code",
		"fromZipFile",
		[]interface{}{filePath, branch},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	if err := c.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

