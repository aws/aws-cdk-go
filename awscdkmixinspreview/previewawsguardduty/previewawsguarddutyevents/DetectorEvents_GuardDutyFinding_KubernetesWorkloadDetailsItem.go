package previewawsguarddutyevents


// Type definition for KubernetesWorkloadDetailsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesWorkloadDetailsItem := &KubernetesWorkloadDetailsItem{
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	ImagePrefix: []*string{
//   		jsii.String("imagePrefix"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	SecurityContext: &SecurityContext{
//   		Privileged: []*string{
//   			jsii.String("privileged"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_KubernetesWorkloadDetailsItem struct {
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// imagePrefix property.
	//
	// Specify an array of string values to match this event if the actual value of imagePrefix is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImagePrefix *[]*string `field:"optional" json:"imagePrefix" yaml:"imagePrefix"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// securityContext property.
	//
	// Specify an array of string values to match this event if the actual value of securityContext is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecurityContext *DetectorEvents_GuardDutyFinding_SecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
}

