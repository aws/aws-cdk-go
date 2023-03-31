package awsappflow


// Specifies the configuration used when importing incremental records from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incrementalPullConfigProperty := &incrementalPullConfigProperty{
//   	datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   }
//
type CfnFlow_IncrementalPullConfigProperty struct {
	// A field that specifies the date time or timestamp field as the criteria to use when importing incremental records from the source.
	DatetimeTypeFieldName *string `field:"optional" json:"datetimeTypeFieldName" yaml:"datetimeTypeFieldName"`
}

