package awsgreengrass


// The access identity whose permissions are used to run the Lambda function.
//
// This setting overrides the default access identity that's specified for the group (by default, ggc_user and ggc_group). You can override the user, group, or both. For more information, see [Run as](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-access-identity.html) in the *Developer Guide* .
//
// > Running as the root user increases risks to your data and device. Do not run as root (UID/GID=0) unless your business case requires it. For more information and requirements, see [Running a Lambda Function as Root](https://docs.aws.amazon.com/greengrass/latest/developerguide/lambda-group-config.html#lambda-running-as-root) .
//
// In an AWS CloudFormation template, `RunAs` is a property of the [`Execution`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-execution.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runAsProperty := &RunAsProperty{
//   	Gid: jsii.Number(123),
//   	Uid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-runas.html
//
type CfnFunctionDefinition_RunAsProperty struct {
	// The group ID whose permissions are used to run the Lambda function.
	//
	// You can use the `getent group` command on your core device to look up the group ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-runas.html#cfn-greengrass-functiondefinition-runas-gid
	//
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The user ID whose permissions are used to run the Lambda function.
	//
	// You can use the `getent passwd` command on your core device to look up the user ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinition-runas.html#cfn-greengrass-functiondefinition-runas-uid
	//
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

