package previewawsquicksightmixins


// The configuration settings for the email alerts that are sent when a dataset refresh fails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   refreshFailureEmailAlertProperty := &RefreshFailureEmailAlertProperty{
//   	AlertStatus: jsii.String("alertStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshfailureemailalert.html
//
type CfnDataSetPropsMixin_RefreshFailureEmailAlertProperty struct {
	// The status value that determines if email alerts are sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshfailureemailalert.html#cfn-quicksight-dataset-refreshfailureemailalert-alertstatus
	//
	AlertStatus *string `field:"optional" json:"alertStatus" yaml:"alertStatus"`
}

