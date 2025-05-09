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

// checks if the CustomModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomModel{}

// CustomModel The custom_model modifies the routing behaviour of the specified profile. See the [detailed documentation](#tag/Custom-Model).  Below is a complete request example in Berlin that limits all speeds to 100km/h, excludes motorways and makes shorter routes a bit more likely than the default due to a higher distance_influence.  Note that it also includes the `\"ch.disabled\": true` parameter which is required to make use of `custom_model`.  ```json {   \"points\": [     [       13.31543,       52.509535     ],     [       13.29779,       52.512434     ]   ],   \"profile\": \"car\",   \"ch.disable\": true,   \"custom_model\": {     \"speed\": [       {         \"if\": \"true\",         \"limit_to\": \"100\"       }     ],     \"priority\": [       {         \"if\": \"road_class == MOTORWAY\",         \"multiply_by\": \"0\"       }     ],     \"distance_influence\": 100   } }  ``` 
type CustomModel struct {
	// See [speed customization](#tag/Custom-Model/Customizing-speed)
	Speed []map[string]interface{} `json:"speed,omitempty"`
	// See [priority customization](#tag/Custom-Model/Customizing-priority)
	Priority []map[string]interface{} `json:"priority,omitempty"`
	// Use higher values to prefer shorter routes. See [here](#tag/Custom-Model/Customizing-distance_influence) for more details.
	DistanceInfluence *float32 `json:"distance_influence,omitempty"`
	// Areas are given in a GeoJson format as FeatureCollection. See more details and an example [here](#tag/Custom-Model/Define-areas). 
	Areas FeatureCollection `json:"areas,omitempty"`
}

// NewCustomModel instantiates a new CustomModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomModel() *CustomModel {
	this := CustomModel{}
	var distanceInfluence float32 = 70
	this.DistanceInfluence = &distanceInfluence
	return &this
}

// NewCustomModelWithDefaults instantiates a new CustomModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomModelWithDefaults() *CustomModel {
	this := CustomModel{}
	var distanceInfluence float32 = 70
	this.DistanceInfluence = &distanceInfluence
	return &this
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *CustomModel) GetSpeed() []map[string]interface{} {
	if o == nil || IsNil(o.Speed) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomModel) GetSpeedOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *CustomModel) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given []map[string]interface{} and assigns it to the Speed field.
func (o *CustomModel) SetSpeed(v []map[string]interface{}) {
	o.Speed = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CustomModel) GetPriority() []map[string]interface{} {
	if o == nil || IsNil(o.Priority) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomModel) GetPriorityOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CustomModel) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given []map[string]interface{} and assigns it to the Priority field.
func (o *CustomModel) SetPriority(v []map[string]interface{}) {
	o.Priority = v
}

// GetDistanceInfluence returns the DistanceInfluence field value if set, zero value otherwise.
func (o *CustomModel) GetDistanceInfluence() float32 {
	if o == nil || IsNil(o.DistanceInfluence) {
		var ret float32
		return ret
	}
	return *o.DistanceInfluence
}

// GetDistanceInfluenceOk returns a tuple with the DistanceInfluence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomModel) GetDistanceInfluenceOk() (*float32, bool) {
	if o == nil || IsNil(o.DistanceInfluence) {
		return nil, false
	}
	return o.DistanceInfluence, true
}

// HasDistanceInfluence returns a boolean if a field has been set.
func (o *CustomModel) HasDistanceInfluence() bool {
	if o != nil && !IsNil(o.DistanceInfluence) {
		return true
	}

	return false
}

// SetDistanceInfluence gets a reference to the given float32 and assigns it to the DistanceInfluence field.
func (o *CustomModel) SetDistanceInfluence(v float32) {
	o.DistanceInfluence = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *CustomModel) GetAreas() FeatureCollection {
	if o == nil || IsNil(o.Areas) {
		var ret FeatureCollection
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomModel) GetAreasOk() (FeatureCollection, bool) {
	if o == nil || IsNil(o.Areas) {
		return FeatureCollection{}, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *CustomModel) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given FeatureCollection and assigns it to the Areas field.
func (o *CustomModel) SetAreas(v FeatureCollection) {
	o.Areas = v
}

func (o CustomModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.DistanceInfluence) {
		toSerialize["distance_influence"] = o.DistanceInfluence
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

type NullableCustomModel struct {
	value *CustomModel
	isSet bool
}

func (v NullableCustomModel) Get() *CustomModel {
	return v.value
}

func (v *NullableCustomModel) Set(val *CustomModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomModel(val *CustomModel) *NullableCustomModel {
	return &NullableCustomModel{value: val, isSet: true}
}

func (v NullableCustomModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


