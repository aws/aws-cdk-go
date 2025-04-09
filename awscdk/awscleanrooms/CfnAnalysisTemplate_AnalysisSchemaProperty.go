package awscleanrooms


// A relation within an analysis.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSchemaProperty := &AnalysisSchemaProperty{
//   	ReferencedTables: []*string{
//   		jsii.String("referencedTables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisschema.html
//
type CfnAnalysisTemplate_AnalysisSchemaProperty struct {
	// The tables referenced in the analysis schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysisschema.html#cfn-cleanrooms-analysistemplate-analysisschema-referencedtables
	//
	ReferencedTables *[]*string `field:"required" json:"referencedTables" yaml:"referencedTables"`
}

