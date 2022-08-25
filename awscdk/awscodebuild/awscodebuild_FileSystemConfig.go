package awscodebuild


// The type returned from {@link IFileSystemLocation#bind}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfig := &fileSystemConfig{
//   	location: &projectFileSystemLocationProperty{
//   		identifier: jsii.String("identifier"),
//   		location: jsii.String("location"),
//   		mountPoint: jsii.String("mountPoint"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		mountOptions: jsii.String("mountOptions"),
//   	},
//   }
//
type FileSystemConfig struct {
	// File system location wrapper property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectfilesystemlocation.html
	//
	Location *CfnProject_ProjectFileSystemLocationProperty `field:"required" json:"location" yaml:"location"`
}

