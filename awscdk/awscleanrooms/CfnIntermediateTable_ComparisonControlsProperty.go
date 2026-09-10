package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   comparisonControlsProperty := &ComparisonControlsProperty{
//   	AllowedColumnComparisonColumns: []*string{
//   		jsii.String("allowedColumnComparisonColumns"),
//   	},
//   	AllowedLiteralComparisonColumns: []*string{
//   		jsii.String("allowedLiteralComparisonColumns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-comparisoncontrols.html
//
type CfnIntermediateTable_ComparisonControlsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-comparisoncontrols.html#cfn-cleanrooms-intermediatetable-comparisoncontrols-allowedcolumncomparisoncolumns
	//
	AllowedColumnComparisonColumns *[]*string `field:"required" json:"allowedColumnComparisonColumns" yaml:"allowedColumnComparisonColumns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-comparisoncontrols.html#cfn-cleanrooms-intermediatetable-comparisoncontrols-allowedliteralcomparisoncolumns
	//
	AllowedLiteralComparisonColumns *[]*string `field:"required" json:"allowedLiteralComparisonColumns" yaml:"allowedLiteralComparisonColumns"`
}

