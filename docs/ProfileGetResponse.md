# ProfileGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProfiles** | Pointer to [**[]ProfileResponse**](ProfileResponse.md) | The existing custom profiles including their ids | [optional] 

## Methods

### NewProfileGetResponse

`func NewProfileGetResponse() *ProfileGetResponse`

NewProfileGetResponse instantiates a new ProfileGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileGetResponseWithDefaults

`func NewProfileGetResponseWithDefaults() *ProfileGetResponse`

NewProfileGetResponseWithDefaults instantiates a new ProfileGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProfiles

`func (o *ProfileGetResponse) GetCustomProfiles() []ProfileResponse`

GetCustomProfiles returns the CustomProfiles field if non-nil, zero value otherwise.

### GetCustomProfilesOk

`func (o *ProfileGetResponse) GetCustomProfilesOk() (*[]ProfileResponse, bool)`

GetCustomProfilesOk returns a tuple with the CustomProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfiles

`func (o *ProfileGetResponse) SetCustomProfiles(v []ProfileResponse)`

SetCustomProfiles sets CustomProfiles field to given value.

### HasCustomProfiles

`func (o *ProfileGetResponse) HasCustomProfiles() bool`

HasCustomProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


