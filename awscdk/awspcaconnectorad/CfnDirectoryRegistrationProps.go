package awspcaconnectorad


// Properties for defining a `CfnDirectoryRegistration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDirectoryRegistrationProps := &CfnDirectoryRegistrationProps{
//   	DirectoryId: jsii.String("directoryId"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-directoryregistration.html
//
type CfnDirectoryRegistrationProps struct {
	// The identifier of the Active Directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-directoryregistration.html#cfn-pcaconnectorad-directoryregistration-directoryid
	//
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Metadata assigned to a directory registration consisting of a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorad-directoryregistration.html#cfn-pcaconnectorad-directoryregistration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

