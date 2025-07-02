package awsdatazone


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionProperty := &RegionProperty{
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-region.html
//
type CfnProjectProfile_RegionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-projectprofile-region.html#cfn-datazone-projectprofile-region-regionname
	//
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

