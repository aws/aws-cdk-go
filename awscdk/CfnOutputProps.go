package awscdk


// Example:
//   var cluster cluster
//
//   // add service account
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.GrantReadWrite(serviceAccount)
//
//   mypod := cluster.addManifest(jsii.String("mypod"), map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Pod"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("mypod"),
//   	},
//   	"spec": map[string]interface{}{
//   		"serviceAccountName": serviceAccount.serviceAccountName,
//   		"containers": []map[string]interface{}{
//   			map[string]interface{}{
//   				"name": jsii.String("hello"),
//   				"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   				"ports": []map[string]*f64{
//   					map[string]*f64{
//   						"containerPort": jsii.Number(8080),
//   					},
//   				},
//   			},
//   		},
//   	},
//   })
//
//   // create the resource after the service account.
//   mypod.Node.AddDependency(serviceAccount)
//
//   // print the IAM role arn for this service account
//   // print the IAM role arn for this service account
//   awscdk.NewCfnOutput(this, jsii.String("ServiceAccountIamRole"), &CfnOutputProps{
//   	Value: serviceAccount.Role.RoleArn,
//   })
//
type CfnOutputProps struct {
	// The value of the property returned by the aws cloudformation describe-stacks command.
	//
	// The value of an output can include literals, parameter references, pseudo-parameters,
	// a mapping value, or intrinsic functions.
	Value *string `field:"required" json:"value" yaml:"value"`
	// A condition to associate with this output value.
	//
	// If the condition evaluates
	// to `false`, this output value will not be included in the stack.
	// Default: - No condition is associated with the output.
	//
	Condition CfnCondition `field:"optional" json:"condition" yaml:"condition"`
	// A String type that describes the output value.
	//
	// The description can be a maximum of 4 K in length.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name used to export the value of this output across stacks.
	//
	// To import the value from another stack, use `Fn.importValue(exportName)`.
	// Default: - the output is not exported.
	//
	ExportName *string `field:"optional" json:"exportName" yaml:"exportName"`
	// The key of the property returned by aws cloudformation describe-stacks command.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

