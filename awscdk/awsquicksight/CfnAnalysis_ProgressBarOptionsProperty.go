package awsquicksight


// The options that determine the presentation of the progress bar of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   progressBarOptionsProperty := &ProgressBarOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnAnalysis_ProgressBarOptionsProperty struct {
	// The visibility of the progress bar.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

