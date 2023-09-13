package awsdatasync


// The reporting level for the verified section of your DataSync task report.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   verifiedProperty := &VerifiedProperty{
//   	ReportLevel: jsii.String("reportLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-verified.html
//
type CfnTask_VerifiedProperty struct {
	// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.
	//
	// - `ERRORS_ONLY` : A report shows what DataSync was unable to verify.
	// - `SUCCESSES_AND_ERRORS` : A report shows what DataSync was able and unable to verify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-verified.html#cfn-datasync-task-verified-reportlevel
	//
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
}

