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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-relatedfinding.html
//
type CfnAutomationRule_RelatedFindingProperty struct {
	// The product-generated identifier for a related finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-relatedfinding.html#cfn-securityhub-automationrule-relatedfinding-id
	//
	Id interface{} `field:"required" json:"id" yaml:"id"`
	// The Amazon Resource Name (ARN) for the product that generated a related finding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-relatedfinding.html#cfn-securityhub-automationrule-relatedfinding-productarn
	//
	ProductArn *string `field:"required" json:"productArn" yaml:"productArn"`
}

