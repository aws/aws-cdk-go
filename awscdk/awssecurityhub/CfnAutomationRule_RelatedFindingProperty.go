package awssecurityhub


// Provides details about a list of findings that the current finding relates to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var id interface{}
//
//   relatedFindingProperty := &RelatedFindingProperty{
//   	Id: id,
//   	ProductArn: jsii.String("productArn"),
//   }
//
type CfnAutomationRule_RelatedFindingProperty struct {
	// The product-generated identifier for a related finding.
	Id interface{} `field:"required" json:"id" yaml:"id"`
	// The Amazon Resource Name (ARN) for the product that generated a related finding.
	ProductArn *string `field:"required" json:"productArn" yaml:"productArn"`
}

