package awsservicecatalog


// Properties for ResourceUpdateConstraint.
//
// Example:
//   var portfolio Portfolio
//   var product CloudFormationProduct
//
//
//   // to disable tag updates:
//   portfolio.constrainTagUpdates(product, &TagUpdateConstraintOptions{
//   	Allow: jsii.Boolean(false),
//   })
//
type TagUpdateConstraintOptions struct {
	// The description of the constraint.
	// Default: - No description provided.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Default: - English.
	//
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// Toggle for if users should be allowed to change/update tags on provisioned products.
	// Default: true.
	//
	Allow *bool `field:"optional" json:"allow" yaml:"allow"`
}

