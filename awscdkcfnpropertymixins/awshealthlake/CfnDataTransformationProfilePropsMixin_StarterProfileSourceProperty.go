package awshealthlake


// Create the profile from a predefined starter profile of transformation templates.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   starterProfileSourceProperty := &StarterProfileSourceProperty{
//   	StarterProfileName: jsii.String("starterProfileName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-starterprofilesource.html
//
type CfnDataTransformationProfilePropsMixin_StarterProfileSourceProperty struct {
	// The name of the starter profile to seed the profile from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-starterprofilesource.html#cfn-healthlake-datatransformationprofile-starterprofilesource-starterprofilename
	//
	StarterProfileName *string `field:"optional" json:"starterProfileName" yaml:"starterProfileName"`
}

