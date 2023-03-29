package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionNavigationOperationProperty := &CustomActionNavigationOperationProperty{
//   	LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   		TargetSheetId: jsii.String("targetSheetId"),
//   	},
//   }
//
type CfnDashboard_CustomActionNavigationOperationProperty struct {
	// `CfnDashboard.CustomActionNavigationOperationProperty.LocalNavigationConfiguration`.
	LocalNavigationConfiguration interface{} `field:"optional" json:"localNavigationConfiguration" yaml:"localNavigationConfiguration"`
}

