package previewawsguarddutyevents


// Type definition for KubernetesWorkloadDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesWorkloadDetails := &KubernetesWorkloadDetails{
//   	Containers: []KubernetesWorkloadDetailsItem{
//   		&KubernetesWorkloadDetailsItem{
//   			Image: []*string{
//   				jsii.String("image"),
//   			},
//   			ImagePrefix: []*string{
//   				jsii.String("imagePrefix"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			SecurityContext: &SecurityContext{
//   				Privileged: []*string{
//   					jsii.String("privileged"),
//   				},
//   			},
//   		},
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Namespace: []*string{
//   		jsii.String("namespace"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   	Uid: []*string{
//   		jsii.String("uid"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_KubernetesWorkloadDetails struct {
	// containers property.
	//
	// Specify an array of string values to match this event if the actual value of containers is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Containers *[]*DetectorEvents_GuardDutyFinding_KubernetesWorkloadDetailsItem `field:"optional" json:"containers" yaml:"containers"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// namespace property.
	//
	// Specify an array of string values to match this event if the actual value of namespace is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Namespace *[]*string `field:"optional" json:"namespace" yaml:"namespace"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
	// uid property.
	//
	// Specify an array of string values to match this event if the actual value of uid is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Uid *[]*string `field:"optional" json:"uid" yaml:"uid"`
}

