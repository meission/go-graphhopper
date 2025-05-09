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

// checks if the BBox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BBox{}

// BBox struct for BBox
type BBox struct {
	// A rectangular area given as an array [minLon, minLat, maxLon, maxLat]. The created profile will only work in this area. The maximum area size is 160 000 square kilometers.
	Bbox []float64 `json:"bbox,omitempty"`
}

// NewBBox instantiates a new BBox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBBox() *BBox {
	this := BBox{}
	return &this
}

// NewBBoxWithDefaults instantiates a new BBox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBBoxWithDefaults() *BBox {
	this := BBox{}
	return &this
}

// GetBbox returns the Bbox field value if set, zero value otherwise.
func (o *BBox) GetBbox() []float64 {
	if o == nil || IsNil(o.Bbox) {
		var ret []float64
		return ret
	}
	return o.Bbox
}

// GetBboxOk returns a tuple with the Bbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BBox) GetBboxOk() ([]float64, bool) {
	if o == nil || IsNil(o.Bbox) {
		return nil, false
	}
	return o.Bbox, true
}

// HasBbox returns a boolean if a field has been set.
func (o *BBox) HasBbox() bool {
	if o != nil && !IsNil(o.Bbox) {
		return true
	}

	return false
}

// SetBbox gets a reference to the given []float64 and assigns it to the Bbox field.
func (o *BBox) SetBbox(v []float64) {
	o.Bbox = v
}

func (o BBox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BBox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bbox) {
		toSerialize["bbox"] = o.Bbox
	}
	return toSerialize, nil
}

type NullableBBox struct {
	value *BBox
	isSet bool
}

func (v NullableBBox) Get() *BBox {
	return v.value
}

func (v *NullableBBox) Set(val *BBox) {
	v.value = val
	v.isSet = true
}

func (v NullableBBox) IsSet() bool {
	return v.isSet
}

func (v *NullableBBox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBBox(val *BBox) *NullableBBox {
	return &NullableBBox{value: val, isSet: true}
}

func (v NullableBBox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBBox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


