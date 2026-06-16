package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sloSourceProperty := &SloSourceProperty{
//   	PolicyName: jsii.String("policyName"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-slosource.html
//
type CfnService_SloSourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-slosource.html#cfn-resiliencehubv2-service-slosource-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-slosource.html#cfn-resiliencehubv2-service-slosource-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

