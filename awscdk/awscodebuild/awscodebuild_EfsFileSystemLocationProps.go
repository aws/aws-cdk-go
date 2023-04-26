package awscodebuild


// Construction properties for {@link EfsFileSystemLocation}.
//
// Example:
//   codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	FileSystemLocations: []iFileSystemLocation{
//   		codebuild.FileSystemLocation_Efs(&EfsFileSystemLocationProps{
//   			Identifier: jsii.String("myidentifier2"),
//   			Location: jsii.String("myclodation.mydnsroot.com:/loc"),
//   			MountPoint: jsii.String("/media"),
//   			MountOptions: jsii.String("opts"),
//   		}),
//   	},
//   })
//
// Experimental.
type EfsFileSystemLocationProps struct {
	// The name used to access a file system created by Amazon EFS.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// A string that specifies the location of the file system, like Amazon EFS.
	//
	// This value looks like `fs-abcd1234.efs.us-west-2.amazonaws.com:/my-efs-mount-directory`.
	// Experimental.
	Location *string `field:"required" json:"location" yaml:"location"`
	// The location in the container where you mount the file system.
	// Experimental.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// The mount options for a file system such as Amazon EFS.
	// Experimental.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

