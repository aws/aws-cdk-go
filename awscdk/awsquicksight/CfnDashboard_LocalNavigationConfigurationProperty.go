package awsquicksight


// The navigation configuration for `CustomActionNavigationOperation` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localNavigationConfigurationProperty := &LocalNavigationConfigurationProperty{
//   	TargetSheetId: jsii.String("targetSheetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-localnavigationconfiguration.html
//
type CfnDashboard_LocalNavigationConfigurationProperty struct {
	// The sheet that is targeted for navigation in the same analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-localnavigationconfiguration.html#cfn-quicksight-dashboard-localnavigationconfiguration-targetsheetid
	//
	TargetSheetId *string `field:"required" json:"targetSheetId" yaml:"targetSheetId"`
}

