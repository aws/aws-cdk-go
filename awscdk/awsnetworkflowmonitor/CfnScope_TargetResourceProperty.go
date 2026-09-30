package awsnetworkflowmonitor


// A target resource in a scope, identified by a Region and target identifier.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetResourceProperty := &TargetResourceProperty{
//   	Region: jsii.String("region"),
//   	TargetIdentifier: &TargetIdentifierProperty{
//   		TargetId: &TargetIdProperty{
//   			AccountId: jsii.String("accountId"),
//   		},
//   		TargetType: jsii.String("targetType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetresource.html
//
type CfnScope_TargetResourceProperty struct {
	// The AWS Region for the target resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetresource.html#cfn-networkflowmonitor-scope-targetresource-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// A target identifier is a pair of identifying information for a scope target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkflowmonitor-scope-targetresource.html#cfn-networkflowmonitor-scope-targetresource-targetidentifier
	//
	TargetIdentifier interface{} `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

