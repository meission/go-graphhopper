# \CustomProfilesAPI

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProfile**](CustomProfilesAPI.md#DeleteProfile) | **Delete** /profiles/{profileId} | Delete a custom routing profile
[**GetProfile**](CustomProfilesAPI.md#GetProfile) | **Get** /profiles | List your custom routing profiles
[**PostProfile**](CustomProfilesAPI.md#PostProfile) | **Post** /profiles | Create a custom routing profile



## DeleteProfile

> DeleteProfile(ctx, profileId).Execute()

Delete a custom routing profile

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
	profileId := "cp_shc_1_4390ba24-e6cb-406b-acc6-3402dd4d229d" // string | The profile to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomProfilesAPI.DeleteProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The profile to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> ProfileGetResponse GetProfile(ctx).Execute()

List your custom routing profiles

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomProfilesAPI.GetProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: ProfileGetResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


### Return type

[**ProfileGetResponse**](ProfileGetResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProfile

> ProfileResponse PostProfile(ctx).ProfileRequest(profileRequest).Execute()

Create a custom routing profile

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
	profileRequest := *openapiclient.NewProfileRequest() // ProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomProfilesAPI.PostProfile(context.Background()).ProfileRequest(profileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.PostProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProfile`: ProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.PostProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileRequest** | [**ProfileRequest**](ProfileRequest.md) |  | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

