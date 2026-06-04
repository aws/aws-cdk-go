// The CDK Construct Library for AWS::DSQL
package awscdkdsqlalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-dsql-alpha.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnectAdmin", GoMethod: "GrantConnectAdmin"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceName", GoGetter: "VpcEndpointServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-dsql-alpha.ClusterAttributes",
		reflect.TypeOf((*ClusterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-dsql-alpha.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-dsql-alpha.ICluster",
		reflect.TypeOf((*ICluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnect", GoMethod: "GrantConnect"},
			_jsii_.MemberMethod{JsiiMethod: "grantConnectAdmin", GoMethod: "GrantConnectAdmin"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpcEndpointServiceName", GoGetter: "VpcEndpointServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ICluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
}
