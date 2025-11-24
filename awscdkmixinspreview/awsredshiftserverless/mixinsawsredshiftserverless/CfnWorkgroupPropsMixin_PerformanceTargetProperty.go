package mixinsawsredshiftserverless


// An object that represents the price performance target settings for the workgroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   performanceTargetProperty := &PerformanceTargetProperty{
//   	Level: jsii.Number(123),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-performancetarget.html
//
type CfnWorkgroupPropsMixin_PerformanceTargetProperty struct {
	// The target price performance level for the workgroup.
	//
	// Valid values include 1, 25, 50, 75, and 100. These correspond to the price performance levels LOW_COST, ECONOMICAL, BALANCED, RESOURCEFUL, and HIGH_PERFORMANCE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-performancetarget.html#cfn-redshiftserverless-workgroup-performancetarget-level
	//
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// Whether the price performance target is enabled for the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-performancetarget.html#cfn-redshiftserverless-workgroup-performancetarget-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

