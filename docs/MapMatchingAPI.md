# \MapMatchingAPI

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostGPX**](MapMatchingAPI.md#PostGPX) | **Post** /match | Map-match a GPX file



## PostGPX

> RouteResponse PostGPX(ctx).GpsAccuracy(gpsAccuracy).Profile(profile).Locale(locale).Elevation(elevation).Details(details).Instructions(instructions).CalcPoints(calcPoints).PointsEncoded(pointsEncoded).Execute()

Map-match a GPX file



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
	gpsAccuracy := int32(56) // int32 | Specify the precision of a point, in meter (optional)
	profile := "profile_example" // string |  (optional) (default to "car")
	locale := "locale_example" // string | The locale of the resulting turn instructions. E.g. `pt_PT` for Portuguese or `de` for German.  (optional) (default to "en")
	elevation := true // bool | If `true`, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the `points` and `snapped_waypoints` fields of the response, in both their encodings. Unless you switch off the `points_encoded` parameter, you need special code on the client side that can handle three-dimensional coordinates.  (optional) (default to false)
	details := []string{"Inner_example"} // []string | Optional parameter. The following path details are available: `street_name`, `street_ref`, `street_destination`, `leg_time`, `leg_distance`, `roundabout`, `country`, `time`, `distance`, `max_speed`, `max_weight`, `max_width`, `toll`, `road_class`, `road_class_link`, `road_access`, `road_environment`, `hazmat`, `hazmat_tunnel`, `hazmat_water`,  `lanes`, `surface`, `smoothness`, `hike_rating`, `mtb_rating`, `foot_network`, `bike_network`. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  (optional)
	instructions := true // bool | If instructions should be calculated and returned  (optional) (default to true)
	calcPoints := true // bool | If the points for the route should be calculated at all.  (optional) (default to true)
	pointsEncoded := true // bool | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to `false` to switch the encoding to simple coordinate pairs like `[lon,lat]`, or `[lon,lat,elevation]`. See the description of the response format for more information.  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MapMatchingAPI.PostGPX(context.Background()).GpsAccuracy(gpsAccuracy).Profile(profile).Locale(locale).Elevation(elevation).Details(details).Instructions(instructions).CalcPoints(calcPoints).PointsEncoded(pointsEncoded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MapMatchingAPI.PostGPX``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGPX`: RouteResponse
	fmt.Fprintf(os.Stdout, "Response from `MapMatchingAPI.PostGPX`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGPXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpsAccuracy** | **int32** | Specify the precision of a point, in meter | 
 **profile** | **string** |  | [default to &quot;car&quot;]
 **locale** | **string** | The locale of the resulting turn instructions. E.g. &#x60;pt_PT&#x60; for Portuguese or &#x60;de&#x60; for German.  | [default to &quot;en&quot;]
 **elevation** | **bool** | If &#x60;true&#x60;, a third coordinate, the altitude, is included with all positions in the response. This changes the format of the &#x60;points&#x60; and &#x60;snapped_waypoints&#x60; fields of the response, in both their encodings. Unless you switch off the &#x60;points_encoded&#x60; parameter, you need special code on the client side that can handle three-dimensional coordinates.  | [default to false]
 **details** | **[]string** | Optional parameter. The following path details are available: &#x60;street_name&#x60;, &#x60;street_ref&#x60;, &#x60;street_destination&#x60;, &#x60;leg_time&#x60;, &#x60;leg_distance&#x60;, &#x60;roundabout&#x60;, &#x60;country&#x60;, &#x60;time&#x60;, &#x60;distance&#x60;, &#x60;max_speed&#x60;, &#x60;max_weight&#x60;, &#x60;max_width&#x60;, &#x60;toll&#x60;, &#x60;road_class&#x60;, &#x60;road_class_link&#x60;, &#x60;road_access&#x60;, &#x60;road_environment&#x60;, &#x60;hazmat&#x60;, &#x60;hazmat_tunnel&#x60;, &#x60;hazmat_water&#x60;,  &#x60;lanes&#x60;, &#x60;surface&#x60;, &#x60;smoothness&#x60;, &#x60;hike_rating&#x60;, &#x60;mtb_rating&#x60;, &#x60;foot_network&#x60;, &#x60;bike_network&#x60;. Read more about the usage of path details [here](https://discuss.graphhopper.com/t/2539).  | 
 **instructions** | **bool** | If instructions should be calculated and returned  | [default to true]
 **calcPoints** | **bool** | If the points for the route should be calculated at all.  | [default to true]
 **pointsEncoded** | **bool** | Allows changing the encoding of location data in the response. The default is polyline encoding, which is compact but requires special client code to unpack. (We provide it in our JavaScript client library!) Set this parameter to &#x60;false&#x60; to switch the encoding to simple coordinate pairs like &#x60;[lon,lat]&#x60;, or &#x60;[lon,lat,elevation]&#x60;. See the description of the response format for more information.  | [default to true]

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/gpx+xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

