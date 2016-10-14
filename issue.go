package main

type Issue struct {
	Id		string `json:"id,omitempty"`
	Status		string `json:"status,omitempty"`
	Project		Project
	Activity		[]Activity
}
