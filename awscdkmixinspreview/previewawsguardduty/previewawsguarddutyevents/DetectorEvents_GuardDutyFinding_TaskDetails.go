package previewawsguarddutyevents


// Type definition for TaskDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskDetails := &TaskDetails{
//   	Arn: []*string{
//   		jsii.String("arn"),
//   	},
//   	Containers: []TaskDetailsItem{
//   		&TaskDetailsItem{
//   			Image: []*string{
//   				jsii.String("image"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	DefinitionArn: []*string{
//   		jsii.String("definitionArn"),
//   	},
//   	StartedAt: []*string{
//   		jsii.String("startedAt"),
//   	},
//   	StartedBy: []*string{
//   		jsii.String("startedBy"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_TaskDetails struct {
	// arn property.
	//
	// Specify an array of string values to match this event if the actual value of arn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Arn *[]*string `field:"optional" json:"arn" yaml:"arn"`
	// containers property.
	//
	// Specify an array of string values to match this event if the actual value of containers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Containers *[]*DetectorEvents_GuardDutyFinding_TaskDetailsItem `field:"optional" json:"containers" yaml:"containers"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// definitionArn property.
	//
	// Specify an array of string values to match this event if the actual value of definitionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DefinitionArn *[]*string `field:"optional" json:"definitionArn" yaml:"definitionArn"`
	// startedAt property.
	//
	// Specify an array of string values to match this event if the actual value of startedAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedAt *[]*string `field:"optional" json:"startedAt" yaml:"startedAt"`
	// startedBy property.
	//
	// Specify an array of string values to match this event if the actual value of startedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartedBy *[]*string `field:"optional" json:"startedBy" yaml:"startedBy"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

