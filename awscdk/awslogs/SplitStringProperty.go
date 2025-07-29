package awslogs


// Use this processor to split a field into an array of strings using a delimiting character.
//
// For more information about this processor including examples, see splitString in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   splitStringProperty := &SplitStringProperty{
//   	Entries: []splitStringEntryProperty{
//   		&splitStringEntryProperty{
//   			Delimiter: awscdk.Aws_logs.DelimiterCharacter_COMMA,
//   			Source: jsii.String("source"),
//   		},
//   	},
//   }
//
type SplitStringProperty struct {
	// An array of SplitStringEntry objects, where each object contains information about one field to split.
	Entries *[]*SplitStringEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

