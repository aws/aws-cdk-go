package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIVisualStandardLayoutProperty := &KPIVisualStandardLayoutProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisualstandardlayout.html
//
type CfnTemplate_KPIVisualStandardLayoutProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-kpivisualstandardlayout.html#cfn-quicksight-template-kpivisualstandardlayout-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

