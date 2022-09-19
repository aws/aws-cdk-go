package awscodebuild


// Construction properties for {@link EfsFileSystemLocation}.
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
type EfsFileSystemLocationProps struct {
	// The name used to access a file system created by Amazon EFS.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// A string that specifies the location of the file system, like Amazon EFS.
	//
	// This value looks like `fs-abcd1234.efs.us-west-2.amazonaws.com:/my-efs-mount-directory`.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The location in the container where you mount the file system.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// The mount options for a file system such as Amazon EFS.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

