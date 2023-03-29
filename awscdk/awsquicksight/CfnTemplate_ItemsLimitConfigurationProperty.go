package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   itemsLimitConfigurationProperty := &ItemsLimitConfigurationProperty{
//   	ItemsLimit: jsii.Number(123),
//   	OtherCategories: jsii.String("otherCategories"),
//   }
//
type CfnTemplate_ItemsLimitConfigurationProperty struct {
	// `CfnTemplate.ItemsLimitConfigurationProperty.ItemsLimit`.
	ItemsLimit *float64 `field:"optional" json:"itemsLimit" yaml:"itemsLimit"`
	// `CfnTemplate.ItemsLimitConfigurationProperty.OtherCategories`.
	OtherCategories *string `field:"optional" json:"otherCategories" yaml:"otherCategories"`
}

