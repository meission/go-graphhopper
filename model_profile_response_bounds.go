/*
GraphHopper Directions API

 Integrate A-to-B route planning, turn-by-turn navigation, route optimization, isochrone calculations, location clustering and other tools into your application.    ##### Authentication      1. [Sign up for GraphHopper](https://graphhopper.com/dashboard/#/signup)   2. [Create an API key](https://support.graphhopper.com/a/solutions/articles/44001976027)    Authenticate to the API by passing your key as a query parameter in every request.    ##### API Explorer    You can also try all API parts without registration in our [API explorer](https://explorer.graphhopper.com/).    ##### Client Libraries    To speed up development and make coding easier, we offer a [JavaScript client](https://github.com/graphhopper/directions-api-js-client) and a   [Java client](https://github.com/graphhopper/graphhopper/tree/master/client-hc).    ##### Optimize Response Speed    1. Reuse SSL/TLS sessions    You should utilize the SSL session to speed up responses after the initial response or use a library that does this. E.g. for Java the   [OkHttp library](https://square.github.io/okhttp/) automatically reuses SSL/TLS sessions and also the browser takes care of this automatically.   For python you can use the [`requests` library](https://requests.readthedocs.io/en/latest/user/advanced/): first you create a   session (`session = requests.Session()`) and then do requests only with this session instead of directly using \"requests\".    2. Bandwidth reduction    If you create your own client, make sure it supports http/2 and gzipped responses for best speed.   If you use the Matrix, the Route Optimization API or the  and want to solve large problems, we recommend you to reduce bandwidth   by [compressing your POST request](https://gist.github.com/karussell/82851e303ea7b3459b2dea01f18949f4) and specifying the header   as follows: `Content-Encoding: gzip`. This will also avoid the HTTP 413 error \"Request Entity Too Large\". 

API version: 1.0.0
Contact: support@graphhopper.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProfileResponseBounds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileResponseBounds{}

// ProfileResponseBounds struct for ProfileResponseBounds
type ProfileResponseBounds struct {
	// A rectangular area given as an array [minLon, minLat, maxLon, maxLat]. The profile will only work in this area.
	Bbox []float64 `json:"bbox,omitempty"`
}

// NewProfileResponseBounds instantiates a new ProfileResponseBounds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileResponseBounds() *ProfileResponseBounds {
	this := ProfileResponseBounds{}
	return &this
}

// NewProfileResponseBoundsWithDefaults instantiates a new ProfileResponseBounds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileResponseBoundsWithDefaults() *ProfileResponseBounds {
	this := ProfileResponseBounds{}
	return &this
}

// GetBbox returns the Bbox field value if set, zero value otherwise.
func (o *ProfileResponseBounds) GetBbox() []float64 {
	if o == nil || IsNil(o.Bbox) {
		var ret []float64
		return ret
	}
	return o.Bbox
}

// GetBboxOk returns a tuple with the Bbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileResponseBounds) GetBboxOk() ([]float64, bool) {
	if o == nil || IsNil(o.Bbox) {
		return nil, false
	}
	return o.Bbox, true
}

// HasBbox returns a boolean if a field has been set.
func (o *ProfileResponseBounds) HasBbox() bool {
	if o != nil && !IsNil(o.Bbox) {
		return true
	}

	return false
}

// SetBbox gets a reference to the given []float64 and assigns it to the Bbox field.
func (o *ProfileResponseBounds) SetBbox(v []float64) {
	o.Bbox = v
}

func (o ProfileResponseBounds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileResponseBounds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bbox) {
		toSerialize["bbox"] = o.Bbox
	}
	return toSerialize, nil
}

type NullableProfileResponseBounds struct {
	value *ProfileResponseBounds
	isSet bool
}

func (v NullableProfileResponseBounds) Get() *ProfileResponseBounds {
	return v.value
}

func (v *NullableProfileResponseBounds) Set(val *ProfileResponseBounds) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileResponseBounds) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileResponseBounds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileResponseBounds(val *ProfileResponseBounds) *NullableProfileResponseBounds {
	return &NullableProfileResponseBounds{value: val, isSet: true}
}

func (v NullableProfileResponseBounds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileResponseBounds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


