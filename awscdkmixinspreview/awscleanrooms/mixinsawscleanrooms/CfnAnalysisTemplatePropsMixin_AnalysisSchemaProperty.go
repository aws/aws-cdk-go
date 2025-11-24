package mixinsawscleanrooms


// A relation within an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analysisSchemaProperty := &AnalysisSchemaProperty{
//   	ReferencedTables: []*string{
//   		jsii.String("referencedTables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisschema.html
//
type CfnAnalysisTemplatePropsMixin_AnalysisSchemaProperty struct {
	// The tables referenced in the analysis schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisschema.html#cfn-cleanrooms-analysistemplate-analysisschema-referencedtables
	//
	ReferencedTables *[]*string `field:"optional" json:"referencedTables" yaml:"referencedTables"`
}

