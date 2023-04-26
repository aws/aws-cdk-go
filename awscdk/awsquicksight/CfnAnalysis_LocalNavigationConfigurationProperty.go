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
type CfnAnalysis_LocalNavigationConfigurationProperty struct {
	// The sheet that is targeted for navigation in the same analysis.
	TargetSheetId *string `field:"required" json:"targetSheetId" yaml:"targetSheetId"`
}

