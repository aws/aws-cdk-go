package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Create files on the EC2 instance.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//
//   	Init: ec2.CloudFormationInit_FromElements(ec2.InitService_SystemdConfigFile(jsii.String("simpleserver"), &SystemdConfigFileOptions{
//   		Command: jsii.String("/usr/bin/python3 -m http.server 8080"),
//   		Cwd: jsii.String("/var/www/html"),
//   	}), ec2.InitService_Enable(jsii.String("simpleserver"), &InitServiceOptions{
//   		ServiceManager: ec2.ServiceManager_SYSTEMD,
//   	}), ec2.InitFile_FromString(jsii.String("/var/www/html/index.html"), jsii.String("Hello! It's working!"))),
//   })
//
type InitFile interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitFile
type jsiiProxy_InitFile struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitFile) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


func NewInitFile_Override(i InitFile, fileName *string, options *InitFileOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InitFile",
		[]interface{}{fileName, options},
		i,
	)
}

// Create an asset from the given file.
//
// This is appropriate for files that are too large to embed into the template.
func InitFile_FromAsset(targetFileName *string, path *string, options *InitFileAssetOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromAssetParameters(targetFileName, path, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromAsset",
		[]interface{}{targetFileName, path, options},
		&returns,
	)

	return returns
}

// Use a file from an asset at instance startup time.
func InitFile_FromExistingAsset(targetFileName *string, asset awss3assets.Asset, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromExistingAssetParameters(targetFileName, asset, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromExistingAsset",
		[]interface{}{targetFileName, asset, options},
		&returns,
	)

	return returns
}

// Read a file from disk and use its contents.
//
// The file will be embedded in the template, so care should be taken to not
// exceed the template size.
//
// If options.base64encoded is set to true, this will base64-encode the file's contents.
func InitFile_FromFileInline(targetFileName *string, sourceFileName *string, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromFileInlineParameters(targetFileName, sourceFileName, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromFileInline",
		[]interface{}{targetFileName, sourceFileName, options},
		&returns,
	)

	return returns
}

// Use a JSON-compatible object as the file content, write it to a JSON file.
//
// May contain tokens.
func InitFile_FromObject(fileName *string, obj *map[string]interface{}, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromObjectParameters(fileName, obj, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromObject",
		[]interface{}{fileName, obj, options},
		&returns,
	)

	return returns
}

// Download a file from an S3 bucket at instance startup time.
func InitFile_FromS3Object(fileName *string, bucket awss3.IBucket, key *string, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromS3ObjectParameters(fileName, bucket, key, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromS3Object",
		[]interface{}{fileName, bucket, key, options},
		&returns,
	)

	return returns
}

// Use a literal string as the file content.
func InitFile_FromString(fileName *string, content *string, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromStringParameters(fileName, content, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromString",
		[]interface{}{fileName, content, options},
		&returns,
	)

	return returns
}

// Download from a URL at instance startup time.
func InitFile_FromUrl(fileName *string, url *string, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_FromUrlParameters(fileName, url, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"fromUrl",
		[]interface{}{fileName, url, options},
		&returns,
	)

	return returns
}

// Write a symlink with the given symlink target.
func InitFile_Symlink(fileName *string, target *string, options *InitFileOptions) InitFile {
	_init_.Initialize()

	if err := validateInitFile_SymlinkParameters(fileName, target, options); err != nil {
		panic(err)
	}
	var returns InitFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitFile",
		"symlink",
		[]interface{}{fileName, target, options},
		&returns,
	)

	return returns
}

