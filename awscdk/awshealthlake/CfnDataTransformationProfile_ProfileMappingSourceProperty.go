package awshealthlake


// Create the profile from raw Velocity template mapping content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileMappingSourceProperty := &ProfileMappingSourceProperty{
//   	ProfileMapping: map[string]*string{
//   		"profileMappingKey": jsii.String("profileMapping"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-profilemappingsource.html
//
type CfnDataTransformationProfile_ProfileMappingSourceProperty struct {
	// Map of template file paths to their Velocity template content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-profilemappingsource.html#cfn-healthlake-datatransformationprofile-profilemappingsource-profilemapping
	//
	ProfileMapping interface{} `field:"required" json:"profileMapping" yaml:"profileMapping"`
}

