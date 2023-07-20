package awscognito


// A map containing a priority as a key, and recovery method name as a value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recoveryOptionProperty := &RecoveryOptionProperty{
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-recoveryoption.html
//
type CfnUserPool_RecoveryOptionProperty struct {
	// Specifies the recovery method for a user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-recoveryoption.html#cfn-cognito-userpool-recoveryoption-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A positive integer specifying priority of a method with 1 being the highest priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-recoveryoption.html#cfn-cognito-userpool-recoveryoption-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
}

