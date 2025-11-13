package interfacesawsglue


// A reference to a DataQualityRuleset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataQualityRulesetReference := &DataQualityRulesetReference{
//   	DataQualityRulesetId: jsii.String("dataQualityRulesetId"),
//   }
//
type DataQualityRulesetReference struct {
	// The Id of the DataQualityRuleset resource.
	DataQualityRulesetId *string `field:"required" json:"dataQualityRulesetId" yaml:"dataQualityRulesetId"`
}

