package awsdeadline


// Properties for defining a `CfnStorageProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStorageProfileProps := &CfnStorageProfileProps{
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	OsFamily: jsii.String("osFamily"),
//
//   	// the properties below are optional
//   	FileSystemLocations: []interface{}{
//   		&FileSystemLocationProperty{
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html
//
type CfnStorageProfileProps struct {
	// The display name of the storage profile summary to update.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html#cfn-deadline-storageprofile-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The unique identifier of the farm that contains the storage profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html#cfn-deadline-storageprofile-farmid
	//
	FarmId *string `field:"required" json:"farmId" yaml:"farmId"`
	// The operating system (OS) family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html#cfn-deadline-storageprofile-osfamily
	//
	OsFamily *string `field:"required" json:"osFamily" yaml:"osFamily"`
	// Operating system specific file system path to the storage location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-storageprofile.html#cfn-deadline-storageprofile-filesystemlocations
	//
	FileSystemLocations interface{} `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
}

