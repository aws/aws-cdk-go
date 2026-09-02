package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createdByInfoProperty := &CreatedByInfoProperty{
//   	AwsIdentityArn: jsii.String("awsIdentityArn"),
//   	ConnectUserArn: jsii.String("connectUserArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-createdbyinfo.html
//
type CfnMetric_CreatedByInfoProperty struct {
	// STS or IAM ARN representing the identity of API Caller.
	//
	// SDK users cannot populate this and this value is calculated automatically if ConnectUserArn is not provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-createdbyinfo.html#cfn-connect-metric-createdbyinfo-awsidentityarn
	//
	AwsIdentityArn *string `field:"optional" json:"awsIdentityArn" yaml:"awsIdentityArn"`
	// An agent ARN representing a connect user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-metric-createdbyinfo.html#cfn-connect-metric-createdbyinfo-connectuserarn
	//
	ConnectUserArn *string `field:"optional" json:"connectUserArn" yaml:"connectUserArn"`
}

