package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ontapUnixFileSystemUserProperty := &OntapUnixFileSystemUserProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapunixfilesystemuser.html
//
type CfnS3AccessPointAttachment_OntapUnixFileSystemUserProperty struct {
	// The name of the UNIX user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-ontapunixfilesystemuser.html#cfn-fsx-s3accesspointattachment-ontapunixfilesystemuser-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

