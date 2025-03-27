package awsappflow


// Specifies the configuration used when importing incremental records from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incrementalPullConfigProperty := &IncrementalPullConfigProperty{
//   	DatetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-incrementalpullconfig.html
//
type CfnFlow_IncrementalPullConfigProperty struct {
	// A field that specifies the date time or timestamp field as the criteria to use when importing incremental records from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-incrementalpullconfig.html#cfn-appflow-flow-incrementalpullconfig-datetimetypefieldname
	//
	DatetimeTypeFieldName *string `field:"optional" json:"datetimeTypeFieldName" yaml:"datetimeTypeFieldName"`
}

