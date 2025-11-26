package previewawsdatazonemixins


// The AWS Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   regionProperty := &RegionProperty{
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-region.html
//
type CfnProjectProfilePropsMixin_RegionProperty struct {
	// The AWS Region name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-region.html#cfn-datazone-projectprofile-region-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

