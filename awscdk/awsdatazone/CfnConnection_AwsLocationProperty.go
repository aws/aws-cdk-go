package awsdatazone


// AWS Location of project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsLocationProperty := &AwsLocationProperty{
//   	AccessRole: jsii.String("accessRole"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	AwsRegion: jsii.String("awsRegion"),
//   	IamConnectionId: jsii.String("iamConnectionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-awslocation.html
//
type CfnConnection_AwsLocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-awslocation.html#cfn-datazone-connection-awslocation-accessrole
	//
	AccessRole *string `field:"optional" json:"accessRole" yaml:"accessRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-awslocation.html#cfn-datazone-connection-awslocation-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-awslocation.html#cfn-datazone-connection-awslocation-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-awslocation.html#cfn-datazone-connection-awslocation-iamconnectionid
	//
	IamConnectionId *string `field:"optional" json:"iamConnectionId" yaml:"iamConnectionId"`
}

