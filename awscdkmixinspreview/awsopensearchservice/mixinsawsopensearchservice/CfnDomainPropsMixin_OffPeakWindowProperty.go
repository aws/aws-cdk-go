package mixinsawsopensearchservice


// A custom 10-hour, low-traffic window during which OpenSearch Service can perform mandatory configuration changes on the domain.
//
// These actions can include scheduled service software updates and blue/green Auto-Tune enhancements. OpenSearch Service will schedule these actions during the window that you specify. If you don't specify a window start time, it defaults to 10:00 P.M. local time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   offPeakWindowProperty := &OffPeakWindowProperty{
//   	WindowStartTime: &WindowStartTimeProperty{
//   		Hours: jsii.Number(123),
//   		Minutes: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-offpeakwindow.html
//
type CfnDomainPropsMixin_OffPeakWindowProperty struct {
	// The desired start time for an off-peak maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-offpeakwindow.html#cfn-opensearchservice-domain-offpeakwindow-windowstarttime
	//
	WindowStartTime interface{} `field:"optional" json:"windowStartTime" yaml:"windowStartTime"`
}

