package awsapplicationsignals


// A reference to a GroupingConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupingConfigurationReference := &GroupingConfigurationReference{
//   	AccountId: jsii.String("accountId"),
//   }
//
type GroupingConfigurationReference struct {
	// The AccountId of the GroupingConfiguration resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
}

