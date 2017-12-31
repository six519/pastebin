package pastebin

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"reflect"
	"strings"
	"errors"
	"encoding/xml"
)

const API_URL string = "https://pastebin.com/api/"

type Pastes struct {
	XMLName xml.Name `xml:"pastes"`
	Pastes []Paste `xml:"paste"`
}

type Paste struct {
	XMLName xml.Name `xml:"paste"`
	Paste_key string `xml:"paste_key"`
	Paste_date string `xml:"paste_date"`
	Paste_title string `xml:"paste_title"`
	Paste_size string `xml:"paste_size"`
	Paste_expire_date string `xml:"paste_expire_date"`
	Paste_private string `xml:"paste_private"`
	Paste_format_long string `xml:"paste_format_long"`
	Paste_format_short string `xml:"paste_format_short"`
	Paste_url string `xml:"paste_url"`
	Paste_hits string `xml:"paste_hits"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	User_name string `xml:"user_name"`
	User_format_short string `xml:"user_format_short"`
	User_expiration string `xml:"user_expiration"`
	User_avatar_url string `xml:"user_avatar_url"`
	User_private string `xml:"user_private"`
	User_website string `xml:"user_website"`
	User_email string `xml:"user_email"`
	User_location string `xml:"user_location"`
	User_account_type string `xml:"user_account_type"`
}

type Pastebin struct {
	Api_dev_key string
	Api_user_key string
}

type PastebinOption struct {
	Api_paste_name string
	Api_paste_format string
	Api_paste_private string
	Api_paste_expire_date string
	Api_results_limit string

	api_user_key string
}

func convertToValues(options PastebinOption) url.Values {
	v := url.Values{}

	r := reflect.ValueOf(&options)
	s := r.Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		fieldName := typeOfT.Field(i).Name
		f := reflect.Indirect(r).FieldByName(fieldName)
		fieldValue := f.String()

		if fieldValue != "" {
			v.Set(strings.ToLower(fieldName), fieldValue)
		}
	}

	return v
}

func request(pastebin Pastebin, method string, params url.Values) (string, error) {

	if pastebin.Api_user_key != "" {
		params.Set("api_user_key", pastebin.Api_user_key)
	}

	params.Set("api_dev_key", pastebin.Api_dev_key)

	resp, err := http.PostForm(API_URL + method, params)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)

	if strings.Contains(strBody, "Bad API request") {
		return "", errors.New(strBody)
	}

	return strBody, nil
}

func (pastebin Pastebin) Paste(api_paste_code string, other_options PastebinOption) (string, error) {

	v := convertToValues(other_options)
	v.Set("api_option", "paste")
	v.Set("api_paste_code", api_paste_code)
	
	return request(pastebin, "api_post.php", v)
}

func (pastebin *Pastebin) Login(username string, password string) (string, error) {
	v := url.Values{}
	v.Set("api_user_name", username)
	v.Set("api_user_password", password)
	ret, err := request(*pastebin, "api_login.php", v)

	if err == nil {
		pastebin.Api_user_key = ret
	}

	return ret, err
}

func (pastebin Pastebin) Trends() (Pastes, error) {
	v := url.Values{}
	v.Set("api_option", "trends")
	ret, err := request(pastebin, "api_post.php", v)

	pastes := Pastes{}
	ret = "<pastes>" + ret + "</pastes>"
	xml.Unmarshal([]byte(ret), &pastes)
	
	return pastes, err
}

func (pastebin Pastebin) List(other_options PastebinOption) (Pastes, error) {
	v := convertToValues(other_options)
	v.Set("api_option", "list")
	
	ret, err := request(pastebin, "api_post.php", v)

	pastes := Pastes{}
	ret = "<pastes>" + ret + "</pastes>"
	xml.Unmarshal([]byte(ret), &pastes)
	
	return pastes, err
}

func (pastebin Pastebin) Delete(api_paste_key string) (string, error) {
	v := url.Values{}
	v.Set("api_option", "delete")
	v.Set("api_paste_key", api_paste_key)

	return request(pastebin, "api_post.php", v)
}

func (pastebin Pastebin) UserDetails() (User, error) {
	v := url.Values{}
	v.Set("api_option", "userdetails")
	ret, err := request(pastebin, "api_post.php", v)

	user := User{}
	xml.Unmarshal([]byte(ret), &user)
	
	return user, err
}

func (pastebin Pastebin) ShowPaste(api_paste_key string) (string, error) {
	v := url.Values{}
	v.Set("api_option", "show_paste")
	v.Set("api_paste_key", api_paste_key)

	return request(pastebin, "api_raw.php", v)
}

func (pastebin Pastebin) GetKey() string {
	return pastebin.Api_dev_key
}