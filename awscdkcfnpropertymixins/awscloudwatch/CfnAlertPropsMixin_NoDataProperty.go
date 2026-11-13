package awscloudwatch


// How the alert behaves when an evaluation returns no data.
//
// If omitted, the alert reports NODATA.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   noDataProperty := &NoDataProperty{
//   	TreatAs: jsii.String("treatAs"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-nodata.html
//
type CfnAlertPropsMixin_NoDataProperty struct {
	// The state to report when an evaluation returns no data.
	//
	// Valid values: OK, WARNING, CRITICAL, NODATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-alert-nodata.html#cfn-cloudwatch-alert-nodata-treatas
	//
	TreatAs *string `field:"optional" json:"treatAs" yaml:"treatAs"`
}

