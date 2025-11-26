package previewawsguarddutyevents


// Type definition for EksClusterDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eksClusterDetails := &EksClusterDetails{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	Tags: []EcsClusterDetailsItem{
//   		&EcsClusterDetailsItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	VpcId: []*string{
//   		jsii.String("vpcId"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_EksClusterDetails struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*DetectorEvents_GuardDutyFinding_EcsClusterDetailsItem `field:"optional" json:"tags" yaml:"tags"`
	// vpcId property.
	//
	// Specify an array of string values to match this event if the actual value of vpcId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VpcId *[]*string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

