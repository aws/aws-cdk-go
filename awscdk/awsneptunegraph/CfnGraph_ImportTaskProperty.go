package awsneptunegraph


// The import task details to import data into the graph at creation time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   importTaskProperty := &ImportTaskProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	BlankNodeHandling: jsii.String("blankNodeHandling"),
//   	FailOnError: jsii.Boolean(false),
//   	Format: jsii.String("format"),
//   	ImportOptions: &ImportOptionsProperty{
//   		Neptune: &NeptuneImportOptionsProperty{
//   			S3ExportKmsKeyId: jsii.String("s3ExportKmsKeyId"),
//   			S3ExportPath: jsii.String("s3ExportPath"),
//
//   			// the properties below are optional
//   			PreserveDefaultVertexLabels: jsii.Boolean(false),
//   			PreserveEdgeIds: jsii.Boolean(false),
//   		},
//   	},
//   	MaxProvisionedMemory: jsii.Number(123),
//   	MinProvisionedMemory: jsii.Number(123),
//   	ParquetType: jsii.String("parquetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html
//
type CfnGraph_ImportTaskProperty struct {
	// The ARN of the IAM role that will allow access to the data that is to be imported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A URL identifying to the location of the data to be imported.
	//
	// This can be an Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The method to handle blank nodes in the dataset.
	//
	// Currently, only convertToIri is supported, meaning blank nodes are converted to unique IRIs at load time. Must be provided when format is NTRIPLES
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-blanknodehandling
	//
	BlankNodeHandling *string `field:"optional" json:"blankNodeHandling" yaml:"blankNodeHandling"`
	// If set to true, the task halts when an import error is encountered.
	//
	// If set to false, the task skips the data that caused the error and continues if possible.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-failonerror
	//
	FailOnError interface{} `field:"optional" json:"failOnError" yaml:"failOnError"`
	// Specifies the format of S3 data to be imported.
	//
	// Valid values are CSV, which identifies the Gremlin CSV format, OPEN_CYPHER, which identifies the openCypher load format, or NTRIPLES, which identifies the RDF n-triples format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Contains options for controlling the import process.
	//
	// For example, if the failOnError key is set to false, the import skips the data that caused the error and continues if possible (whereas if set to true, the default, or if omitted, the import operation halts immediately when an error is encountered).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-importoptions
	//
	ImportOptions interface{} `field:"optional" json:"importOptions" yaml:"importOptions"`
	// The maximum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.
	//
	// Default: 1024, or the approved upper limit for your account. If both the minimum and maximum values are specified, the final provisioned-memory will be chosen per the actual size of your imported data. If neither value is specified, 128 m-NCUs are used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-maxprovisionedmemory
	//
	MaxProvisionedMemory *float64 `field:"optional" json:"maxProvisionedMemory" yaml:"maxProvisionedMemory"`
	// The minimum provisioned memory-optimized Neptune Capacity Units (m-NCUs) to use for the graph.
	//
	// Default: 16.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-minprovisionedmemory
	//
	MinProvisionedMemory *float64 `field:"optional" json:"minProvisionedMemory" yaml:"minProvisionedMemory"`
	// The parquet type of the import task.
	//
	// Required when Format is PARQUET.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-importtask.html#cfn-neptunegraph-graph-importtask-parquettype
	//
	ParquetType *string `field:"optional" json:"parquetType" yaml:"parquetType"`
}

