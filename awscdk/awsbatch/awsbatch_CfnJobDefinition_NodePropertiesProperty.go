package awsbatch


// An object representing the node properties of a multi-node parallel job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   nodePropertiesProperty := &nodePropertiesProperty{
//   	mainNode: jsii.Number(123),
//   	nodeRangeProperties: []interface{}{
//   		&nodeRangePropertyProperty{
//   			targetNodes: jsii.String("targetNodes"),
//
//   			// the properties below are optional
//   			container: &containerPropertiesProperty{
//   				image: jsii.String("image"),
//
//   				// the properties below are optional
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				environment: []interface{}{
//   					&environmentProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				fargatePlatformConfiguration: &fargatePlatformConfigurationProperty{
//   					platformVersion: jsii.String("platformVersion"),
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				jobRoleArn: jsii.String("jobRoleArn"),
//   				linuxParameters: &linuxParametersProperty{
//   					devices: []interface{}{
//   						&deviceProperty{
//   							containerPath: jsii.String("containerPath"),
//   							hostPath: jsii.String("hostPath"),
//   							permissions: []*string{
//   								jsii.String("permissions"),
//   							},
//   						},
//   					},
//   					initProcessEnabled: jsii.Boolean(false),
//   					maxSwap: jsii.Number(123),
//   					sharedMemorySize: jsii.Number(123),
//   					swappiness: jsii.Number(123),
//   					tmpfs: []interface{}{
//   						&tmpfsProperty{
//   							containerPath: jsii.String("containerPath"),
//   							size: jsii.Number(123),
//
//   							// the properties below are optional
//   							mountOptions: []*string{
//   								jsii.String("mountOptions"),
//   							},
//   						},
//   					},
//   				},
//   				logConfiguration: &logConfigurationProperty{
//   					logDriver: jsii.String("logDriver"),
//
//   					// the properties below are optional
//   					options: options,
//   					secretOptions: []interface{}{
//   						&secretProperty{
//   							name: jsii.String("name"),
//   							valueFrom: jsii.String("valueFrom"),
//   						},
//   					},
//   				},
//   				memory: jsii.Number(123),
//   				mountPoints: []interface{}{
//   					&mountPointsProperty{
//   						containerPath: jsii.String("containerPath"),
//   						readOnly: jsii.Boolean(false),
//   						sourceVolume: jsii.String("sourceVolume"),
//   					},
//   				},
//   				networkConfiguration: &networkConfigurationProperty{
//   					assignPublicIp: jsii.String("assignPublicIp"),
//   				},
//   				privileged: jsii.Boolean(false),
//   				readonlyRootFilesystem: jsii.Boolean(false),
//   				resourceRequirements: []interface{}{
//   					&resourceRequirementProperty{
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				secrets: []interface{}{
//   					&secretProperty{
//   						name: jsii.String("name"),
//   						valueFrom: jsii.String("valueFrom"),
//   					},
//   				},
//   				ulimits: []interface{}{
//   					&ulimitProperty{
//   						hardLimit: jsii.Number(123),
//   						name: jsii.String("name"),
//   						softLimit: jsii.Number(123),
//   					},
//   				},
//   				user: jsii.String("user"),
//   				vcpus: jsii.Number(123),
//   				volumes: []interface{}{
//   					&volumesProperty{
//   						efsVolumeConfiguration: &efsVolumeConfigurationProperty{
//   							fileSystemId: jsii.String("fileSystemId"),
//
//   							// the properties below are optional
//   							authorizationConfig: &authorizationConfigProperty{
//   								accessPointId: jsii.String("accessPointId"),
//   								iam: jsii.String("iam"),
//   							},
//   							rootDirectory: jsii.String("rootDirectory"),
//   							transitEncryption: jsii.String("transitEncryption"),
//   							transitEncryptionPort: jsii.Number(123),
//   						},
//   						host: &volumesHostProperty{
//   							sourcePath: jsii.String("sourcePath"),
//   						},
//   						name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	numNodes: jsii.Number(123),
//   }
//
type CfnJobDefinition_NodePropertiesProperty struct {
	// Specifies the node index for the main node of a multi-node parallel job.
	//
	// This node index value must be fewer than the number of nodes.
	MainNode *float64 `field:"required" json:"mainNode" yaml:"mainNode"`
	// A list of node ranges and their properties associated with a multi-node parallel job.
	NodeRangeProperties interface{} `field:"required" json:"nodeRangeProperties" yaml:"nodeRangeProperties"`
	// The number of nodes associated with a multi-node parallel job.
	NumNodes *float64 `field:"required" json:"numNodes" yaml:"numNodes"`
}

