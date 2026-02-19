package awsconnect


// After Contact Work configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   afterContactWorkConfigProperty := &AfterContactWorkConfigProperty{
//   	AfterContactWorkTimeLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfig.html
//
type CfnUser_AfterContactWorkConfigProperty struct {
	// The After Call Work (ACW) timeout setting, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfig.html#cfn-connect-user-aftercontactworkconfig-aftercontactworktimelimit
	//
	AfterContactWorkTimeLimit *float64 `field:"optional" json:"afterContactWorkTimeLimit" yaml:"afterContactWorkTimeLimit"`
}

