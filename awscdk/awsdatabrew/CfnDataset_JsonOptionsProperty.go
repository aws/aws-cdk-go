package awsdatabrew


// Represents the JSON-specific options that define how input is to be interpreted by AWS Glue DataBrew .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonOptionsProperty := &JsonOptionsProperty{
//   	MultiLine: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-jsonoptions.html
//
type CfnDataset_JsonOptionsProperty struct {
	// A value that specifies whether JSON input contains embedded new line characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-jsonoptions.html#cfn-databrew-dataset-jsonoptions-multiline
	//
	MultiLine interface{} `field:"optional" json:"multiLine" yaml:"multiLine"`
}

