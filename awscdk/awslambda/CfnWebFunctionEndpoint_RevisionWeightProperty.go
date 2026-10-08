package awslambda


// A revision routing entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   revisionWeightProperty := &RevisionWeightProperty{
//   	RevisionId: jsii.String("revisionId"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-revisionweight.html
//
type CfnWebFunctionEndpoint_RevisionWeightProperty struct {
	// The revision identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-revisionweight.html#cfn-lambda-webfunctionendpoint-revisionweight-revisionid
	//
	RevisionId *string `field:"required" json:"revisionId" yaml:"revisionId"`
	// The traffic weight for this revision.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-webfunctionendpoint-revisionweight.html#cfn-lambda-webfunctionendpoint-revisionweight-weight
	//
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

