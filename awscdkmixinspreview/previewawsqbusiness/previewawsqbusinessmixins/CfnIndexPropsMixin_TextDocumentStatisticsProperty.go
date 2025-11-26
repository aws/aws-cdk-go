package previewawsqbusinessmixins


// Provides information about text documents in an index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textDocumentStatisticsProperty := &TextDocumentStatisticsProperty{
//   	IndexedTextBytes: jsii.Number(123),
//   	IndexedTextDocumentCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-textdocumentstatistics.html
//
type CfnIndexPropsMixin_TextDocumentStatisticsProperty struct {
	// The total size, in bytes, of the indexed documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-textdocumentstatistics.html#cfn-qbusiness-index-textdocumentstatistics-indexedtextbytes
	//
	IndexedTextBytes *float64 `field:"optional" json:"indexedTextBytes" yaml:"indexedTextBytes"`
	// The number of text documents indexed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-index-textdocumentstatistics.html#cfn-qbusiness-index-textdocumentstatistics-indexedtextdocumentcount
	//
	IndexedTextDocumentCount *float64 `field:"optional" json:"indexedTextDocumentCount" yaml:"indexedTextDocumentCount"`
}

