package mixinsawsqbusiness


// Provides information about the number of documents in an index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   indexStatisticsProperty := &IndexStatisticsProperty{
//   	TextDocumentStatistics: &TextDocumentStatisticsProperty{
//   		IndexedTextBytes: jsii.Number(123),
//   		IndexedTextDocumentCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-indexstatistics.html
//
type CfnIndexPropsMixin_IndexStatisticsProperty struct {
	// The number of documents indexed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-indexstatistics.html#cfn-qbusiness-index-indexstatistics-textdocumentstatistics
	//
	TextDocumentStatistics interface{} `field:"optional" json:"textDocumentStatistics" yaml:"textDocumentStatistics"`
}

