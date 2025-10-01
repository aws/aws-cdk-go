package awsarcregionswitch


// Configuration for AWS Lambda functions used in a Region switch plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdasProperty := &LambdasProperty{
//   	Arn: jsii.String("arn"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdas.html
//
type CfnPlan_LambdasProperty struct {
	// The Amazon Resource Name (ARN) of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdas.html#cfn-arcregionswitch-plan-lambdas-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The cross account role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdas.html#cfn-arcregionswitch-plan-lambdas-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The external ID (secret key) for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdas.html#cfn-arcregionswitch-plan-lambdas-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

