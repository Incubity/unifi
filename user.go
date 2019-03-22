// Copyright (c) 2014 The unifi Authors. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

package unifi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"
)

type User struct {
	FirstSeen int `json:"first_seen"`
	Hostname  string
	IsGuest   bool `json:"is_guest"`
	LastSeen  int  `json:"last_seen"`
	Mac       string
	Oui       string
}

// Request structs
type NewUser struct {
	UserGroupID string `json:"usergroup_id"`
	Mac         string `json:"mac"`
	Name        string `json:"name"`
}

type Data struct {
	Data NewUser `json:"data"`
}

type Objects struct {
	Objects []Data `json:"objects"`
}

//Value with parameters for create New User
var Nu NewUser

//Functions creating new Users

func (u *Unifi) NewUser(site *Site, nu NewUser) ([]User, error) {
	var response struct {
		Data []User
		Meta meta
	}

	Nu = nu
	err := u.parseNewUser(site, "group/user", &response)
	return response.Data, err
}

func (u *Unifi) apicmdNewUser(site *Site, cmd string) ([]byte, error) {
	jsonData := Objects{[]Data{{Nu}}}

	// Setup url
	cmdurl := u.apiURL
	cmdurl += fmt.Sprintf("s/%s/%s", site.Name, cmd)

	data, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	fmt.Println(bytes.NewBuffer(data))

	val := url.Values{"json": {string(data)}}

	resp, err := u.client.PostForm(cmdurl, val)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (u *Unifi) parseNewUser(site *Site, cmd string, v interface{}) error {
	body, err := u.apicmdNewUser(site, cmd)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &v); err != nil {
		log.Println(cmd)
		log.Println(string(body))
		return err
	}
	m := reflect.ValueOf(v).Elem().FieldByName("Meta").Interface().(meta)
	if m.Rc != "ok" {
		return fmt.Errorf("Bad request: %s", m.Rc)
	}
	return nil
}
