package awsdatabrew


// Represents a set of options that define how DataBrew will write a comma-separated value (CSV) file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   csvOutputOptionsProperty := &csvOutputOptionsProperty{
//   	delimiter: jsii.String("delimiter"),
//   }
//
type CfnJob_CsvOutputOptionsProperty struct {
	// A single character that specifies the delimiter used to create CSV job output.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

