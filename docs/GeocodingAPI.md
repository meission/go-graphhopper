# \GeocodingAPI

All URIs are relative to *https://graphhopper.com/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeocode**](GeocodingAPI.md#GetGeocode) | **Get** /geocode | Geocoding Endpoint



## GetGeocode

> GeocodingResponse GetGeocode(ctx).Q(q).Locale(locale).Limit(limit).Reverse(reverse).Debug(debug).Point(point).Provider(provider).Execute()

Geocoding Endpoint

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
	q := "q_example" // string | A textual description of the address you are looking for. Required for forward geocoding. Note that the `default` geocoding provider does prefix searches preferable for \"autocomplete\" use cases, but may lead to sub-optimal results if used without user interaction. See e.g. `provider=nominatim` as an appropriate alternative for less interactive use cases.  (optional)
	locale := "locale_example" // string | Display the search results for the specified locale. Currently French (fr), English (en) and German (de) are explicitly supported. Otherwise leave the locale empty. (optional) (default to "en")
	limit := int32(56) // int32 | Specify the maximum number of results to return (optional) (default to 5)
	reverse := true // bool | It is `required` to be `true` if you want to do a reverse geocoding request. If it is `true`, `point` must be defined as well, and `q` must not be used. (optional) (default to false)
	debug := true // bool | If `true`, the output will be formatted. (optional) (default to false)
	point := "point_example" // string | _Forward geocoding_: The location bias in the format 'latitude,longitude' e.g. point=45.93272,11.58803. _Reverse geocoding_: The location to find amenities, cities. (optional)
	provider := "provider_example" // string | The provider parameter is currently under development and can fall back to `default` at any time. The intend is to provide alternatives to our default geocoder. Each provider has its own strenghts and might fit better for certain scenarios, so it's worth to compare the different providers. To try it append the `provider`parameter to the URL like `&provider=nominatim`, the result structure should be identical in all cases - if not, please report this back to us. Keep in mind that some providers do not support certain parameters or don't return some fields, for example `osm_id` and `osm_type` are not supported by every geocoding provider. If you would like to use additional parameters of one of the providers, but it's not available for the GraphHopper Geocoding API, yet? Please contact us.  The credit costs can be different for all providers - see [here](https://support.graphhopper.com/support/solutions/articles/44000718211-what-is-one-credit-) for more information about it.  Currently, only the default provider and gisgraphy support autocompletion of partial search strings.  All providers support normal \"forward\" geocoding and reverse geocoding via `reverse=true`.  #### Default (`provider=default`)  This provider returns results of our internal geocoding engine. It is best suited for use cases with user interaction as it does prefix searches required for an \"autocomplete\" use case (a user types an address into a search field).  In addition to the above documented parameters the following parameters are possible:  * `osm_tag` - you can filter `key:value` or exclude places with certain OpenStreetMap tags `!key:value`. E.g. `osm_tag=tourism:museum` or just the key `osm_tag=tourism`. To exclude multiple tags you add multiple `osm_tag` parameters. * `location_bias_scale` - describes how much the prominence of a result should still be taken into account. Sensible values go from 0.0 (ignore prominence almost completely) to 1.0 (prominence has approximately the same influence as the location). The default is 0.2. * `zoom` - describes the radius around the center to focus on. This is a number that should correspond roughly to the map zoom parameter of a corresponding map. The default is zoom=16. * `bbox` - the expected format is `minLon,minLat,maxLon,maxLat`. This requires reverse=false. * `radius` - the search radius in km for the reverse search. This requires reverse=true.  #### Nominatim (`provider=nominatim`)  The GraphHopper Directions API uses a commercially hosted Nominatim geocoder (hosted by [OpenCageData](https://opencagedata.com/)). It is best suited for use cases without or less user interaction like batch processing or detailed location data retrieval. It is not suited for \"autocomplete\".  In addition to the above documented parameters we currently support the following parameters:  * countrycode - The country code is a two letter code as defined by the ISO 3166-1 Alpha 2 standard. E.g. gb for the United Kingdom, fr for France, us for United States. * bounds - the expected format is `minLon,minLat,maxLon,maxLat`  #### Gisgraphy (`provider=gisgraphy`)  This provider returns results from the Gisgraphy geocoder which you can try [here](https://services.gisgraphy.com/static/leaflet/index.html).  **Limitations:** The `locale` parameter is not supported. Gisgraphy does not return OSM tags or an extent.  Gisgraphy has a special autocomplete API, which you can use by adding `autocomplete=true` (does not work with `reverse=true`). The autocomplete API is optimized on predicting text input, but returns less information.  In addition to the above documented parameters Gisgraphy allows to use the following parameters, which can be used as documented [here](https://www.gisgraphy.com/documentation/user-guide.php#geocodingservice):  * `radius` - radius in meters * `country` - restrict search for the specified country. The value must be the ISO 3166 Alpha 2 code of the country.  #### NetToolKit (`provider=nettoolkit`)  This provider returns results from the NetToolKit provider which is specialized for US addresses and provides a wrapper around Nominatim for other addresses. You can try it [here](https://www.nettoolkit.com/geo/demo).  The following additional NetToolKit parameters are supported (read [here](https://www.nettoolkit.com/docs/geo/geocoding) for more details): - `source`: User can choose which source provider to geocode the address, this value is \"NetToolKit\" by default - `country_code`: an iso-3166-2 country code (e.g : US) filter the results to the specify country code  **Limitations:** NetToolKit does not support the `locale` parameter. NetToolKit does not return OSM tags (e.g. osm_id, osm_type, osm_value).  #### OpenCage Data (`provider=opencagedata`)  This provider returns results from the OpenCageData geocoder which you can try [here](https://geocoder.opencagedata.com/demo). The difference to the `nominatim` provider is that [other geocoders](https://opencagedata.com/credits) might be used under the hood.  In addition to the above documented parameters OpenCage Data allows to use the following parameters, which can be used as documented [here](https://geocoder.opencagedata.com/api#forward-opt):  * countrycode - The country code is a two letter code as defined by the ISO 3166-1 Alpha 2 standard. E.g. gb for the United Kingdom, fr for France, us for United States.  * bounds - the expected format is `minLon,minLat,maxLon,maxLat`  (optional) (default to "default")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeocodingAPI.GetGeocode(context.Background()).Q(q).Locale(locale).Limit(limit).Reverse(reverse).Debug(debug).Point(point).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeocodingAPI.GetGeocode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeocode`: GeocodingResponse
	fmt.Fprintf(os.Stdout, "Response from `GeocodingAPI.GetGeocode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGeocodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | A textual description of the address you are looking for. Required for forward geocoding. Note that the &#x60;default&#x60; geocoding provider does prefix searches preferable for \&quot;autocomplete\&quot; use cases, but may lead to sub-optimal results if used without user interaction. See e.g. &#x60;provider&#x3D;nominatim&#x60; as an appropriate alternative for less interactive use cases.  | 
 **locale** | **string** | Display the search results for the specified locale. Currently French (fr), English (en) and German (de) are explicitly supported. Otherwise leave the locale empty. | [default to &quot;en&quot;]
 **limit** | **int32** | Specify the maximum number of results to return | [default to 5]
 **reverse** | **bool** | It is &#x60;required&#x60; to be &#x60;true&#x60; if you want to do a reverse geocoding request. If it is &#x60;true&#x60;, &#x60;point&#x60; must be defined as well, and &#x60;q&#x60; must not be used. | [default to false]
 **debug** | **bool** | If &#x60;true&#x60;, the output will be formatted. | [default to false]
 **point** | **string** | _Forward geocoding_: The location bias in the format &#39;latitude,longitude&#39; e.g. point&#x3D;45.93272,11.58803. _Reverse geocoding_: The location to find amenities, cities. | 
 **provider** | **string** | The provider parameter is currently under development and can fall back to &#x60;default&#x60; at any time. The intend is to provide alternatives to our default geocoder. Each provider has its own strenghts and might fit better for certain scenarios, so it&#39;s worth to compare the different providers. To try it append the &#x60;provider&#x60;parameter to the URL like &#x60;&amp;provider&#x3D;nominatim&#x60;, the result structure should be identical in all cases - if not, please report this back to us. Keep in mind that some providers do not support certain parameters or don&#39;t return some fields, for example &#x60;osm_id&#x60; and &#x60;osm_type&#x60; are not supported by every geocoding provider. If you would like to use additional parameters of one of the providers, but it&#39;s not available for the GraphHopper Geocoding API, yet? Please contact us.  The credit costs can be different for all providers - see [here](https://support.graphhopper.com/support/solutions/articles/44000718211-what-is-one-credit-) for more information about it.  Currently, only the default provider and gisgraphy support autocompletion of partial search strings.  All providers support normal \&quot;forward\&quot; geocoding and reverse geocoding via &#x60;reverse&#x3D;true&#x60;.  #### Default (&#x60;provider&#x3D;default&#x60;)  This provider returns results of our internal geocoding engine. It is best suited for use cases with user interaction as it does prefix searches required for an \&quot;autocomplete\&quot; use case (a user types an address into a search field).  In addition to the above documented parameters the following parameters are possible:  * &#x60;osm_tag&#x60; - you can filter &#x60;key:value&#x60; or exclude places with certain OpenStreetMap tags &#x60;!key:value&#x60;. E.g. &#x60;osm_tag&#x3D;tourism:museum&#x60; or just the key &#x60;osm_tag&#x3D;tourism&#x60;. To exclude multiple tags you add multiple &#x60;osm_tag&#x60; parameters. * &#x60;location_bias_scale&#x60; - describes how much the prominence of a result should still be taken into account. Sensible values go from 0.0 (ignore prominence almost completely) to 1.0 (prominence has approximately the same influence as the location). The default is 0.2. * &#x60;zoom&#x60; - describes the radius around the center to focus on. This is a number that should correspond roughly to the map zoom parameter of a corresponding map. The default is zoom&#x3D;16. * &#x60;bbox&#x60; - the expected format is &#x60;minLon,minLat,maxLon,maxLat&#x60;. This requires reverse&#x3D;false. * &#x60;radius&#x60; - the search radius in km for the reverse search. This requires reverse&#x3D;true.  #### Nominatim (&#x60;provider&#x3D;nominatim&#x60;)  The GraphHopper Directions API uses a commercially hosted Nominatim geocoder (hosted by [OpenCageData](https://opencagedata.com/)). It is best suited for use cases without or less user interaction like batch processing or detailed location data retrieval. It is not suited for \&quot;autocomplete\&quot;.  In addition to the above documented parameters we currently support the following parameters:  * countrycode - The country code is a two letter code as defined by the ISO 3166-1 Alpha 2 standard. E.g. gb for the United Kingdom, fr for France, us for United States. * bounds - the expected format is &#x60;minLon,minLat,maxLon,maxLat&#x60;  #### Gisgraphy (&#x60;provider&#x3D;gisgraphy&#x60;)  This provider returns results from the Gisgraphy geocoder which you can try [here](https://services.gisgraphy.com/static/leaflet/index.html).  **Limitations:** The &#x60;locale&#x60; parameter is not supported. Gisgraphy does not return OSM tags or an extent.  Gisgraphy has a special autocomplete API, which you can use by adding &#x60;autocomplete&#x3D;true&#x60; (does not work with &#x60;reverse&#x3D;true&#x60;). The autocomplete API is optimized on predicting text input, but returns less information.  In addition to the above documented parameters Gisgraphy allows to use the following parameters, which can be used as documented [here](https://www.gisgraphy.com/documentation/user-guide.php#geocodingservice):  * &#x60;radius&#x60; - radius in meters * &#x60;country&#x60; - restrict search for the specified country. The value must be the ISO 3166 Alpha 2 code of the country.  #### NetToolKit (&#x60;provider&#x3D;nettoolkit&#x60;)  This provider returns results from the NetToolKit provider which is specialized for US addresses and provides a wrapper around Nominatim for other addresses. You can try it [here](https://www.nettoolkit.com/geo/demo).  The following additional NetToolKit parameters are supported (read [here](https://www.nettoolkit.com/docs/geo/geocoding) for more details): - &#x60;source&#x60;: User can choose which source provider to geocode the address, this value is \&quot;NetToolKit\&quot; by default - &#x60;country_code&#x60;: an iso-3166-2 country code (e.g : US) filter the results to the specify country code  **Limitations:** NetToolKit does not support the &#x60;locale&#x60; parameter. NetToolKit does not return OSM tags (e.g. osm_id, osm_type, osm_value).  #### OpenCage Data (&#x60;provider&#x3D;opencagedata&#x60;)  This provider returns results from the OpenCageData geocoder which you can try [here](https://geocoder.opencagedata.com/demo). The difference to the &#x60;nominatim&#x60; provider is that [other geocoders](https://opencagedata.com/credits) might be used under the hood.  In addition to the above documented parameters OpenCage Data allows to use the following parameters, which can be used as documented [here](https://geocoder.opencagedata.com/api#forward-opt):  * countrycode - The country code is a two letter code as defined by the ISO 3166-1 Alpha 2 standard. E.g. gb for the United Kingdom, fr for France, us for United States.  * bounds - the expected format is &#x60;minLon,minLat,maxLon,maxLat&#x60;  | [default to &quot;default&quot;]

### Return type

[**GeocodingResponse**](GeocodingResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

