package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// FileSystemLocation provider definition for a CodeBuild Project.
//
// Example:
//   codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	fileSystemLocations: []iFileSystemLocation{
//   		codebuild.fileSystemLocation.efs(&efsFileSystemLocationProps{
//   			identifier: jsii.String("myidentifier2"),
//   			location: jsii.String("myclodation.mydnsroot.com:/loc"),
//   			mountPoint: jsii.String("/media"),
//   			mountOptions: jsii.String("opts"),
//   		}),
//   	},
//   })
//
type FileSystemLocation interface {
}

// The jsii proxy struct for FileSystemLocation
type jsiiProxy_FileSystemLocation struct {
	_ byte // padding
}

func NewFileSystemLocation() FileSystemLocation {
	_init_.Initialize()

	j := jsiiProxy_FileSystemLocation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.FileSystemLocation",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFileSystemLocation_Override(f FileSystemLocation) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.FileSystemLocation",
		nil, // no parameters
		f,
	)
}

// EFS file system provider.
func FileSystemLocation_Efs(props *EfsFileSystemLocationProps) IFileSystemLocation {
	_init_.Initialize()

	var returns IFileSystemLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.FileSystemLocation",
		"efs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

