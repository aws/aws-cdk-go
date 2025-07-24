package awsworkspacesinstances


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   creditSpecificationRequestProperty := &CreditSpecificationRequestProperty{
//   	CpuCredits: jsii.String("cpuCredits"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-creditspecificationrequest.html
//
type CfnWorkspaceInstance_CreditSpecificationRequestProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesinstances-workspaceinstance-creditspecificationrequest.html#cfn-workspacesinstances-workspaceinstance-creditspecificationrequest-cpucredits
	//
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

