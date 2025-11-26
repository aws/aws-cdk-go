package previewawsec2events


// Type definition for LaunchTemplate_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchTemplate1 := &LaunchTemplate1{
//   	CreatedBy: []*string{
//   		jsii.String("createdBy"),
//   	},
//   	CreateTime: []*string{
//   		jsii.String("createTime"),
//   	},
//   	DefaultVersionNumber: []*string{
//   		jsii.String("defaultVersionNumber"),
//   	},
//   	LatestVersionNumber: []*string{
//   		jsii.String("latestVersionNumber"),
//   	},
//   	LaunchTemplateId: []*string{
//   		jsii.String("launchTemplateId"),
//   	},
//   	LaunchTemplateName: []*string{
//   		jsii.String("launchTemplateName"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_LaunchTemplate1 struct {
	// createdBy property.
	//
	// Specify an array of string values to match this event if the actual value of createdBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedBy *[]*string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// createTime property.
	//
	// Specify an array of string values to match this event if the actual value of createTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateTime *[]*string `field:"optional" json:"createTime" yaml:"createTime"`
	// defaultVersionNumber property.
	//
	// Specify an array of string values to match this event if the actual value of defaultVersionNumber is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DefaultVersionNumber *[]*string `field:"optional" json:"defaultVersionNumber" yaml:"defaultVersionNumber"`
	// latestVersionNumber property.
	//
	// Specify an array of string values to match this event if the actual value of latestVersionNumber is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LatestVersionNumber *[]*string `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// launchTemplateId property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplateId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateId *[]*string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// launchTemplateName property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplateName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplateName *[]*string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

