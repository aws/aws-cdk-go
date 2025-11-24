package mixinsawsapptest


// Specifies an output file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputFileProperty := &OutputFileProperty{
//   	FileLocation: jsii.String("fileLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-outputfile.html
//
type CfnTestCasePropsMixin_OutputFileProperty struct {
	// The file location of the output file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-outputfile.html#cfn-apptest-testcase-outputfile-filelocation
	//
	FileLocation *string `field:"optional" json:"fileLocation" yaml:"fileLocation"`
}

