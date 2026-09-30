package awsneptunegraph


// Options for importing data from a Neptune database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   neptuneImportOptionsProperty := &NeptuneImportOptionsProperty{
//   	S3ExportKmsKeyId: jsii.String("s3ExportKmsKeyId"),
//   	S3ExportPath: jsii.String("s3ExportPath"),
//
//   	// the properties below are optional
//   	PreserveDefaultVertexLabels: jsii.Boolean(false),
//   	PreserveEdgeIds: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-neptuneimportoptions.html
//
type CfnGraph_NeptuneImportOptionsProperty struct {
	// The KMS key to use to encrypt data in the S3 bucket where the graph data is exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-neptuneimportoptions.html#cfn-neptunegraph-graph-neptuneimportoptions-s3exportkmskeyid
	//
	S3ExportKmsKeyId *string `field:"required" json:"s3ExportKmsKeyId" yaml:"s3ExportKmsKeyId"`
	// The path to an S3 bucket from which to import data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-neptuneimportoptions.html#cfn-neptunegraph-graph-neptuneimportoptions-s3exportpath
	//
	S3ExportPath *string `field:"required" json:"s3ExportPath" yaml:"s3ExportPath"`
	// Neptune Analytics supports label-less vertices and no labels are assigned unless one is explicitly provided.
	//
	// Neptune assigns default labels when none is explicitly provided. When importing the data into Neptune Analytics, the default vertex labels can be omitted by setting preserveDefaultVertexLabels to false. Note that if the vertex only has default labels, and has no other properties or edges, then the vertex will effectively not get imported into Neptune Analytics when preserveDefaultVertexLabels is set to false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-neptuneimportoptions.html#cfn-neptunegraph-graph-neptuneimportoptions-preservedefaultvertexlabels
	//
	PreserveDefaultVertexLabels interface{} `field:"optional" json:"preserveDefaultVertexLabels" yaml:"preserveDefaultVertexLabels"`
	// Neptune Analytics currently does not support user defined edge ids.
	//
	// The edge ids are not imported by default. They are imported if preserveEdgeIds is set to true, and ids are stored as properties on the relationships with the property name neptuneEdgeId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptunegraph-graph-neptuneimportoptions.html#cfn-neptunegraph-graph-neptuneimportoptions-preserveedgeids
	//
	PreserveEdgeIds interface{} `field:"optional" json:"preserveEdgeIds" yaml:"preserveEdgeIds"`
}

