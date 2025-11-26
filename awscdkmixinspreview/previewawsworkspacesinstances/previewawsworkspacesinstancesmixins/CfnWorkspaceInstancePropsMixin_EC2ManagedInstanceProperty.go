package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ManagedInstanceProperty := &EC2ManagedInstanceProperty{
//   	InstanceId: jsii.String("instanceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-ec2managedinstance.html
//
type CfnWorkspaceInstancePropsMixin_EC2ManagedInstanceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-ec2managedinstance.html#cfn-workspacesinstances-workspaceinstance-ec2managedinstance-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

