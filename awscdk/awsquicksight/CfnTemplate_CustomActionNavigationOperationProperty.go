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
type CfnTemplate_CustomActionNavigationOperationProperty struct {
	// `CfnTemplate.CustomActionNavigationOperationProperty.LocalNavigationConfiguration`.
	LocalNavigationConfiguration interface{} `field:"optional" json:"localNavigationConfiguration" yaml:"localNavigationConfiguration"`
}

