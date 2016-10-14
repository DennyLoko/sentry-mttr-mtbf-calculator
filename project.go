package main

type Project struct {
	Name         string `json:"name,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Organization Organization
}
