package previewawsquicksightmixins


// The navigation operation that navigates between different sheets in the same analysis.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customActionNavigationOperationProperty := &CustomActionNavigationOperationProperty{
//   	LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   		TargetSheetId: jsii.String("targetSheetId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customactionnavigationoperation.html
//
type CfnDashboardPropsMixin_CustomActionNavigationOperationProperty struct {
	// The configuration that chooses the navigation target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customactionnavigationoperation.html#cfn-quicksight-dashboard-customactionnavigationoperation-localnavigationconfiguration
	//
	LocalNavigationConfiguration interface{} `field:"optional" json:"localNavigationConfiguration" yaml:"localNavigationConfiguration"`
}

