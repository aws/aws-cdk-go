package awsapplicationsignals


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   compositeSliComponentProperty := &CompositeSliComponentProperty{
//   	OperationName: jsii.String("operationName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositeslicomponent.html
//
type CfnServiceLevelObjective_CompositeSliComponentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationsignals-servicelevelobjective-compositeslicomponent.html#cfn-applicationsignals-servicelevelobjective-compositeslicomponent-operationname
	//
	OperationName *string `field:"required" json:"operationName" yaml:"operationName"`
}

