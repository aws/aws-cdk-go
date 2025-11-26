package previewawsroute53recoveryreadinessevents


// Type definition for State.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   state := &State{
//   	ReadinessStatus: []*string{
//   		jsii.String("readinessStatus"),
//   	},
//   }
//
// Experimental.
type ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_State struct {
	// readiness-status property.
	//
	// Specify an array of string values to match this event if the actual value of readiness-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ReadinessStatus *[]*string `field:"optional" json:"readinessStatus" yaml:"readinessStatus"`
}

