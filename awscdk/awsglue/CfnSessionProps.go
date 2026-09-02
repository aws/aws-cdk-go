package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSessionProps := &CfnSessionProps{
//   	Command: &SessionCommandProperty{
//   		Name: jsii.String("name"),
//   		PythonVersion: jsii.String("pythonVersion"),
//   	},
//   	Id: jsii.String("id"),
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	Connections: &ConnectionsListProperty{
//   		Connections: []*string{
//   			jsii.String("connections"),
//   		},
//   	},
//   	DefaultArguments: map[string]*string{
//   		"defaultArgumentsKey": jsii.String("defaultArguments"),
//   	},
//   	Description: jsii.String("description"),
//   	GlueVersion: jsii.String("glueVersion"),
//   	IdleTimeout: jsii.Number(123),
//   	MaxCapacity: jsii.Number(123),
//   	NumberOfWorkers: jsii.Number(123),
//   	RequestOrigin: jsii.String("requestOrigin"),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   	WorkerType: jsii.String("workerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html
//
type CfnSessionProps struct {
	// The SessionCommand that runs the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-command
	//
	Command interface{} `field:"required" json:"command" yaml:"command"`
	// The ID of the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The IAM Role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
	// Specifies the connections used by the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-connections
	//
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// A map array of key-value pairs.
	//
	// Max is 75 pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-defaultarguments
	//
	DefaultArguments interface{} `field:"optional" json:"defaultArguments" yaml:"defaultArguments"`
	// The description of the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Glue version determines the versions of Apache Spark and Python that Glue supports.
	//
	// The GlueVersion must be greater than 2.0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-glueversion
	//
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// The number of minutes when idle before session times out.
	//
	// Default is the value of Timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-idletimeout
	//
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// The number of Glue data processing units (DPUs) that can be allocated when the job runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The number of workers of a defined WorkerType to use for the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-numberofworkers
	//
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The origin of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-requestorigin
	//
	RequestOrigin *string `field:"optional" json:"requestOrigin" yaml:"requestOrigin"`
	// The name of the SecurityConfiguration structure to be used with the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-securityconfiguration
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The tags belonging to the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The number of minutes before session times out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The type of predefined worker that is allocated when a session runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-session.html#cfn-glue-session-workertype
	//
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

