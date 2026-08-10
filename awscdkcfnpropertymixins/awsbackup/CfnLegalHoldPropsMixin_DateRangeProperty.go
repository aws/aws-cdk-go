package awsbackup


// A date range for filtering recovery points.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dateRangeProperty := &DateRangeProperty{
//   	FromDate: jsii.String("fromDate"),
//   	ToDate: jsii.String("toDate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-daterange.html
//
type CfnLegalHoldPropsMixin_DateRangeProperty struct {
	// The beginning date, inclusive.
	//
	// ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-daterange.html#cfn-backup-legalhold-daterange-fromdate
	//
	FromDate *string `field:"optional" json:"fromDate" yaml:"fromDate"`
	// The end date, inclusive.
	//
	// ISO 8601 date-time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-legalhold-daterange.html#cfn-backup-legalhold-daterange-todate
	//
	ToDate *string `field:"optional" json:"toDate" yaml:"toDate"`
}

