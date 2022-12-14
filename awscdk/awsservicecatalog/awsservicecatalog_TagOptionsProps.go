package awsservicecatalog


// Properties for TagOptions.
//
// Example:
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//
//   tagOptionsForPortfolio := servicecatalog.NewTagOptions(this, jsii.String("OrgTagOptions"), &tagOptionsProps{
//   	allowedValuesForTags: map[string][]*string{
//   		"Group": []*string{
//   			jsii.String("finance"),
//   			jsii.String("engineering"),
//   			jsii.String("marketing"),
//   			jsii.String("research"),
//   		},
//   		"CostCenter": []*string{
//   			jsii.String("01"),
//   			jsii.String("02"),
//   			jsii.String("03"),
//   		},
//   	},
//   })
//   portfolio.associateTagOptions(tagOptionsForPortfolio)
//
//   tagOptionsForProduct := servicecatalog.NewTagOptions(this, jsii.String("ProductTagOptions"), &tagOptionsProps{
//   	allowedValuesForTags: map[string][]*string{
//   		"Environment": []*string{
//   			jsii.String("dev"),
//   			jsii.String("alpha"),
//   			jsii.String("prod"),
//   		},
//   	},
//   })
//   product.associateTagOptions(tagOptionsForProduct)
//
type TagOptionsProps struct {
	// The values that are allowed to be set for specific tags.
	//
	// The keys of the map represent the tag keys,
	// and the values of the map are a list of allowed values for that particular tag key.
	AllowedValuesForTags *map[string]*[]*string `field:"required" json:"allowedValuesForTags" yaml:"allowedValuesForTags"`
}

