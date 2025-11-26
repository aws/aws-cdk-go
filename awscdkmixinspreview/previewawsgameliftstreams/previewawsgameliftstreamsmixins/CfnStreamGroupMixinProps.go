package previewawsgameliftstreamsmixins


// Properties for CfnStreamGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStreamGroupMixinProps := &CfnStreamGroupMixinProps{
//   	DefaultApplication: &DefaultApplicationProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Description: jsii.String("description"),
//   	LocationConfigurations: []interface{}{
//   		&LocationConfigurationProperty{
//   			AlwaysOnCapacity: jsii.Number(123),
//   			LocationName: jsii.String("locationName"),
//   			OnDemandCapacity: jsii.Number(123),
//   		},
//   	},
//   	StreamClass: jsii.String("streamClass"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html
//
type CfnStreamGroupMixinProps struct {
	// Object that identifies the Amazon GameLift Streams application to stream with this stream group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-defaultapplication
	//
	DefaultApplication interface{} `field:"optional" json:"defaultApplication" yaml:"defaultApplication"`
	// A descriptive label for the stream group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A set of one or more locations and the streaming capacity for each location.
	//
	// One of the locations MUST be your primary location, which is the AWS Region where you are specifying this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-locationconfigurations
	//
	LocationConfigurations interface{} `field:"optional" json:"locationConfigurations" yaml:"locationConfigurations"`
	// The target stream quality for sessions that are hosted in this stream group.
	//
	// Set a stream class that is appropriate to the type of content that you're streaming. Stream class determines the type of computing resources Amazon GameLift Streams uses and impacts the cost of streaming. The following options are available:
	//
	// A stream class can be one of the following:
	//
	// - *`gen5n_win2022` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.4, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA A10G Tensor GPU.
	//
	// - Reference resolution: 1080p
	// - Reference frame rate: 60 fps
	// - Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM
	// - Tenancy: Supports 1 concurrent stream session
	// - *`gen5n_high` (NVIDIA, high)* Supports applications with moderate to high 3D scene complexity. Uses NVIDIA A10G Tensor GPU.
	//
	// - Reference resolution: 1080p
	// - Reference frame rate: 60 fps
	// - Workload specifications: 4 vCPUs, 16 GB RAM, 12 GB VRAM
	// - Tenancy: Supports up to 2 concurrent stream sessions
	// - *`gen5n_ultra` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Uses dedicated NVIDIA A10G Tensor GPU.
	//
	// - Reference resolution: 1080p
	// - Reference frame rate: 60 fps
	// - Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM
	// - Tenancy: Supports 1 concurrent stream session
	// - *`gen4n_win2022` (NVIDIA, ultra)* Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.4, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA T4 Tensor GPU.
	//
	// - Reference resolution: 1080p
	// - Reference frame rate: 60 fps
	// - Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM
	// - Tenancy: Supports 1 concurrent stream session
	// - *`gen4n_high` (NVIDIA, high)* Supports applications with moderate to high 3D scene complexity. Uses NVIDIA T4 Tensor GPU.
	//
	// - Reference resolution: 1080p
	// - Reference frame rate: 60 fps
	// - Workload specifications: 4 vCPUs, 16 GB RAM, 8 GB VRAM
	// - Tenancy: Supports up to 2 concurrent stream sessions
	// - *`gen4n_ultra` (NVIDIA, ultra)* Supports applications with high 3D scene complexity. Uses dedicated NVIDIA T4 Tensor GPU.
	//
	// - Reference resolution: 1080p
	// - Reference frame rate: 60 fps
	// - Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM
	// - Tenancy: Supports 1 concurrent stream session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-streamclass
	//
	StreamClass *string `field:"optional" json:"streamClass" yaml:"streamClass"`
	// A list of labels to assign to the new stream group resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gameliftstreams-streamgroup.html#cfn-gameliftstreams-streamgroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

