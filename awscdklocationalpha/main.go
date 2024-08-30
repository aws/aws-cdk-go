// The CDK Construct Library for AWS::Location
package awscdklocationalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
		map[string]interface{}{
			"ESRI": DataSource_ESRI,
			"HERE": DataSource_HERE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.GeofenceCollection",
		reflect.TypeOf((*GeofenceCollection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionArn", GoGetter: "GeofenceCollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionCreateTime", GoGetter: "GeofenceCollectionCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionName", GoGetter: "GeofenceCollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionUpdateTime", GoGetter: "GeofenceCollectionUpdateTime"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GeofenceCollection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGeofenceCollection)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.GeofenceCollectionProps",
		reflect.TypeOf((*GeofenceCollectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IGeofenceCollection",
		reflect.TypeOf((*IGeofenceCollection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionArn", GoGetter: "GeofenceCollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "geofenceCollectionName", GoGetter: "GeofenceCollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IGeofenceCollection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-location-alpha.IPlaceIndex",
		reflect.TypeOf((*IPlaceIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexArn", GoGetter: "PlaceIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexName", GoGetter: "PlaceIndexName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IPlaceIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-location-alpha.IntendedUse",
		reflect.TypeOf((*IntendedUse)(nil)).Elem(),
		map[string]interface{}{
			"SINGLE_USE": IntendedUse_SINGLE_USE,
			"STORAGE": IntendedUse_STORAGE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-location-alpha.PlaceIndex",
		reflect.TypeOf((*PlaceIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantSearch", GoMethod: "GrantSearch"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexArn", GoGetter: "PlaceIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexCreateTime", GoGetter: "PlaceIndexCreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexName", GoGetter: "PlaceIndexName"},
			_jsii_.MemberProperty{JsiiProperty: "placeIndexUpdateTime", GoGetter: "PlaceIndexUpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PlaceIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPlaceIndex)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-location-alpha.PlaceIndexProps",
		reflect.TypeOf((*PlaceIndexProps)(nil)).Elem(),
	)
}
