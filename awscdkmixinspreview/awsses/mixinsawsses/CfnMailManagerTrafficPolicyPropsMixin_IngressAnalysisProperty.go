package mixinsawsses


// The Add On ARN and its returned value that is evaluated in a policy statement's conditional expression to either deny or block the incoming email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressAnalysisProperty := &IngressAnalysisProperty{
//   	Analyzer: jsii.String("analyzer"),
//   	ResultField: jsii.String("resultField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressanalysis.html
//
type CfnMailManagerTrafficPolicyPropsMixin_IngressAnalysisProperty struct {
	// The Amazon Resource Name (ARN) of an Add On.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressanalysis.html#cfn-ses-mailmanagertrafficpolicy-ingressanalysis-analyzer
	//
	Analyzer *string `field:"optional" json:"analyzer" yaml:"analyzer"`
	// The returned value from an Add On.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressanalysis.html#cfn-ses-mailmanagertrafficpolicy-ingressanalysis-resultfield
	//
	ResultField *string `field:"optional" json:"resultField" yaml:"resultField"`
}

