package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshFailureEmailAlertProperty := &RefreshFailureEmailAlertProperty{
//   	AlertStatus: jsii.String("alertStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshfailureemailalert.html
//
type CfnDataSet_RefreshFailureEmailAlertProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-refreshfailureemailalert.html#cfn-quicksight-dataset-refreshfailureemailalert-alertstatus
	//
	AlertStatus *string `field:"optional" json:"alertStatus" yaml:"alertStatus"`
}

