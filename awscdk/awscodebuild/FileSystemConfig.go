package awscodebuild


// The type returned from `IFileSystemLocation#bind`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfig := &FileSystemConfig{
//   	Location: &ProjectFileSystemLocationProperty{
//   		Identifier: jsii.String("identifier"),
//   		Location: jsii.String("location"),
//   		MountPoint: jsii.String("mountPoint"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		MountOptions: jsii.String("mountOptions"),
//   	},
//   }
//
type FileSystemConfig struct {
	// File system location wrapper property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectfilesystemlocation.html
	//
	Location *CfnProject_ProjectFileSystemLocationProperty `field:"required" json:"location" yaml:"location"`
}

