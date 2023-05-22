// An experiment to bundle the entire CDK into a single module
package awscdk


// Construction properties for {@link Resource}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceProps := &ResourceProps{
//   	Account: jsii.String("account"),
//   	EnvironmentFromArn: jsii.String("environmentFromArn"),
//   	PhysicalName: jsii.String("physicalName"),
//   	Region: jsii.String("region"),
//   }
//
// Experimental.
type ResourceProps struct {
	// The AWS account ID this resource belongs to.
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// ARN to deduce region and account from.
	//
	// The ARN is parsed and the account and region are taken from the ARN.
	// This should be used for imported resources.
	//
	// Cannot be supplied together with either `account` or `region`.
	// Experimental.
	EnvironmentFromArn *string `field:"optional" json:"environmentFromArn" yaml:"environmentFromArn"`
	// The value passed in by users to the physical name prop of the resource.
	//
	// - `undefined` implies that a physical name will be allocated by
	//    CloudFormation during deployment.
	// - a concrete value implies a specific physical name
	// - `PhysicalName.GENERATE_IF_NEEDED` is a marker that indicates that a physical will only be generated
	//    by the CDK if it is needed for cross-environment references. Otherwise, it will be allocated by CloudFormation.
	// Experimental.
	PhysicalName *string `field:"optional" json:"physicalName" yaml:"physicalName"`
	// The AWS region this resource belongs to.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

