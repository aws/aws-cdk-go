package mixinsawsglue


// Properties for CfnDevEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var arguments_ interface{}
//   var tags interface{}
//
//   cfnDevEndpointMixinProps := &CfnDevEndpointMixinProps{
//   	Arguments: arguments_,
//   	EndpointName: jsii.String("endpointName"),
//   	ExtraJarsS3Path: jsii.String("extraJarsS3Path"),
//   	ExtraPythonLibsS3Path: jsii.String("extraPythonLibsS3Path"),
//   	GlueVersion: jsii.String("glueVersion"),
//   	NumberOfNodes: jsii.Number(123),
//   	NumberOfWorkers: jsii.Number(123),
//   	PublicKey: jsii.String("publicKey"),
//   	PublicKeys: []*string{
//   		jsii.String("publicKeys"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: tags,
//   	WorkerType: jsii.String("workerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html
//
type CfnDevEndpointMixinProps struct {
	// A map of arguments used to configure the `DevEndpoint` .
	//
	// Valid arguments are:
	//
	// - `"--enable-glue-datacatalog": ""`
	// - `"GLUE_PYTHON_VERSION": "3"`
	// - `"GLUE_PYTHON_VERSION": "2"`
	//
	// You can specify a version of Python support for development endpoints by using the `Arguments` parameter in the `CreateDevEndpoint` or `UpdateDevEndpoint` APIs. If no arguments are provided, the version defaults to Python 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-arguments
	//
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// The name of the `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// The path to one or more Java `.jar` files in an S3 bucket that should be loaded in your `DevEndpoint` .
	//
	// > You can only use pure Java/Scala libraries with a `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-extrajarss3path
	//
	ExtraJarsS3Path *string `field:"optional" json:"extraJarsS3Path" yaml:"extraJarsS3Path"`
	// The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded in your `DevEndpoint` .
	//
	// Multiple values must be complete paths separated by a comma.
	//
	// > You can only use pure Python libraries with a `DevEndpoint` . Libraries that rely on C extensions, such as the [pandas](https://docs.aws.amazon.com/http://pandas.pydata.org/) Python data analysis library, are not currently supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-extrapythonlibss3path
	//
	ExtraPythonLibsS3Path *string `field:"optional" json:"extraPythonLibsS3Path" yaml:"extraPythonLibsS3Path"`
	// The AWS Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// The Python version indicates the version supported for running your ETL scripts on development endpoints.
	//
	// For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.
	//
	// Development endpoints that are created without specifying a Glue version default to Glue 0.9.
	//
	// You can specify a version of Python support for development endpoints by using the `Arguments` parameter in the `CreateDevEndpoint` or `UpdateDevEndpoint` APIs. If no arguments are provided, the version defaults to Python 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-glueversion
	//
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// The number of AWS Glue Data Processing Units (DPUs) allocated to this `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-numberofnodes
	//
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// The number of workers of a defined `workerType` that are allocated to the development endpoint.
	//
	// The maximum number of workers you can define are 299 for `G.1X` , and 149 for `G.2X` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-numberofworkers
	//
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The public key to be used by this `DevEndpoint` for authentication.
	//
	// This attribute is provided for backward compatibility because the recommended attribute to use is public keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-publickey
	//
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// A list of public keys to be used by the `DevEndpoints` for authentication.
	//
	// Using this attribute is preferred over a single public key because the public keys allow you to have a different private key per client.
	//
	// > If you previously created an endpoint with a public key, you must remove that key to be able to set a list of public keys. Call the `UpdateDevEndpoint` API operation with the public key content in the `deletePublicKeys` attribute, and the list of new keys in the `addPublicKeys` attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-publickeys
	//
	PublicKeys *[]*string `field:"optional" json:"publicKeys" yaml:"publicKeys"`
	// The Amazon Resource Name (ARN) of the IAM role used in this `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the `SecurityConfiguration` structure to be used with this `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-securityconfiguration
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// A list of security group identifiers used in this `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet ID for this `DevEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The tags to use with this DevEndpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of predefined worker that is allocated to the development endpoint.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	// - For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	//
	// Known issue: when a development endpoint is created with the `G.2X` `WorkerType` configuration, the Spark drivers for the development endpoint will run on 4 vCPU, 16 GB of memory, and a 64 GB disk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-devendpoint.html#cfn-glue-devendpoint-workertype
	//
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

