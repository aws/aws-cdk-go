package awscdk


// Resource violating a specific rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyViolatingResource := &PolicyViolatingResource{
//   	Locations: []*string{
//   		jsii.String("locations"),
//   	},
//
//   	// the properties below are optional
//   	ConstructPath: jsii.String("constructPath"),
//   	ResourceLogicalId: jsii.String("resourceLogicalId"),
//   	TemplatePath: jsii.String("templatePath"),
//   }
//
type PolicyViolatingResource struct {
	// The locations in the CloudFormation template that pose the violations.
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// The construct path of the violating construct.
	//
	// Use this for violations that originate from constructs rather than
	// CloudFormation resources (e.g. annotations added via `Annotations.of()`
	// or `Validations.of()`). When provided, the report will use this path
	// directly instead of deriving it from the resource logical ID.
	// Mutually exclusive with `resourceLogicalId`.
	// Default: - construct path is derived from the resource logical ID.
	//
	ConstructPath *string `field:"optional" json:"constructPath" yaml:"constructPath"`
	// The logical ID of the resource in the CloudFormation template.
	//
	// Required for plugin-sourced violations that operate on CloudFormation
	// templates. Mutually exclusive with `constructPath`.
	// Default: - no resource logical ID.
	//
	ResourceLogicalId *string `field:"optional" json:"resourceLogicalId" yaml:"resourceLogicalId"`
	// The path to the CloudFormation template that contains this resource.
	// Default: - no template path.
	//
	TemplatePath *string `field:"optional" json:"templatePath" yaml:"templatePath"`
}

