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

// checks if the Algorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Algorithm{}

// Algorithm Use `objectives` instead.
type Algorithm struct {
	ProblemType *string `json:"problem_type,omitempty"`
	Objective *string `json:"objective,omitempty"`
}

// NewAlgorithm instantiates a new Algorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgorithm() *Algorithm {
	this := Algorithm{}
	return &this
}

// NewAlgorithmWithDefaults instantiates a new Algorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgorithmWithDefaults() *Algorithm {
	this := Algorithm{}
	return &this
}

// GetProblemType returns the ProblemType field value if set, zero value otherwise.
func (o *Algorithm) GetProblemType() string {
	if o == nil || IsNil(o.ProblemType) {
		var ret string
		return ret
	}
	return *o.ProblemType
}

// GetProblemTypeOk returns a tuple with the ProblemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetProblemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProblemType) {
		return nil, false
	}
	return o.ProblemType, true
}

// HasProblemType returns a boolean if a field has been set.
func (o *Algorithm) HasProblemType() bool {
	if o != nil && !IsNil(o.ProblemType) {
		return true
	}

	return false
}

// SetProblemType gets a reference to the given string and assigns it to the ProblemType field.
func (o *Algorithm) SetProblemType(v string) {
	o.ProblemType = &v
}

// GetObjective returns the Objective field value if set, zero value otherwise.
func (o *Algorithm) GetObjective() string {
	if o == nil || IsNil(o.Objective) {
		var ret string
		return ret
	}
	return *o.Objective
}

// GetObjectiveOk returns a tuple with the Objective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Algorithm) GetObjectiveOk() (*string, bool) {
	if o == nil || IsNil(o.Objective) {
		return nil, false
	}
	return o.Objective, true
}

// HasObjective returns a boolean if a field has been set.
func (o *Algorithm) HasObjective() bool {
	if o != nil && !IsNil(o.Objective) {
		return true
	}

	return false
}

// SetObjective gets a reference to the given string and assigns it to the Objective field.
func (o *Algorithm) SetObjective(v string) {
	o.Objective = &v
}

func (o Algorithm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Algorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProblemType) {
		toSerialize["problem_type"] = o.ProblemType
	}
	if !IsNil(o.Objective) {
		toSerialize["objective"] = o.Objective
	}
	return toSerialize, nil
}

type NullableAlgorithm struct {
	value *Algorithm
	isSet bool
}

func (v NullableAlgorithm) Get() *Algorithm {
	return v.value
}

func (v *NullableAlgorithm) Set(val *Algorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgorithm(val *Algorithm) *NullableAlgorithm {
	return &NullableAlgorithm{value: val, isSet: true}
}

func (v NullableAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


