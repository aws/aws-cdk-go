package awslex


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dTMFSpecificationProperty := &dTMFSpecificationProperty{
//   	deletionCharacter: jsii.String("deletionCharacter"),
//   	endCharacter: jsii.String("endCharacter"),
//   	endTimeoutMs: jsii.Number(123),
//   	maxLength: jsii.Number(123),
//   }
//
type CfnBot_DTMFSpecificationProperty struct {
	// `CfnBot.DTMFSpecificationProperty.DeletionCharacter`.
	DeletionCharacter *string `field:"required" json:"deletionCharacter" yaml:"deletionCharacter"`
	// `CfnBot.DTMFSpecificationProperty.EndCharacter`.
	EndCharacter *string `field:"required" json:"endCharacter" yaml:"endCharacter"`
	// `CfnBot.DTMFSpecificationProperty.EndTimeoutMs`.
	EndTimeoutMs *float64 `field:"required" json:"endTimeoutMs" yaml:"endTimeoutMs"`
	// `CfnBot.DTMFSpecificationProperty.MaxLength`.
	MaxLength *float64 `field:"required" json:"maxLength" yaml:"maxLength"`
}

