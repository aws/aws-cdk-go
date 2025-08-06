package awsapplicationsignals


// Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dependencyConfigProperty := &DependencyConfigProperty{
//   	DependencyKeyAttributes: map[string]*string{
//   		"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   	},
//   	DependencyOperationName: jsii.String("dependencyOperationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-dependencyconfig.html
//
type CfnServiceLevelObjective_DependencyConfigProperty struct {
	// If this SLO is related to a metric collected by Application Signals, you must use this field to specify which dependency the SLO metric is related to.
	//
	// - `Type` designates the type of object this is.
	// - `ResourceType` specifies the type of the resource. This field is used only when the value of the `Type` field is `Resource` or `AWS::Resource` .
	// - `Name` specifies the name of the object. This is used only if the value of the `Type` field is `Service` , `RemoteService` , or `AWS::Service` .
	// - `Identifier` identifies the resource objects of this resource. This is used only if the value of the `Type` field is `Resource` or `AWS::Resource` .
	// - `Environment` specifies the location where this object is hosted, or what it belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-dependencyconfig.html#cfn-applicationsignals-servicelevelobjective-dependencyconfig-dependencykeyattributes
	//
	DependencyKeyAttributes interface{} `field:"required" json:"dependencyKeyAttributes" yaml:"dependencyKeyAttributes"`
	// When the SLO monitors a specific operation of the dependency, this field specifies the name of that operation in the dependency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-dependencyconfig.html#cfn-applicationsignals-servicelevelobjective-dependencyconfig-dependencyoperationname
	//
	DependencyOperationName *string `field:"required" json:"dependencyOperationName" yaml:"dependencyOperationName"`
}

