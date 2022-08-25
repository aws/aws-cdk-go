package awsservicecatalog


// Properties for ResourceUpdateConstraint.
//
// Example:
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//
//   // to disable tag updates:
//   portfolio.constrainTagUpdates(product, &tagUpdateConstraintOptions{
//   	allow: jsii.Boolean(false),
//   })
//
type TagUpdateConstraintOptions struct {
	// The description of the constraint.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// Toggle for if users should be allowed to change/update tags on provisioned products.
	Allow *bool `field:"optional" json:"allow" yaml:"allow"`
}

