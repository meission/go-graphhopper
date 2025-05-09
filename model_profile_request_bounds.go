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
	"fmt"
)


// ProfileRequestBounds struct for ProfileRequestBounds
type ProfileRequestBounds struct {
	BBox *BBox
	FeatureCollection *FeatureCollection
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProfileRequestBounds) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BBox
	err = json.Unmarshal(data, &dst.BBox);
	if err == nil {
		jsonBBox, _ := json.Marshal(dst.BBox)
		if string(jsonBBox) == "{}" { // empty struct
			dst.BBox = nil
		} else {
			return nil // data stored in dst.BBox, return on the first match
		}
	} else {
		dst.BBox = nil
	}

	// try to unmarshal JSON data into FeatureCollection
	err = json.Unmarshal(data, &dst.FeatureCollection);
	if err == nil {
		jsonFeatureCollection, _ := json.Marshal(dst.FeatureCollection)
		if string(jsonFeatureCollection) == "{}" { // empty struct
			dst.FeatureCollection = nil
		} else {
			return nil // data stored in dst.FeatureCollection, return on the first match
		}
	} else {
		dst.FeatureCollection = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProfileRequestBounds)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProfileRequestBounds) MarshalJSON() ([]byte, error) {
	if src.BBox != nil {
		return json.Marshal(&src.BBox)
	}

	if src.FeatureCollection != nil {
		return json.Marshal(&src.FeatureCollection)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableProfileRequestBounds struct {
	value *ProfileRequestBounds
	isSet bool
}

func (v NullableProfileRequestBounds) Get() *ProfileRequestBounds {
	return v.value
}

func (v *NullableProfileRequestBounds) Set(val *ProfileRequestBounds) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileRequestBounds) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileRequestBounds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileRequestBounds(val *ProfileRequestBounds) *NullableProfileRequestBounds {
	return &NullableProfileRequestBounds{value: val, isSet: true}
}

func (v NullableProfileRequestBounds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileRequestBounds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


