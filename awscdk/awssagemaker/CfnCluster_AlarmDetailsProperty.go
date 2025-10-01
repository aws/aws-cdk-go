package awssagemaker


// The details of the alarm to monitor during the AMI update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmDetailsProperty := &AlarmDetailsProperty{
//   	AlarmName: jsii.String("alarmName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-alarmdetails.html
//
type CfnCluster_AlarmDetailsProperty struct {
	// The name of the alarm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-alarmdetails.html#cfn-sagemaker-cluster-alarmdetails-alarmname
	//
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
}

