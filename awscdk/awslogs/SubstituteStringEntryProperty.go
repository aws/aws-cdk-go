package awslogs


// This object defines one log field key that will be replaced using the substituteString processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   substituteStringEntryProperty := &SubstituteStringEntryProperty{
//   	From: jsii.String("from"),
//   	Source: jsii.String("source"),
//   	To: jsii.String("to"),
//   }
//
type SubstituteStringEntryProperty struct {
	// The regular expression string to be replaced.
	From *string `field:"required" json:"from" yaml:"from"`
	// The key to modify.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The string to be substituted for each match of from.
	To *string `field:"required" json:"to" yaml:"to"`
}

