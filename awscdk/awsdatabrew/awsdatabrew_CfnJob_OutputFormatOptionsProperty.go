package awsdatabrew


// Represents a set of options that define the structure of comma-separated (CSV) job output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputFormatOptionsProperty := &outputFormatOptionsProperty{
//   	csv: &csvOutputOptionsProperty{
//   		delimiter: jsii.String("delimiter"),
//   	},
//   }
//
type CfnJob_OutputFormatOptionsProperty struct {
	// Represents a set of options that define the structure of comma-separated value (CSV) job output.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
}

