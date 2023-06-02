package awssimspaceweaver


// Properties for defining a `CfnSimulation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSimulationProps := &CfnSimulationProps{
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	MaximumDuration: jsii.String("maximumDuration"),
//   	SchemaS3Location: &S3LocationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   	},
//   	SnapshotS3Location: &S3LocationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		ObjectKey: jsii.String("objectKey"),
//   	},
//   }
//
type CfnSimulationProps struct {
	// The name of the simulation.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management ( IAM ) role that the simulation assumes to perform actions.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* . For more information about IAM roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the *AWS Identity and Access Management User Guide* .
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The maximum running time of the simulation, specified as a number of minutes (m or M), hours (h or H), or days (d or D).
	//
	// The simulation stops when it reaches this limit. The maximum value is `14D` , or its equivalent in the other units. The default value is `14D` . A value equivalent to `0` makes the simulation immediately transition to `STOPPING` as soon as it reaches `STARTED` .
	MaximumDuration *string `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// The location of the simulation schema in Amazon Simple Storage Service ( Amazon S3 ).
	//
	// For more information about Amazon S3 , see the [*Amazon Simple Storage Service User Guide*](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html) .
	//
	// Provide a `SchemaS3Location` to start your simulation from a schema.
	//
	// If you provide a `SchemaS3Location` then you can't provide a `SnapshotS3Location` .
	SchemaS3Location interface{} `field:"optional" json:"schemaS3Location" yaml:"schemaS3Location"`
	// The location of the snapshot in Amazon Simple Storage Service ( Amazon S3 ).
	//
	// For more information about Amazon S3 , see the [*Amazon Simple Storage Service User Guide*](https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html) .
	//
	// Provide a `SnapshotS3Location` to start your simulation from a snapshot.
	//
	// If you provide a `SnapshotS3Location` then you can't provide a `SchemaS3Location` .
	SnapshotS3Location interface{} `field:"optional" json:"snapshotS3Location" yaml:"snapshotS3Location"`
}

