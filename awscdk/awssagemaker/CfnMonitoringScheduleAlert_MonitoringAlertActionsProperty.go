package awssagemaker


// A list of alert actions taken in response to an alert going into InAlert status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   monitoringAlertActionsProperty := &MonitoringAlertActionsProperty{
//   	ModelDashboardIndicator: &ModelDashboardIndicatorActionProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedulealert-monitoringalertactions.html
//
type CfnMonitoringScheduleAlert_MonitoringAlertActionsProperty struct {
	// An alert action taken to light up an icon on the Amazon SageMaker Model Dashboard when an alert goes into InAlert status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedulealert-monitoringalertactions.html#cfn-sagemaker-monitoringschedulealert-monitoringalertactions-modeldashboardindicator
	//
	ModelDashboardIndicator interface{} `field:"optional" json:"modelDashboardIndicator" yaml:"modelDashboardIndicator"`
}

