package mixinsawsquicksight


// The navigation configuration for `CustomActionNavigationOperation` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   localNavigationConfigurationProperty := &LocalNavigationConfigurationProperty{
//   	TargetSheetId: jsii.String("targetSheetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-localnavigationconfiguration.html
//
type CfnAnalysisPropsMixin_LocalNavigationConfigurationProperty struct {
	// The sheet that is targeted for navigation in the same analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-localnavigationconfiguration.html#cfn-quicksight-analysis-localnavigationconfiguration-targetsheetid
	//
	TargetSheetId *string `field:"optional" json:"targetSheetId" yaml:"targetSheetId"`
}

