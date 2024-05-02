package awsgamelift


// Specifies how much memory is available to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryLimitsProperty := &MemoryLimitsProperty{
//   	HardLimit: jsii.Number(123),
//   	SoftLimit: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-memorylimits.html
//
type CfnContainerGroupDefinition_MemoryLimitsProperty struct {
	// The hard limit of memory to reserve for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-memorylimits.html#cfn-gamelift-containergroupdefinition-memorylimits-hardlimit
	//
	HardLimit *float64 `field:"optional" json:"hardLimit" yaml:"hardLimit"`
	// The amount of memory that is reserved for the container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containergroupdefinition-memorylimits.html#cfn-gamelift-containergroupdefinition-memorylimits-softlimit
	//
	SoftLimit *float64 `field:"optional" json:"softLimit" yaml:"softLimit"`
}

