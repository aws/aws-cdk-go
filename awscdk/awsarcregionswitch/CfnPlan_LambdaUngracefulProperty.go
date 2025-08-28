package awsarcregionswitch


// Configuration for handling failures when invoking Lambda functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaUngracefulProperty := &LambdaUngracefulProperty{
//   	Behavior: jsii.String("behavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaungraceful.html
//
type CfnPlan_LambdaUngracefulProperty struct {
	// The ungraceful behavior for a Lambda function, which must be set to `skip` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaungraceful.html#cfn-arcregionswitch-plan-lambdaungraceful-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
}

