package previewawsworkspacesinstancesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cpuOptionsRequestProperty := &CpuOptionsRequestProperty{
//   	CoreCount: jsii.Number(123),
//   	ThreadsPerCore: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-cpuoptionsrequest.html
//
type CfnWorkspaceInstancePropsMixin_CpuOptionsRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-cpuoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-cpuoptionsrequest-corecount
	//
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-cpuoptionsrequest.html#cfn-workspacesinstances-workspaceinstance-cpuoptionsrequest-threadspercore
	//
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

