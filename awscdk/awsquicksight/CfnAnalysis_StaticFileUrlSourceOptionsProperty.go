package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticFileUrlSourceOptionsProperty := &StaticFileUrlSourceOptionsProperty{
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfileurlsourceoptions.html
//
type CfnAnalysis_StaticFileUrlSourceOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-staticfileurlsourceoptions.html#cfn-quicksight-analysis-staticfileurlsourceoptions-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
}

