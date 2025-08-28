package awsarczonalshift


// Properties for defining a `CfnAutoshiftObserverNotificationStatus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoshiftObserverNotificationStatusProps := &CfnAutoshiftObserverNotificationStatusProps{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-autoshiftobservernotificationstatus.html
//
type CfnAutoshiftObserverNotificationStatusProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-arczonalshift-autoshiftobservernotificationstatus.html#cfn-arczonalshift-autoshiftobservernotificationstatus-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
}

