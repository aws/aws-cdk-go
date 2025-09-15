package awslogs


// This object defines one log field that will be split with the splitString processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splitStringEntryProperty := &SplitStringEntryProperty{
//   	Delimiter: awscdk.Aws_logs.DelimiterCharacter_COMMA,
//   	Source: jsii.String("source"),
//   }
//
type SplitStringEntryProperty struct {
	// The separator character to split the string on.
	Delimiter DelimiterCharacter `field:"required" json:"delimiter" yaml:"delimiter"`
	// The key of the field to split.
	Source *string `field:"required" json:"source" yaml:"source"`
}

