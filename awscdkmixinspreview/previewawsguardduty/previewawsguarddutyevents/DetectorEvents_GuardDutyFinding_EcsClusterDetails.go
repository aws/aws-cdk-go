package previewawsguarddutyevents


// Type definition for EcsClusterDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ecsClusterDetails := &EcsClusterDetails{
//   	Arn: []*string{
//   		jsii.String("arn"),
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
//   	TaskDetails: &TaskDetails{
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		Containers: []TaskDetailsItem{
//   			&TaskDetailsItem{
//   				Image: []*string{
//   					jsii.String("image"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   			},
//   		},
//   		CreatedAt: []*string{
//   			jsii.String("createdAt"),
//   		},
//   		DefinitionArn: []*string{
//   			jsii.String("definitionArn"),
//   		},
//   		StartedAt: []*string{
//   			jsii.String("startedAt"),
//   		},
//   		StartedBy: []*string{
//   			jsii.String("startedBy"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_EcsClusterDetails struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
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
	// taskDetails property.
	//
	// Specify an array of string values to match this event if the actual value of taskDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskDetails *DetectorEvents_GuardDutyFinding_TaskDetails `field:"optional" json:"taskDetails" yaml:"taskDetails"`
}

