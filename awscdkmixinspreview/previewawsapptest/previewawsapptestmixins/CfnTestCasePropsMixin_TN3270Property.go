package previewawsapptestmixins


// Specifies the TN3270 protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tN3270Property := &TN3270Property{
//   	ExportDataSetNames: []*string{
//   		jsii.String("exportDataSetNames"),
//   	},
//   	Script: &ScriptProperty{
//   		ScriptLocation: jsii.String("scriptLocation"),
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-tn3270.html
//
type CfnTestCasePropsMixin_TN3270Property struct {
	// The data set names of the TN3270 protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-tn3270.html#cfn-apptest-testcase-tn3270-exportdatasetnames
	//
	ExportDataSetNames *[]*string `field:"optional" json:"exportDataSetNames" yaml:"exportDataSetNames"`
	// The script of the TN3270 protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-tn3270.html#cfn-apptest-testcase-tn3270-script
	//
	Script interface{} `field:"optional" json:"script" yaml:"script"`
}

