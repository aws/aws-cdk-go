package previewawscodebuildevents


// Type definition for Environment_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   environment1 := &Environment1{
//   	ComputeType: []*string{
//   		jsii.String("computeType"),
//   	},
//   	EnvironmentVariables: []EnvironmentItem{
//   		&EnvironmentItem{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Type: []*string{
//   				jsii.String("type"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	Image: []*string{
//   		jsii.String("image"),
//   	},
//   	ImagePullCredentialsType: []*string{
//   		jsii.String("imagePullCredentialsType"),
//   	},
//   	PrivilegedMode: []*string{
//   		jsii.String("privilegedMode"),
//   	},
//   	Type: []*string{
//   		jsii.String("type"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Environment1 struct {
	// computeType property.
	//
	// Specify an array of string values to match this event if the actual value of computeType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ComputeType *[]*string `field:"optional" json:"computeType" yaml:"computeType"`
	// environmentVariables property.
	//
	// Specify an array of string values to match this event if the actual value of environmentVariables is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EnvironmentVariables *[]*AWSAPICallViaCloudTrail_EnvironmentItem `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// image property.
	//
	// Specify an array of string values to match this event if the actual value of image is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Image *[]*string `field:"optional" json:"image" yaml:"image"`
	// imagePullCredentialsType property.
	//
	// Specify an array of string values to match this event if the actual value of imagePullCredentialsType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImagePullCredentialsType *[]*string `field:"optional" json:"imagePullCredentialsType" yaml:"imagePullCredentialsType"`
	// privilegedMode property.
	//
	// Specify an array of string values to match this event if the actual value of privilegedMode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PrivilegedMode *[]*string `field:"optional" json:"privilegedMode" yaml:"privilegedMode"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *[]*string `field:"optional" json:"type" yaml:"type"`
}

