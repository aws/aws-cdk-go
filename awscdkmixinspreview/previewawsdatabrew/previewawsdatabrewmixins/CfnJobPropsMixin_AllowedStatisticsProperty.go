package previewawsdatabrewmixins


// Configuration of statistics that are allowed to be run on columns that contain detected entities.
//
// When undefined, no statistics will be computed on columns that contain detected entities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   allowedStatisticsProperty := &AllowedStatisticsProperty{
//   	Statistics: []*string{
//   		jsii.String("statistics"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-allowedstatistics.html
//
type CfnJobPropsMixin_AllowedStatisticsProperty struct {
	// One or more column statistics to allow for columns that contain detected entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-allowedstatistics.html#cfn-databrew-job-allowedstatistics-statistics
	//
	Statistics *[]*string `field:"optional" json:"statistics" yaml:"statistics"`
}

