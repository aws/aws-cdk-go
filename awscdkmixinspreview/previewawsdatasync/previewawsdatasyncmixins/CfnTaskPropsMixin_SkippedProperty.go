package previewawsdatasyncmixins


// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to skip during your transfer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   skippedProperty := &SkippedProperty{
//   	ReportLevel: jsii.String("reportLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-skipped.html
//
type CfnTaskPropsMixin_SkippedProperty struct {
	// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-skipped.html#cfn-datasync-task-skipped-reportlevel
	//
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
}

