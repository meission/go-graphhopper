# \RoutingAPI

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoute**](RoutingAPI.md#GetRoute) | **Get** /route | Calculate a route
[**PostRoute**](RoutingAPI.md#PostRoute) | **Post** /route | Calculate a route



## GetRoute

> RouteResponse GetRoute(ctx).Point(point).Profile(profile).PointHint(pointHint).SnapPrevention(snapPrevention).Curbside(curbside).Locale(locale).Elevation(elevation).Details(details).Optimize(optimize).Instructions(instructions).CalcPoints(calcPoints).Debug(debug).PointsEncoded(pointsEncoded).ChDisable(chDisable).Heading(heading).HeadingPenalty(headingPenalty).PassThrough(passThrough).Algorithm(algorithm).RoundTripDistance(roundTripDistance).RoundTripSeed(roundTripSeed).AlternativeRouteMaxPaths(alternativeRouteMaxPaths).AlternativeRouteMaxWeightFactor(alternativeRouteMaxWeightFactor).AlternativeRouteMaxShareFactor(alternativeRouteMaxShareFactor).Execute()

Calculate a route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/meission/go-graphhopper"
)

func main() {
	point := []string{"Inner_example"} // []string | The points for which the route should be calculated. Format: `latitude,longitude`. Specify at least an origin and a destination. Via points are possible. The maximum number depends on your plan. 
	profile := "profile_example" // string |  (optional) (default to "car")
	pointHint := []string{"Inner_example"} // []string | The `point_hint` is typically a road name to which the associated `point` parameter should be snapped to. Specify no `point_hint` parameter or the same number as you have `point` parameters.  (optional)
	snapPrevention := []string{"Inner_example"} // []string | Optional parameter. 'Snapping' is the process of finding the closest road location for GPS coordinates provided in the `point` parameter. The `snap_prevention` parameter allows you to prevent snapping to specific types of roads. For example, if `snap_prevention` is set to bridge, the routing engine will avoid snapping to a bridge, even if it is the closest road for the given `point`. Current supported values: `motorway`, `trunk`, `ferry`, `tunnel`, `bridge` and `ford`. Multiple values are specified like `snap_prevention=ferry&snap_prevention=motorway`. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the `custom_model`.  (optional)
	curbside := []string{"Curbside_example"} // []string | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap.  (optional)
	locale := "locale_example" // string | The locale of the resulting turn instructions. E.g. `pt_PT` for Portuguese or `de` for German.  (optional) (default to "en")
	elevation := true // bool | If `true`, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the `points` and `snapped_waypoints` fields of the response, in both their encodings. Unless you switch off the `points_encoded` parameter, you need special code on the client side that can handle three-dimensional coordinates.  (optional) (default to false)
	details := []string{"Inner_example"} // []string | Optional parameter. The following path details are available: `street_name`, `street_ref`, `street_destination`, `leg_time`, `leg_distance`, `roundabout`, `country`, `time`, `distance`, `max_speed`, `max_weight`, `max_width`, `toll`, `road_class`, `road_class_link`, `road_access`, `road_environment`, `hazmat`, `hazmat_tunnel`, `hazmat_water`,  `lanes`, `surface`, `smoothness`, `hike_rating`, `mtb_rating`, `foot_network`, `bike_network`. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  (optional)
	optimize := "optimize_example" // string | Normally, the calculated route will visit the points in the order you specified them. If you have more than two points, you can set this parameter to `\"true\"` and the points may be re-ordered to minimize the total travel time. Keep in mind that the limits on the number of locations of the Route Optimization API applies, and the request costs more credits.  (optional) (default to "false")
	instructions := true // bool | If instructions should be calculated and returned  (optional) (default to true)
	calcPoints := true // bool | If the points for the route should be calculated at all.  (optional) (default to true)
	debug := true // bool | If `true`, the output will be formatted.  (optional) (default to false)
	pointsEncoded := true // bool | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to `false` to switch the encoding to simple coordinate pairs like `[lon,lat]`, or `[lon,lat,elevation]`. See the description of the response format for more information.  (optional) (default to true)
	chDisable := true // bool | Combine this parameter with any of the following options below.  (optional) (default to false)
	heading := []int32{int32(123)} // []int32 | Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with `algorithm=round_trip` and forces the initial direction.  Requires `ch.disable=true`.  (optional)
	headingPenalty := int32(56) // int32 | Time penalty in seconds for not obeying a specified heading. Requires `ch.disable=true`.  (optional) (default to 300)
	passThrough := true // bool | If `true`, u-turns are avoided at via-points with regard to the `heading_penalty`. Requires `ch.disable=true`.  (optional) (default to false)
	algorithm := "algorithm_example" // string | Rather than looking for the shortest or fastest path, this parameter lets you solve two different problems related to routing: With `alternative_route`, we give you not one but several routes that are close to optimal, but not too similar to each other.  With `round_trip`, the route will get you back to where you started. This is meant for fun (think of a bike trip), so we will add some randomness. The `round_trip` option requires `ch.disable=true`. You can control both of these features with additional parameters, see below.   (optional)
	roundTripDistance := int32(56) // int32 | If `algorithm=round_trip`, this parameter configures approximative length of the resulting round trip. Requires `ch.disable=true`.  (optional) (default to 10000)
	roundTripSeed := int64(789) // int64 | If `algorithm=round_trip`, this sets the random seed. Change this to get a different tour for each value.  (optional)
	alternativeRouteMaxPaths := int32(56) // int32 | If `algorithm=alternative_route`, this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives.  (optional) (default to 2)
	alternativeRouteMaxWeightFactor := float32(8.14) // float32 | If `algorithm=alternative_route`, this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives.  (optional) (default to 1.4)
	alternativeRouteMaxShareFactor := float32(8.14) // float32 | If `algorithm=alternative_route`, this parameter specifies how similar an alternative route can be to the optimal route. Increasing can lead to worse alternatives.  (optional) (default to 0.6)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.GetRoute(context.Background()).Point(point).Profile(profile).PointHint(pointHint).SnapPrevention(snapPrevention).Curbside(curbside).Locale(locale).Elevation(elevation).Details(details).Optimize(optimize).Instructions(instructions).CalcPoints(calcPoints).Debug(debug).PointsEncoded(pointsEncoded).ChDisable(chDisable).Heading(heading).HeadingPenalty(headingPenalty).PassThrough(passThrough).Algorithm(algorithm).RoundTripDistance(roundTripDistance).RoundTripSeed(roundTripSeed).AlternativeRouteMaxPaths(alternativeRouteMaxPaths).AlternativeRouteMaxWeightFactor(alternativeRouteMaxWeightFactor).AlternativeRouteMaxShareFactor(alternativeRouteMaxShareFactor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.GetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoute`: RouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.GetRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **point** | **[]string** | The points for which the route should be calculated. Format: &#x60;latitude,longitude&#x60;. Specify at least an origin and a destination. Via points are possible. The maximum number depends on your plan.  | 
 **profile** | **string** |  | [default to &quot;car&quot;]
 **pointHint** | **[]string** | The &#x60;point_hint&#x60; is typically a road name to which the associated &#x60;point&#x60; parameter should be snapped to. Specify no &#x60;point_hint&#x60; parameter or the same number as you have &#x60;point&#x60; parameters.  | 
 **snapPrevention** | **[]string** | Optional parameter. &#39;Snapping&#39; is the process of finding the closest road location for GPS coordinates provided in the &#x60;point&#x60; parameter. The &#x60;snap_prevention&#x60; parameter allows you to prevent snapping to specific types of roads. For example, if &#x60;snap_prevention&#x60; is set to bridge, the routing engine will avoid snapping to a bridge, even if it is the closest road for the given &#x60;point&#x60;. Current supported values: &#x60;motorway&#x60;, &#x60;trunk&#x60;, &#x60;ferry&#x60;, &#x60;tunnel&#x60;, &#x60;bridge&#x60; and &#x60;ford&#x60;. Multiple values are specified like &#x60;snap_prevention&#x3D;ferry&amp;snap_prevention&#x3D;motorway&#x60;. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the &#x60;custom_model&#x60;.  | 
 **curbside** | **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap.  | 
 **locale** | **string** | The locale of the resulting turn instructions. E.g. &#x60;pt_PT&#x60; for Portuguese or &#x60;de&#x60; for German.  | [default to &quot;en&quot;]
 **elevation** | **bool** | If &#x60;true&#x60;, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the &#x60;points&#x60; and &#x60;snapped_waypoints&#x60; fields of the response, in both their encodings. Unless you switch off the &#x60;points_encoded&#x60; parameter, you need special code on the client side that can handle three-dimensional coordinates.  | [default to false]
 **details** | **[]string** | Optional parameter. The following path details are available: &#x60;street_name&#x60;, &#x60;street_ref&#x60;, &#x60;street_destination&#x60;, &#x60;leg_time&#x60;, &#x60;leg_distance&#x60;, &#x60;roundabout&#x60;, &#x60;country&#x60;, &#x60;time&#x60;, &#x60;distance&#x60;, &#x60;max_speed&#x60;, &#x60;max_weight&#x60;, &#x60;max_width&#x60;, &#x60;toll&#x60;, &#x60;road_class&#x60;, &#x60;road_class_link&#x60;, &#x60;road_access&#x60;, &#x60;road_environment&#x60;, &#x60;hazmat&#x60;, &#x60;hazmat_tunnel&#x60;, &#x60;hazmat_water&#x60;,  &#x60;lanes&#x60;, &#x60;surface&#x60;, &#x60;smoothness&#x60;, &#x60;hike_rating&#x60;, &#x60;mtb_rating&#x60;, &#x60;foot_network&#x60;, &#x60;bike_network&#x60;. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  | 
 **optimize** | **string** | Normally, the calculated route will visit the points in the order you specified them. If you have more than two points, you can set this parameter to &#x60;\&quot;true\&quot;&#x60; and the points may be re-ordered to minimize the total travel time. Keep in mind that the limits on the number of locations of the Route Optimization API applies, and the request costs more credits.  | [default to &quot;false&quot;]
 **instructions** | **bool** | If instructions should be calculated and returned  | [default to true]
 **calcPoints** | **bool** | If the points for the route should be calculated at all.  | [default to true]
 **debug** | **bool** | If &#x60;true&#x60;, the output will be formatted.  | [default to false]
 **pointsEncoded** | **bool** | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to &#x60;false&#x60; to switch the encoding to simple coordinate pairs like &#x60;[lon,lat]&#x60;, or &#x60;[lon,lat,elevation]&#x60;. See the description of the response format for more information.  | [default to true]
 **chDisable** | **bool** | Combine this parameter with any of the following options below.  | [default to false]
 **heading** | **[]int32** | Favour a heading direction for a certain point. Specify either one heading for the start point or as many as there are points. In this case headings are associated by their order to the specific points. Headings are given as north based clockwise angle between 0 and 360 degree. This parameter also influences the tour generated with &#x60;algorithm&#x3D;round_trip&#x60; and forces the initial direction.  Requires &#x60;ch.disable&#x3D;true&#x60;.  | 
 **headingPenalty** | **int32** | Time penalty in seconds for not obeying a specified heading. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [default to 300]
 **passThrough** | **bool** | If &#x60;true&#x60;, u-turns are avoided at via-points with regard to the &#x60;heading_penalty&#x60;. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [default to false]
 **algorithm** | **string** | Rather than looking for the shortest or fastest path, this parameter lets you solve two different problems related to routing: With &#x60;alternative_route&#x60;, we give you not one but several routes that are close to optimal, but not too similar to each other.  With &#x60;round_trip&#x60;, the route will get you back to where you started. This is meant for fun (think of a bike trip), so we will add some randomness. The &#x60;round_trip&#x60; option requires &#x60;ch.disable&#x3D;true&#x60;. You can control both of these features with additional parameters, see below.   | 
 **roundTripDistance** | **int32** | If &#x60;algorithm&#x3D;round_trip&#x60;, this parameter configures approximative length of the resulting round trip. Requires &#x60;ch.disable&#x3D;true&#x60;.  | [default to 10000]
 **roundTripSeed** | **int64** | If &#x60;algorithm&#x3D;round_trip&#x60;, this sets the random seed. Change this to get a different tour for each value.  | 
 **alternativeRouteMaxPaths** | **int32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter sets the number of maximum paths which should be calculated. Increasing can lead to worse alternatives.  | [default to 2]
 **alternativeRouteMaxWeightFactor** | **float32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter sets the factor by which the alternatives routes can be longer than the optimal route. Increasing can lead to worse alternatives.  | [default to 1.4]
 **alternativeRouteMaxShareFactor** | **float32** | If &#x60;algorithm&#x3D;alternative_route&#x60;, this parameter specifies how similar an alternative route can be to the optimal route. Increasing can lead to worse alternatives.  | [default to 0.6]

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRoute

> RouteResponse PostRoute(ctx).RouteRequest(routeRequest).Execute()

Calculate a route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/meission/go-graphhopper"
)

func main() {
	routeRequest := *openapiclient.NewRouteRequest() // RouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutingAPI.PostRoute(context.Background()).RouteRequest(routeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingAPI.PostRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRoute`: RouteResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutingAPI.PostRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routeRequest** | [**RouteRequest**](RouteRequest.md) |  | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

