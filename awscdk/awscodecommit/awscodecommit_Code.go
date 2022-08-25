package awscodecommit

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents the contents to initialize the repository with.
//
// Example:
//   repo := codecommit.NewRepository(this, jsii.String("Repository"), &repositoryProps{
//   	repositoryName: jsii.String("MyRepositoryName"),
//   	code: codecommit.code.fromDirectory(path.join(__dirname, jsii.String("directory/")), jsii.String("develop")),
//   })
//
// Experimental.
type Code interface {
	// This method is called after a repository is passed this instance of Code in its 'code' property.
	// Experimental.
	Bind(scope constructs.Construct) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

// Experimental.
func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codecommit.Code",
		nil, // no parameters
		c,
	)
}

// Code from user-supplied asset.
// Experimental.
func Code_FromAsset(asset awss3assets.Asset, branch *string) Code {
	_init_.Initialize()

	var returns Code

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Code",
		"fromAsset",
		[]interface{}{asset, branch},
		&returns,
	)

	return returns
}

// Code from directory.
// Experimental.
func Code_FromDirectory(directoryPath *string, branch *string) Code {
	_init_.Initialize()

	var returns Code

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Code",
		"fromDirectory",
		[]interface{}{directoryPath, branch},
		&returns,
	)

	return returns
}

// Code from preexisting ZIP file.
// Experimental.
func Code_FromZipFile(filePath *string, branch *string) Code {
	_init_.Initialize()

	var returns Code

	_jsii_.StaticInvoke(
		"monocdk.aws_codecommit.Code",
		"fromZipFile",
		[]interface{}{filePath, branch},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

