package awsquicksight


// Determines the row alternate color options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowAlternateColorOptionsProperty := &RowAlternateColorOptionsProperty{
//   	RowAlternateColors: []*string{
//   		jsii.String("rowAlternateColors"),
//   	},
//   	Status: jsii.String("status"),
//   	UsePrimaryBackgroundColor: jsii.String("usePrimaryBackgroundColor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rowalternatecoloroptions.html
//
type CfnTemplate_RowAlternateColorOptionsProperty struct {
	// Determines the list of row alternate colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rowalternatecoloroptions.html#cfn-quicksight-template-rowalternatecoloroptions-rowalternatecolors
	//
	RowAlternateColors *[]*string `field:"optional" json:"rowAlternateColors" yaml:"rowAlternateColors"`
	// Determines the widget status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rowalternatecoloroptions.html#cfn-quicksight-template-rowalternatecoloroptions-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The primary background color options for alternate rows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rowalternatecoloroptions.html#cfn-quicksight-template-rowalternatecoloroptions-useprimarybackgroundcolor
	//
	UsePrimaryBackgroundColor *string `field:"optional" json:"usePrimaryBackgroundColor" yaml:"usePrimaryBackgroundColor"`
}

