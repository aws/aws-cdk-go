package awsfis


// A reference to a TargetAccountConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetAccountConfigurationReference := &TargetAccountConfigurationReference{
//   	AccountId: jsii.String("accountId"),
//   	ExperimentTemplateId: jsii.String("experimentTemplateId"),
//   }
//
type TargetAccountConfigurationReference struct {
	// The AccountId of the TargetAccountConfiguration resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ExperimentTemplateId of the TargetAccountConfiguration resource.
	ExperimentTemplateId *string `field:"required" json:"experimentTemplateId" yaml:"experimentTemplateId"`
}

