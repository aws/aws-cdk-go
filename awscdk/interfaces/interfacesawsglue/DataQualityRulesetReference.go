package interfacesawsglue


// A reference to a DataQualityRuleset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityRulesetReference := &DataQualityRulesetReference{
//   	DataQualityRulesetName: jsii.String("dataQualityRulesetName"),
//   }
//
type DataQualityRulesetReference struct {
	// The Name of the DataQualityRuleset resource.
	DataQualityRulesetName *string `field:"required" json:"dataQualityRulesetName" yaml:"dataQualityRulesetName"`
}

