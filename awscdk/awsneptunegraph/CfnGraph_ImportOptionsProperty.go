package awsneptunegraph


// Contains options for controlling the import process.
//
// For example, if the failOnError key is set to false, the import skips the data that caused the error and continues if possible (whereas if set to true, the default, or if omitted, the import operation halts immediately when an error is encountered).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   importOptionsProperty := &ImportOptionsProperty{
//   	Neptune: &NeptuneImportOptionsProperty{
//   		S3ExportKmsKeyId: jsii.String("s3ExportKmsKeyId"),
//   		S3ExportPath: jsii.String("s3ExportPath"),
//
//   		// the properties below are optional
//   		PreserveDefaultVertexLabels: jsii.Boolean(false),
//   		PreserveEdgeIds: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importoptions.html
//
type CfnGraph_ImportOptionsProperty struct {
	// Options for importing data from a Neptune database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importoptions.html#cfn-neptunegraph-graph-importoptions-neptune
	//
	Neptune interface{} `field:"optional" json:"neptune" yaml:"neptune"`
}

