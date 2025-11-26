package previewawsguarddutyevents


// Type definition for KubernetesDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesDetails := &KubernetesDetails{
//   	KubernetesUserDetails: &KubernetesUserDetails{
//   		Groups: []*string{
//   			jsii.String("groups"),
//   		},
//   		Uid: []*string{
//   			jsii.String("uid"),
//   		},
//   		Username: []*string{
//   			jsii.String("username"),
//   		},
//   	},
//   	KubernetesWorkloadDetails: &KubernetesWorkloadDetails{
//   		Containers: []KubernetesWorkloadDetailsItem{
//   			&KubernetesWorkloadDetailsItem{
//   				Image: []*string{
//   					jsii.String("image"),
//   				},
//   				ImagePrefix: []*string{
//   					jsii.String("imagePrefix"),
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				SecurityContext: &SecurityContext{
//   					Privileged: []*string{
//   						jsii.String("privileged"),
//   					},
//   				},
//   			},
//   		},
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		Namespace: []*string{
//   			jsii.String("namespace"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   		Uid: []*string{
//   			jsii.String("uid"),
//   		},
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_KubernetesDetails struct {
	// kubernetesUserDetails property.
	//
	// Specify an array of string values to match this event if the actual value of kubernetesUserDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KubernetesUserDetails *DetectorEvents_GuardDutyFinding_KubernetesUserDetails `field:"optional" json:"kubernetesUserDetails" yaml:"kubernetesUserDetails"`
	// kubernetesWorkloadDetails property.
	//
	// Specify an array of string values to match this event if the actual value of kubernetesWorkloadDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	KubernetesWorkloadDetails *DetectorEvents_GuardDutyFinding_KubernetesWorkloadDetails `field:"optional" json:"kubernetesWorkloadDetails" yaml:"kubernetesWorkloadDetails"`
}

