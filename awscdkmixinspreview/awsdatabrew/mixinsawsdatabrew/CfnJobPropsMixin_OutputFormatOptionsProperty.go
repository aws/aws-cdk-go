package mixinsawsdatabrew


// Represents a set of options that define the structure of comma-separated (CSV) job output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputFormatOptionsProperty := &OutputFormatOptionsProperty{
//   	Csv: &CsvOutputOptionsProperty{
//   		Delimiter: jsii.String("delimiter"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-outputformatoptions.html
//
type CfnJobPropsMixin_OutputFormatOptionsProperty struct {
	// Represents a set of options that define the structure of comma-separated value (CSV) job output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-outputformatoptions.html#cfn-databrew-job-outputformatoptions-csv
	//
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
}

