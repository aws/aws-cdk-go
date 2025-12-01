package awscloudfront


// A connection function association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionFunctionAssociationProperty := &ConnectionFunctionAssociationProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-connectionfunctionassociation.html
//
type CfnDistribution_ConnectionFunctionAssociationProperty struct {
	// The association's ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-connectionfunctionassociation.html#cfn-cloudfront-distribution-connectionfunctionassociation-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
}

