package awsbatch


// Determine whether your data volume persists on the host container instance and where it's stored.
//
// If this parameter is empty, then the Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to persist after the containers that are associated with it stop running.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumesHostProperty := &VolumesHostProperty{
//   	SourcePath: jsii.String("sourcePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html
//
type CfnJobDefinition_VolumesHostProperty struct {
	// The path on the host container instance that's presented to the container.
	//
	// If this parameter is empty, then the Docker daemon has assigned a host path for you. If this parameter contains a file location, then the data volume persists at the specified location on the host container instance until you delete it manually. If the source path location doesn't exist on the host container instance, the Docker daemon creates it. If the location does exist, the contents of the source path folder are exported.
	//
	// > This parameter isn't applicable to jobs that run on Fargate resources. Don't provide this for these jobs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html#cfn-batch-jobdefinition-volumeshost-sourcepath
	//
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

