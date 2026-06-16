package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   disasterRecoverySourceProperty := &DisasterRecoverySourceProperty{
//   	PolicyName: jsii.String("policyName"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-disasterrecoverysource.html
//
type CfnService_DisasterRecoverySourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-disasterrecoverysource.html#cfn-resiliencehubv2-service-disasterrecoverysource-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-disasterrecoverysource.html#cfn-resiliencehubv2-service-disasterrecoverysource-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

