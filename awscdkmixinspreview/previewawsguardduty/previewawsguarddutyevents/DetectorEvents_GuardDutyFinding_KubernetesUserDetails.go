package previewawsguarddutyevents


// Type definition for KubernetesUserDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kubernetesUserDetails := &KubernetesUserDetails{
//   	Groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	Uid: []*string{
//   		jsii.String("uid"),
//   	},
//   	Username: []*string{
//   		jsii.String("username"),
//   	},
//   }
//
// Experimental.
type DetectorEvents_GuardDutyFinding_KubernetesUserDetails struct {
	// groups property.
	//
	// Specify an array of string values to match this event if the actual value of groups is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// uid property.
	//
	// Specify an array of string values to match this event if the actual value of uid is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Uid *[]*string `field:"optional" json:"uid" yaml:"uid"`
	// username property.
	//
	// Specify an array of string values to match this event if the actual value of username is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Username *[]*string `field:"optional" json:"username" yaml:"username"`
}

