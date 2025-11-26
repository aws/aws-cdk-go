package previewawsec2events


// Type definition for SpotOptions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spotOptions := &SpotOptions{
//   	AllocationStrategy: []*string{
//   		jsii.String("allocationStrategy"),
//   	},
//   	InstancePoolConstraintFilterDisabled: []*string{
//   		jsii.String("instancePoolConstraintFilterDisabled"),
//   	},
//   	InstancePoolsToUseCount: []*string{
//   		jsii.String("instancePoolsToUseCount"),
//   	},
//   	MaxInstanceCount: []*string{
//   		jsii.String("maxInstanceCount"),
//   	},
//   	MaxTargetCapacity: []*string{
//   		jsii.String("maxTargetCapacity"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_SpotOptions struct {
	// AllocationStrategy property.
	//
	// Specify an array of string values to match this event if the actual value of AllocationStrategy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AllocationStrategy *[]*string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// InstancePoolConstraintFilterDisabled property.
	//
	// Specify an array of string values to match this event if the actual value of InstancePoolConstraintFilterDisabled is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstancePoolConstraintFilterDisabled *[]*string `field:"optional" json:"instancePoolConstraintFilterDisabled" yaml:"instancePoolConstraintFilterDisabled"`
	// InstancePoolsToUseCount property.
	//
	// Specify an array of string values to match this event if the actual value of InstancePoolsToUseCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstancePoolsToUseCount *[]*string `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// MaxInstanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of MaxInstanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxInstanceCount *[]*string `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// MaxTargetCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of MaxTargetCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxTargetCapacity *[]*string `field:"optional" json:"maxTargetCapacity" yaml:"maxTargetCapacity"`
}

