package awsssm


// The `Targets` property type specifies adding a target to a maintenance window target in AWS Systems Manager .
//
// `Targets` is a property of the [AWS::SSM::MaintenanceWindowTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetsProperty := &targetsProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnMaintenanceWindowTarget_TargetsProperty struct {
	// User-defined criteria for sending commands that target managed nodes that meet the criteria.
	Key *string `field:"required" json:"key" yaml:"key"`
	// User-defined criteria that maps to `Key` .
	//
	// For example, if you specified `tag:ServerRole` , you could specify `value:WebServer` to run a command on instances that include EC2 tags of `ServerRole,WebServer` .
	//
	// Depending on the type of target, the maximum number of values for a key might be lower than the global maximum of 50.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

