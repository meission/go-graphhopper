# \MatricesAPI

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateMatrix**](MatricesAPI.md#CalculateMatrix) | **Post** /matrix/calculate | Submit a matrix computation job
[**GetMatrix**](MatricesAPI.md#GetMatrix) | **Get** /matrix | Compute a matrix
[**GetMatrixSolution**](MatricesAPI.md#GetMatrixSolution) | **Get** /matrix/solution/{jobId} | Retrieve result of a matrix computation job
[**PostMatrix**](MatricesAPI.md#PostMatrix) | **Post** /matrix | Compute a matrix



## CalculateMatrix

> JobId CalculateMatrix(ctx).PostMatrixRequest(postMatrixRequest).Execute()

Submit a matrix computation job



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
	postMatrixRequest := openapiclient.postMatrix_request{MatrixRequest: openapiclient.NewMatrixRequest()} // PostMatrixRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatricesAPI.CalculateMatrix(context.Background()).PostMatrixRequest(postMatrixRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatricesAPI.CalculateMatrix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateMatrix`: JobId
	fmt.Fprintf(os.Stdout, "Response from `MatricesAPI.CalculateMatrix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postMatrixRequest** | [**PostMatrixRequest**](PostMatrixRequest.md) |  | 

### Return type

[**JobId**](JobId.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatrix

> MatrixResponse GetMatrix(ctx).Profile(profile).Point(point).FromPoint(fromPoint).ToPoint(toPoint).PointHint(pointHint).FromPointHint(fromPointHint).ToPointHint(toPointHint).SnapPrevention(snapPrevention).Curbside(curbside).FromCurbside(fromCurbside).ToCurbside(toCurbside).OutArray(outArray).FailFast(failFast).Execute()

Compute a matrix



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
	profile := "profile_example" // string |  (optional) (default to "car")
	point := []string{"Inner_example"} // []string | Specify multiple points in `latitude,longitude` for which the weight-, route-, time- or distance-matrix should be calculated. In this case the starts are identical to the destinations. If there are N points, then NxN entries will be calculated. The order of the point parameter is important. Specify at least three points. Cannot be used together with from_point or to_point. (optional)
	fromPoint := []string{"Inner_example"} // []string | The origin points for the routes in `latitude,longitude`. E.g. if you want to calculate the three routes A-&gt;1, A-&gt;2, A-&gt;3 then you have one from_point parameter and three to_point parameters. (optional)
	toPoint := []string{"Inner_example"} // []string | The destination points for the routes in `latitude,longitude`. (optional)
	pointHint := []string{"Inner_example"} // []string | Optional parameter. Specifies a hint for each `point` parameter to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. (optional)
	fromPointHint := []string{"Inner_example"} // []string | For the from_point parameter. See point_hint (optional)
	toPointHint := []string{"Inner_example"} // []string | For the to_point parameter. See point_hint (optional)
	snapPrevention := []string{"Inner_example"} // []string | Optional parameter. 'Snapping' is the process of finding the closest road location for GPS coordinates provided in the `point` parameter. The `snap_prevention` parameter allows you to prevent snapping to specific types of roads. For example, if `snap_prevention` is set to bridge, the routing engine will avoid snapping to a bridge, even if it is the closest road for the given `point`. Current supported values: `motorway`, `trunk`, `ferry`, `tunnel`, `bridge` and `ford`. Multiple values are specified like `snap_prevention=ferry&snap_prevention=motorway`. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the `custom_model`.  (optional)
	curbside := []string{"Curbside_example"} // []string | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap. (optional)
	fromCurbside := []string{"FromCurbside_example"} // []string | Curbside setting for the from_point parameter. See curbside. (optional)
	toCurbside := []string{"ToCurbside_example"} // []string | Curbside setting for the to_point parameter. See curbside. (optional)
	outArray := []string{"Inner_example"} // []string | Specifies which arrays should be included in the response. Specify one or more of the following options 'weights', 'times', 'distances'. To specify more than one array use e.g. out_array=times&out_array=distances. The units of the entries of distances are meters, of times are seconds and of weights is arbitrary and it can differ for different vehicles or versions of this API. (optional)
	failFast := true // bool | Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to `false` the time/weight/distance matrix will be calculated for all valid points and contain the `null` value for all entries that could not be calculated. The `hint` field of the response will also contain additional information about what went wrong (see its documentation). (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatricesAPI.GetMatrix(context.Background()).Profile(profile).Point(point).FromPoint(fromPoint).ToPoint(toPoint).PointHint(pointHint).FromPointHint(fromPointHint).ToPointHint(toPointHint).SnapPrevention(snapPrevention).Curbside(curbside).FromCurbside(fromCurbside).ToCurbside(toCurbside).OutArray(outArray).FailFast(failFast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatricesAPI.GetMatrix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMatrix`: MatrixResponse
	fmt.Fprintf(os.Stdout, "Response from `MatricesAPI.GetMatrix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | **string** |  | [default to &quot;car&quot;]
 **point** | **[]string** | Specify multiple points in &#x60;latitude,longitude&#x60; for which the weight-, route-, time- or distance-matrix should be calculated. In this case the starts are identical to the destinations. If there are N points, then NxN entries will be calculated. The order of the point parameter is important. Specify at least three points. Cannot be used together with from_point or to_point. | 
 **fromPoint** | **[]string** | The origin points for the routes in &#x60;latitude,longitude&#x60;. E.g. if you want to calculate the three routes A-&amp;gt;1, A-&amp;gt;2, A-&amp;gt;3 then you have one from_point parameter and three to_point parameters. | 
 **toPoint** | **[]string** | The destination points for the routes in &#x60;latitude,longitude&#x60;. | 
 **pointHint** | **[]string** | Optional parameter. Specifies a hint for each &#x60;point&#x60; parameter to prefer a certain street for the closest location lookup. E.g. if there is an address or house with two or more neighboring streets you can control for which street the closest location is looked up. | 
 **fromPointHint** | **[]string** | For the from_point parameter. See point_hint | 
 **toPointHint** | **[]string** | For the to_point parameter. See point_hint | 
 **snapPrevention** | **[]string** | Optional parameter. &#39;Snapping&#39; is the process of finding the closest road location for GPS coordinates provided in the &#x60;point&#x60; parameter. The &#x60;snap_prevention&#x60; parameter allows you to prevent snapping to specific types of roads. For example, if &#x60;snap_prevention&#x60; is set to bridge, the routing engine will avoid snapping to a bridge, even if it is the closest road for the given &#x60;point&#x60;. Current supported values: &#x60;motorway&#x60;, &#x60;trunk&#x60;, &#x60;ferry&#x60;, &#x60;tunnel&#x60;, &#x60;bridge&#x60; and &#x60;ford&#x60;. Multiple values are specified like &#x60;snap_prevention&#x3D;ferry&amp;snap_prevention&#x3D;motorway&#x60;. Note that once snapped the routing algorithm can still route over bridges (or the other values). To avoid this you need to use the &#x60;custom_model&#x60;.  | 
 **curbside** | **[]string** | Optional parameter. It specifies on which side a point should be relative to the driver when she leaves/arrives at a start/target/via point. You need to specify this parameter for either none or all points. Only supported for motor vehicles and OpenStreetMap. | 
 **fromCurbside** | **[]string** | Curbside setting for the from_point parameter. See curbside. | 
 **toCurbside** | **[]string** | Curbside setting for the to_point parameter. See curbside. | 
 **outArray** | **[]string** | Specifies which arrays should be included in the response. Specify one or more of the following options &#39;weights&#39;, &#39;times&#39;, &#39;distances&#39;. To specify more than one array use e.g. out_array&#x3D;times&amp;out_array&#x3D;distances. The units of the entries of distances are meters, of times are seconds and of weights is arbitrary and it can differ for different vehicles or versions of this API. | 
 **failFast** | **bool** | Specifies whether or not the matrix calculation should return with an error as soon as possible in case some points cannot be found or some points are not connected. If set to &#x60;false&#x60; the time/weight/distance matrix will be calculated for all valid points and contain the &#x60;null&#x60; value for all entries that could not be calculated. The &#x60;hint&#x60; field of the response will also contain additional information about what went wrong (see its documentation). | [default to true]

### Return type

[**MatrixResponse**](MatrixResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMatrixSolution

> GetMatrixSolution200Response GetMatrixSolution(ctx, jobId).Execute()

Retrieve result of a matrix computation job

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
	jobId := "jobId_example" // string | The jobId you received when you submitted the job.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatricesAPI.GetMatrixSolution(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatricesAPI.GetMatrixSolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMatrixSolution`: GetMatrixSolution200Response
	fmt.Fprintf(os.Stdout, "Response from `MatricesAPI.GetMatrixSolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The jobId you received when you submitted the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMatrixSolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMatrixSolution200Response**](GetMatrixSolution200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMatrix

> MatrixResponse PostMatrix(ctx).PostMatrixRequest(postMatrixRequest).Execute()

Compute a matrix



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
	postMatrixRequest := openapiclient.postMatrix_request{MatrixRequest: openapiclient.NewMatrixRequest()} // PostMatrixRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MatricesAPI.PostMatrix(context.Background()).PostMatrixRequest(postMatrixRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MatricesAPI.PostMatrix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMatrix`: MatrixResponse
	fmt.Fprintf(os.Stdout, "Response from `MatricesAPI.PostMatrix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postMatrixRequest** | [**PostMatrixRequest**](PostMatrixRequest.md) |  | 

### Return type

[**MatrixResponse**](MatrixResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

