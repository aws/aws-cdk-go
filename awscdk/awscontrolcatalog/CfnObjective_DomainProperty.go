package awscontrolcatalog


// The domain that the objective belongs to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainProperty := &DomainProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-objective-domain.html
//
type CfnObjective_DomainProperty struct {
	// The Amazon Resource Name (ARN) of the related domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-objective-domain.html#cfn-controlcatalog-objective-domain-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of the related domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-objective-domain.html#cfn-controlcatalog-objective-domain-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

