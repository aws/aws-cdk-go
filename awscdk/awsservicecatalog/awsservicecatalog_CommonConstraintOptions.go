package awsservicecatalog


// Properties for governance mechanisms and constraints.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//   var portfolio portfolio
//   var product cloudFormationProduct
//
//
//   topic1 := sns.NewTopic(this, jsii.String("Topic1"))
//   portfolio.notifyOnStackEvents(product, topic1)
//
//   topic2 := sns.NewTopic(this, jsii.String("Topic2"))
//   portfolio.notifyOnStackEvents(product, topic2, &commonConstraintOptions{
//   	description: jsii.String("description for topic2"),
//   })
//
type CommonConstraintOptions struct {
	// The description of the constraint.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
}

