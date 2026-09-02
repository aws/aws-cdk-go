package interfacesawsaccessanalyzer


// A reference to a ArchiveRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveRuleReference := &ArchiveRuleReference{
//   	ArchiveRuleArn: jsii.String("archiveRuleArn"),
//   }
//
type ArchiveRuleReference struct {
	// The Arn of the ArchiveRule resource.
	ArchiveRuleArn *string `field:"required" json:"archiveRuleArn" yaml:"archiveRuleArn"`
}

