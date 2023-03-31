package awskinesisanalyticsv2


// For a SQL-based Kinesis Data Analytics application, provides additional mapping information when JSON is the record format on the streaming source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jSONMappingParametersProperty := &jSONMappingParametersProperty{
//   	recordRowPath: jsii.String("recordRowPath"),
//   }
//
type CfnApplication_JSONMappingParametersProperty struct {
	// The path to the top-level parent that contains the records.
	RecordRowPath *string `field:"required" json:"recordRowPath" yaml:"recordRowPath"`
}

