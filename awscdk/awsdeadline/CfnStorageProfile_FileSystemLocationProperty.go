package awsdeadline


// The details of the file system location for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemLocationProperty := &FileSystemLocationProperty{
//   	Name: jsii.String("name"),
//   	Path: jsii.String("path"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-storageprofile-filesystemlocation.html
//
type CfnStorageProfile_FileSystemLocationProperty struct {
	// The location name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-storageprofile-filesystemlocation.html#cfn-deadline-storageprofile-filesystemlocation-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The file path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-storageprofile-filesystemlocation.html#cfn-deadline-storageprofile-filesystemlocation-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// The type of file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-storageprofile-filesystemlocation.html#cfn-deadline-storageprofile-filesystemlocation-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

