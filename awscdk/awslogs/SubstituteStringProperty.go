package awslogs


// This processor matches a key's value against a regular expression and replaces all matches with a replacement string.
//
// For more information about this processor including examples, see substituteString in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   substituteStringProperty := &SubstituteStringProperty{
//   	Entries: []SubstituteStringEntryProperty{
//   		&SubstituteStringEntryProperty{
//   			From: jsii.String("from"),
//   			Source: jsii.String("source"),
//   			To: jsii.String("to"),
//   		},
//   	},
//   }
//
type SubstituteStringProperty struct {
	// An array of objects, where each object contains information about one key to match and replace.
	Entries *[]*SubstituteStringEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

